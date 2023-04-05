//go:build linux

package utils

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/livekit/protocol/logger"
)

var (
	usageRegex = regexp.MustCompile("usage_usec ([0-9]+)")
)

const (
	cpuStatsPathV1 = "/sys/fs/cgroup/cpu,cpuacct/cpuacct.usage"
	cpuStatsPathV2 = "/sys/fs/cgroup/cpu.stat"

	numCPUPathV1 = "/sys/fs/cgroup/cpu,cpuacct/cpuacct.usage_percpu"
	numCPUPathV2 = "/sys/fs/cgroup/cpu.max"
)

type cpuInfoGetter interface {
	getTotalCPUTime() (int64, error)
	numCPU() (int, error)
}

type cgroupCPUMonitor struct {
	lastSampleTime   int64
	lastTotalCPUTime int64
	nCPU             int

	cg cpuInfoGetter
}

func newPlatformCPUMonitor() (platformCPUMonitor, error) {
	// probe for the cgroup version
	var cg cpuInfoGetter
	for k, v := range map[string]func() cpuInfoGetter{
		cpuStatsPathV1: newCpuInfoGetterV1,
		cpuStatsPathV2: newCpuInfoGetterV2,
	} {
		e, err := fileExists(k)
		if err != nil {
			return nil, err
		}
		if e {
			cg = v()
			break
		}
	}
	if cg == nil {
		logger.Infow("failed reading cgroup specific cpu stats, falling back to system wide implementation")
		return newOsstatCPUMonitor()
	}

	cpu, err := cg.getTotalCPUTime()
	if err != nil {
		return nil, err
	}

	nCPU, err := cg.numCPU()
	if err != nil {
		return nil, err
	}

	return &cgroupCPUMonitor{
		lastSampleTime:   time.Now().UnixNano(),
		lastTotalCPUTime: cpu,
		nCPU:             nCPU,
		cg:               cg,
	}, nil
}

func (p *cgroupCPUMonitor) getCPUIdle() (float64, error) {
	next, err := p.cg.getTotalCPUTime()
	if err != nil {
		return 0, err
	}
	t := time.Now().UnixNano()

	duration := t - p.lastSampleTime
	cpuTime := next - p.lastTotalCPUTime

	busyRatio := float64(cpuTime) / float64(duration)
	idleRatio := float64(p.nCPU) - busyRatio

	// Clamp the value as we do not get all the timestamps at the same time
	if idleRatio > float64(p.nCPU) {
		idleRatio = float64(p.nCPU)
	} else if idleRatio < 0 {
		idleRatio = 0
	}

	p.lastSampleTime = t
	p.lastTotalCPUTime = next

	return idleRatio, nil
}

func (p *cgroupCPUMonitor) numCPU() int {
	return p.nCPU
}

type cpuInfoGetterV1 struct {
}

func newCpuInfoGetterV1() cpuInfoGetter {
	return &cpuInfoGetterV1{}
}

func (cg *cpuInfoGetterV1) getTotalCPUTime() (int64, error) {
	b, err := os.ReadFile(cpuStatsPathV1)
	if err != nil {
		return 0, err
	}

	// Skip the trailing EOL
	i, err := strconv.ParseInt(string(b[:len(b)-1]), 10, 64)
	if err != nil {
		return 0, err
	}

	return i, nil
}

func (cg *cpuInfoGetterV1) numCPU() (int, error) {
	b, err := os.ReadFile(numCPUPathV1)
	if err != nil {
		return 0, err
	}

	// Remove trailing new line if any
	s := strings.TrimSuffix(string(b), "\n")

	// Remove trailing space if any
	s = strings.TrimSuffix(s, " ")

	m := strings.Split(s, " ")
	if len(m) == 0 {
		return 0, errors.New("could not parse cpu stats")
	}

	cpuCount := 0
	for _, v := range m {
		if v != "0" {
			cpuCount++
		}
	}

	return cpuCount, nil
}

type cpuInfoGetterV2 struct {
}

func newCpuInfoGetterV2() cpuInfoGetter {
	return &cpuInfoGetterV2{}
}

func (cg *cpuInfoGetterV2) getTotalCPUTime() (int64, error) {
	b, err := os.ReadFile(cpuStatsPathV2)
	if err != nil {
		return 0, err
	}

	m := usageRegex.FindSubmatch(b)
	if len(m) <= 1 {
		return 0, errors.New("could not parse cpu stats")
	}

	i, err := strconv.ParseInt(string(m[1]), 10, 64)
	if err != nil {
		return 0, err
	}

	// Caller expexts time in ns
	return i * 1000, nil
}

func (cg *cpuInfoGetterV2) numCPU() (int, error) {
	b, err := os.ReadFile(numCPUPathV2)
	if err != nil {
		return 0, err
	}

	s := strings.TrimSuffix(string(b), "\n")

	m := strings.Split(s, " ")
	if len(m) <= 1 {
		return 0, errors.New("could not parse cpu stats")
	}

	if m[0] == "max" {
		// No quota
		return runtime.NumCPU(), nil
	} else {
		n, err := strconv.ParseInt(string(m[0]), 10, 64)
		if err != nil {
			return 0, err
		}

		d, err := strconv.ParseInt(string(m[1]), 10, 64)
		if err != nil {
			return 0, err
		}

		return int(n / d), nil
	}
}

func fileExists(path string) (bool, error) {
	_, err := os.Lstat(path)
	switch {
	case err == nil:
		return true, nil
	case errors.Is(err, os.ErrNotExist):
		return false, nil
	default:
		return false, err
	}
}
