//go:build linux

package utils

import (
	"errors"
	"os"
	"regexp"
	"strconv"
	"time"
)

var (
	usageRegex = regexp.MustCompile("usage_usec ([0-9]+)")
)

const (
	cpuStatPath = "/sys/fs/cgroup/cpu.stat"
)

type platformCPUMonitor struct {
	lastSampleTime   int64
	lastTotalCPUTime int64
}

func newPlatformCPUMonitor() (*platformCPUMonitor, error) {
	cpu, err := getTotalCPUTime()
	if err != nil {
		return nil, err
	}

	return &platformCPUMonitor{
		lastSampleTime:   time.Now().UnixNano(),
		lastTotalCPUTime: cpu,
	}, nil
}

func (p *platformCPUMonitor) getCPUIdle() (float64, error) {
	next, err := getTotalCPUTime()
	if err != nil {
		return 0, err
	}
	t := time.Now().UnixNano() / 1000

	durationUSec := t - p.lastSampleTime
	cpuTime := next - p.lastTotalCPUTime

	idleRatio := float64(cpuTime) / float64(durationUSec)

	p.lastSampleTime = t
	p.lastTotalCPUTime = next

	return idleRatio, nil
}

func getTotalCPUTime() (int64, error) {
	b, err := os.ReadFile(cpuStatPath)
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

	return i, nil
}
