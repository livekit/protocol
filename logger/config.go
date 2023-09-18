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

package logger

import "sync"

type Config struct {
	JSON  bool   `yaml:"json"`
	Level string `yaml:"level"`
	// true to enable log sampling, where the same log message and level will be throttled.
	// we have two layers of sampling
	// 1. global sampling - within a second, it will log the first SampleInitial, then every SampleInterval messages.
	// 2. per participant/track sampling - to be used with Logger.WithItemSampler(). This would be used to throttle
	//    the logs for a particular participant/track.
	Sample bool `yaml:"sample,omitempty"`

	ComponentLevels map[string]string `yaml:"component_levels,omitempty"`

	// global sampling per server
	// when sampling, the first N logs will be logged
	SampleInitial int `yaml:"sample_initial,omitempty"`
	// when sampling, every Mth log will be logged
	SampleInterval int `yaml:"sample_interval,omitempty"`

	// participant/track level sampling
	ItemSampleSeconds  int `yaml:"item_sample_seconds,omitempty"`
	ItemSampleInitial  int `yaml:"item_sample_initial,omitempty"`
	ItemSampleInterval int `yaml:"item_sample_interval,omitempty"`
}

type ConfigObserver func(*Config) error

type updatableConfig struct {
	*Config

	lock               sync.Mutex
	onUpdatedCallbacks []ConfigObserver
}

var uConfig updatableConfig

func (c *Config) Update(o *Config) error {
	u := &uConfig
	u.lock.Lock()
	u.JSON = o.JSON
	u.Level = o.Level
	u.Sample = o.Sample
	u.SampleInitial = o.SampleInitial
	u.SampleInterval = o.SampleInterval
	u.ItemSampleSeconds = o.ItemSampleSeconds
	u.ItemSampleInitial = o.ItemSampleInitial
	u.ItemSampleInterval = o.ItemSampleInterval
	u.ComponentLevels = o.ComponentLevels
	callbacks := u.onUpdatedCallbacks
	u.lock.Unlock()

	for _, cb := range callbacks {
		if err := cb(u.Config); err != nil {
			return err
		}
	}
	return nil
}

func (c *Config) AddUpdateObserver(cb ConfigObserver) {
	u := &uConfig
	u.lock.Lock()
	defer u.lock.Unlock()
	u.onUpdatedCallbacks = append(u.onUpdatedCallbacks, cb)
}
