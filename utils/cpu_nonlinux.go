//go:build !linux

package utils

import (
	"runtime"

	"github.com/mackerelio/go-osstat/cpu"
)

type platformCPUMonitor struct {
	prev *cpu.Stats
}

func newPlatformCPUMonitor() (*platformCPUMonitor, error) {
	stats, err := cpu.Get()
	if err != nil {
		return nil, err
	}

	return &platformCPUMonitor{
		prev: stats,
	}, nil
}

func (p *platformCPUMonitor) getCPUIdle() (float64, error) {
	next, err := cpu.Get()
	if err != nil {
		return 0, err
	}
	idleRatio := float64(next.Idle-p.prev.Idle) / float64(next.Total-p.prev.Total)
	p.prev = next

	return float64(runtime.NumCPU()) * idleRatio, nil

}
