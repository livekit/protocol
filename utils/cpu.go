package utils

import (
	"time"

	"github.com/frostbyte73/go-throttle"
	"github.com/livekit/protocol/logger"
	"go.uber.org/atomic"
)

// This object returns cgroup quota aware cpu stats. On other systems than Linux,
// it falls back to full system stats

type platformCPUMonitor interface {
	getCPUIdle() (float64, error)
	numCPU() int
}

type CPUStats struct {
	idleCPUs atomic.Float64
	platform platformCPUMonitor

	updateCallback  func(idle float64)
	warningThrottle func(func())
	closeChan       chan struct{}
}

func NewCPUStats(updateCallback func(idle float64)) (*CPUStats, error) {
	p, err := newPlatformCPUMonitor()
	if err != nil {
		return nil, err
	}

	c := &CPUStats{
		platform:        p,
		warningThrottle: throttle.New(time.Minute),
		updateCallback:  updateCallback,
		closeChan:       make(chan struct{}),
	}

	go c.monitorCPULoad()

	return c, nil
}

func (c *CPUStats) GetCPUIdle() float64 {
	return c.idleCPUs.Load()
}

func (c *CPUStats) NumCPU() int {
	return c.platform.numCPU()
}

func (c *CPUStats) Stop() {
	close(c.closeChan)
}

func (c *CPUStats) monitorCPULoad() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-c.closeChan:
			return
		case <-ticker.C:
			idle, err := c.platform.getCPUIdle()
			if err != nil {
				logger.Errorw("failed retrieving CPU idle", err)
				continue
			}

			c.idleCPUs.Store(idle)
			idleRatio := idle / float64(c.platform.numCPU())

			if idleRatio < 0.1 {
				c.warningThrottle(func() { logger.Infow("high cpu load", "load", 1-idleRatio) })
			}

			if c.updateCallback != nil {
				c.updateCallback(idle)
			}
		}
	}
}
