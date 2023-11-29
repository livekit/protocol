// Copyright 2023 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"time"

	"github.com/frostbyte73/core"
	"github.com/mitchellh/go-ps"
	"github.com/prometheus/procfs"
	"go.uber.org/atomic"

	"github.com/livekit/protocol/logger"
)

// This object returns cgroup quota aware cpu stats. On other systems than Linux,
// it falls back to full system stats

type platformCPUMonitor interface {
	getCPUIdle() (float64, error)
	numCPU() float64
}

type CPUStats struct {
	idleCPUs atomic.Float64
	platform platformCPUMonitor

	idleCallback    func(idle float64)
	procCallback    func(idle float64, usage map[int]float64)
	warningThrottle core.Throttle
	closeChan       chan struct{}
}

func NewCPUStats(updateCallback func(idle float64)) (*CPUStats, error) {
	p, err := newPlatformCPUMonitor()
	if err != nil {
		return nil, err
	}

	c := &CPUStats{
		platform:        p,
		warningThrottle: core.NewThrottle(time.Minute),
		idleCallback:    updateCallback,
		closeChan:       make(chan struct{}),
	}

	go c.monitorCPULoad()

	return c, nil
}

func NewProcCPUStats(updateCallback func(idle float64, usage map[int]float64)) (*CPUStats, error) {
	p, err := newPlatformCPUMonitor()
	if err != nil {
		return nil, err
	}

	c := &CPUStats{
		platform:        p,
		warningThrottle: core.NewThrottle(time.Minute),
		procCallback:    updateCallback,
		closeChan:       make(chan struct{}),
	}

	go c.monitorProcCPULoad()

	return c, nil
}

func (c *CPUStats) GetCPUIdle() float64 {
	return c.idleCPUs.Load()
}

func (c *CPUStats) NumCPU() float64 {
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
			idleRatio := idle / c.platform.numCPU()

			if idleRatio < 0.1 {
				c.warningThrottle(func() { logger.Infow("high cpu load", "load", 1-idleRatio) })
			}

			if c.idleCallback != nil {
				c.idleCallback(idle)
			}
		}
	}
}

func (c *CPUStats) monitorProcCPULoad() {
	numCPU := c.platform.numCPU()
	fs, err := procfs.NewFS(procfs.DefaultMountPoint)
	if err != nil {
		logger.Errorw("failed read proc fs", err)
		return
	}

	self, err := fs.Self()
	if err != nil {
		logger.Errorw("failed to read self", err)
		return
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	var prevTotalTime float64
	var prevStats map[int]procfs.ProcStat
	for {
		select {
		case <-c.closeChan:
			return
		case <-ticker.C:
			nextStats := make(map[int]procfs.ProcStat)
			procs, err := procfs.AllProcs()
			if err != nil {
				logger.Errorw("failed to read processes", err)
				continue
			}

			total, err := fs.Stat()
			if err != nil {
				logger.Errorw("failed to read stats", err)
				continue
			}

			ppids := make(map[int]int)
			for _, proc := range procs {
				nextStats[proc.PID], err = proc.Stat()
				if err != nil {
					logger.Errorw("failed to read proc stats", err)
					continue
				}
				if proc.PID != self.PID {
					ppids[proc.PID], err = getPPID(proc.PID)
					if err != nil {
						logger.Errorw("failed to get PPID", err)
						continue
					}
				}
			}

			nextTotalTime := total.CPUTotal.User + total.CPUTotal.Nice + total.CPUTotal.System + total.CPUTotal.Idle

			usage := make(map[int]float64)
			totalUsage := 0.0
			for pid, stat := range nextStats {
				t := float64(stat.UTime + stat.STime - prevStats[pid].UTime - prevStats[pid].STime)
				if t == 0 {
					continue
				}

				for ppids[pid] != self.PID && ppids[pid] != 0 {
					// attribute usage to parent process, stopping before service process
					pid = ppids[pid]
				}

				s := numCPU * t / 100 / (nextTotalTime - prevTotalTime)
				usage[pid] += s
				totalUsage += s
			}

			idle := numCPU - totalUsage
			c.idleCPUs.Store(idle)

			if c.procCallback != nil {
				c.procCallback(idle, usage)
			}

			prevTotalTime = nextTotalTime
			prevStats = nextStats
		}
	}
}

func getPPID(pid int) (int, error) {
	p, err := ps.FindProcess(pid)
	if err != nil {
		return 0, err
	}
	return p.PPid(), nil
}
