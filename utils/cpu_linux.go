//go:build linux

package utils

import (
	"errors"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"time"
)

var (
	usageRegex = regexp.MustCompile("usage_usec ([0-9]+)")
)

const (
	cpuStatsPathV1 = "/sys/fs/cgroup/cpu,cpuacct/cpuacct.usage"
	cpuStatsPathV2 = "/sys/fs/cgroup/cpu.stat"
)

type platformCPUMonitor struct {
	lastSampleTime   int64
	lastTotalCPUTime int64

	cpuTimeFunc func() (int64, error)
}

func newPlatformCPUMonitor() (*platformCPUMonitor, error) {
	// probe for the cgroup version
	var cpuTimeFunc func() (int64, error)
	for k, v := range map[string]func() (int64, error){
		cpuStatsPathV1: getTotalCPUTimeV1,
		cpuStatsPathV2: getTotalCPUTimeV2,
	} {
		e, err := fileExists(k)
		if err != nil {
			return nil, err
		}
		if e {
			cpuTimeFunc = v
		}
		break
	}
	if cpuTimeFunc == nil {
		return nil, errors.New("failed reading cpu stats file")
	}

	cpu, err := cpuTimeFunc()
	if err != nil {
		return nil, err
	}

	return &platformCPUMonitor{
		lastSampleTime:   time.Now().UnixNano() / 1000,
		lastTotalCPUTime: cpu,
		cpuTimeFunc:      cpuTimeFunc,
	}, nil
}

func (p *platformCPUMonitor) getCPUIdle() (float64, error) {
	next, err := p.cpuTimeFunc()
	if err != nil {
		return 0, err
	}
	t := time.Now().UnixNano()

	duration := t - p.lastSampleTime
	cpuTime := next - p.lastTotalCPUTime

	busyRatio := float64(cpuTime) / float64(duration)
	idleRatio := float64(runtime.NumCPU()) - busyRatio

	// Clamp the value as we do not get all the timestamps at the same time
	if idleRatio > float64(runtime.NumCPU()) {
		idleRatio = float64(runtime.NumCPU())
	} else if idleRatio < 0 {
		idleRatio = 0
	}

	p.lastSampleTime = t
	p.lastTotalCPUTime = next

	return idleRatio, nil
}

func getTotalCPUTimeV1() (int64, error) {
	b, err := os.ReadFile(cpuStatsPathV1)
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

func getTotalCPUTimeV2() (int64, error) {
	b, err := os.ReadFile(cpuStatsPathV2)
	if err != nil {
		return 0, err
	}

	i, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return 0, err
	}

	return i, nil
}

func fileExists(path string) (bool, error) {
	_, err := os.Lstat(path)
	switch err {
	case nil:
		return true, nil
	case os.ErrNotExist:
		return false, nil
	default:
		return false, err
	}
}
