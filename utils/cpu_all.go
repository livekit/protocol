package utils

import (
	"runtime"

	"github.com/mackerelio/go-osstat/cpu"
)

type osstatCPUMonitor struct {
	prev *cpu.Stats
}

func newOsstatCPUMonitor() (*osstatCPUMonitor, error) {
	stats, err := cpu.Get()
	if err != nil {
		return nil, err
	}

	return &osstatCPUMonitor{
		prev: stats,
	}, nil
}

func (p *osstatCPUMonitor) getCPUIdle() (float64, error) {
	next, err := cpu.Get()
	if err != nil {
		return 0, err
	}
	idleRatio := float64(next.Idle-p.prev.Idle) / float64(next.Total-p.prev.Total)
	p.prev = next

	return float64(p.numCPU()) * idleRatio, nil
}

func (p *osstatCPUMonitor) numCPU() int {
	return runtime.NumCPU()
}
