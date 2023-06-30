package utils

import (
	"runtime"

	"github.com/mackerelio/go-osstat/cpu"
)

type osStatCPUMonitor struct {
	prev *cpu.Stats
}

func newOSStatCPUMonitor() (*osStatCPUMonitor, error) {
	stats, err := cpu.Get()
	if err != nil {
		return nil, err
	}

	return &osStatCPUMonitor{
		prev: stats,
	}, nil
}

func (p *osStatCPUMonitor) getCPUIdle() (float64, error) {
	next, err := cpu.Get()
	if err != nil {
		return 0, err
	}
	idleRatio := float64(next.Idle-p.prev.Idle) / float64(next.Total-p.prev.Total)
	p.prev = next

	return p.numCPU() * idleRatio, nil
}

func (p *osStatCPUMonitor) numCPU() float64 {
	return float64(runtime.NumCPU())
}
