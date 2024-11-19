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
	"errors"
	"time"

	"github.com/livekit/protocol/livekit"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	ErrInvalidMetricLabel           = errors.New("invalid metric label")
	ErrInvalidTimeSeriesMetricIndex = errors.New("invalid time series metric index")
)

type MetricsBatchBuilder struct {
	*livekit.MetricsBatch

	stringData map[string]uint32
}

func NewMetricsBatchBuilder(at time.Time, normalizedAt time.Time) *MetricsBatchBuilder {
	return &MetricsBatchBuilder{
		MetricsBatch: &livekit.MetricsBatch{
			TimestampMs:         at.UnixMilli(),
			NormalizedTimestamp: timestamppb.New(normalizedAt),
		},
		stringData: make(map[string]uint32),
	}
}

func (m *MetricsBatchBuilder) ToProto() *livekit.MetricsBatch {
	return m.MetricsBatch
}

type MetricSample struct {
	At           time.Time
	NormalizedAt time.Time
	Value        float32
}

type TimeSeriesMetric struct {
	MetricLabel         livekit.MetricLabel
	CustomMetricLabel   string
	ParticipantIdentity livekit.ParticipantIdentity
	TrackID             livekit.TrackID
	Samples             []MetricSample
	Rid                 string
}

func (m *MetricsBatchBuilder) AddTimeSeriesMetric(tsm TimeSeriesMetric) (int, error) {
	ptsm := &livekit.TimeSeriesMetric{}

	if tsm.CustomMetricLabel != "" {
		ptsm.Label = m.getStrDataIndex(tsm.CustomMetricLabel)
	} else {
		if tsm.MetricLabel >= livekit.MetricLabel_METRIC_LABEL_PREDEFINED_MAX_VALUE {
			return 0, ErrInvalidMetricLabel
		}
		ptsm.Label = uint32(tsm.MetricLabel)
	}

	if tsm.ParticipantIdentity != "" {
		ptsm.ParticipantIdentity = m.getStrDataIndex(string(tsm.ParticipantIdentity))
	}

	if tsm.TrackID != "" {
		ptsm.TrackSid = m.getStrDataIndex(string(tsm.TrackID))
	}

	for _, sample := range tsm.Samples {
		ptsm.Samples = append(ptsm.Samples, &livekit.MetricSample{
			TimestampMs:         sample.At.UnixMilli(),
			NormalizedTimestamp: timestamppb.New(sample.NormalizedAt),
			Value:               sample.Value,
		})
	}

	if tsm.Rid != "" {
		ptsm.Rid = m.getStrDataIndex(tsm.Rid)
	}

	m.MetricsBatch.TimeSeries = append(m.MetricsBatch.TimeSeries, ptsm)
	return len(m.MetricsBatch.TimeSeries) - 1, nil
}

func (m *MetricsBatchBuilder) AddMetricSamplesToTimeSeriesMetric(timeSeriesMetricIdx int, samples []MetricSample) error {
	if timeSeriesMetricIdx < 0 || timeSeriesMetricIdx >= len(m.MetricsBatch.TimeSeries) {
		return ErrInvalidTimeSeriesMetricIndex
	}

	ptsm := m.MetricsBatch.TimeSeries[timeSeriesMetricIdx]
	for _, sample := range samples {
		ptsm.Samples = append(ptsm.Samples, &livekit.MetricSample{
			TimestampMs:         sample.At.UnixMilli(),
			NormalizedTimestamp: timestamppb.New(sample.NormalizedAt),
			Value:               sample.Value,
		})
	}

	return nil
}

type EventMetric struct {
	MetricLabel         livekit.MetricLabel
	CustomMetricLabel   string
	ParticipantIdentity livekit.ParticipantIdentity
	TrackID             livekit.TrackID
	StartedAt           time.Time
	EndedAt             time.Time
	NormalizedStartedAt time.Time
	NormalizedEndedAt   time.Time
	Metadata            string
	Rid                 string
}

func (m *MetricsBatchBuilder) AddEventMetric(em EventMetric) error {
	pem := &livekit.EventMetric{}

	if em.CustomMetricLabel != "" {
		pem.Label = m.getStrDataIndex(em.CustomMetricLabel)
	} else {
		if em.MetricLabel >= livekit.MetricLabel_METRIC_LABEL_PREDEFINED_MAX_VALUE {
			return ErrInvalidMetricLabel
		}
		pem.Label = uint32(em.MetricLabel)
	}

	if em.ParticipantIdentity != "" {
		pem.ParticipantIdentity = m.getStrDataIndex(string(em.ParticipantIdentity))
	}

	if em.TrackID != "" {
		pem.TrackSid = m.getStrDataIndex(string(em.TrackID))
	}

	pem.StartTimestampMs = em.StartedAt.UnixMilli()
	if !em.EndedAt.IsZero() {
		endTimestampMs := em.EndedAt.UnixMilli()
		pem.EndTimestampMs = &endTimestampMs
	}

	pem.NormalizedStartTimestamp = timestamppb.New(em.NormalizedStartedAt)
	if !em.NormalizedEndedAt.IsZero() {
		pem.NormalizedEndTimestamp = timestamppb.New(em.NormalizedEndedAt)
	}

	pem.Metadata = em.Metadata

	if em.Rid != "" {
		pem.Rid = m.getStrDataIndex(em.Rid)
	}

	m.MetricsBatch.Events = append(m.MetricsBatch.Events, pem)
	return nil
}

func (m *MetricsBatchBuilder) getStrDataIndex(s string) uint32 {
	idx, ok := m.stringData[s]
	if !ok {
		m.MetricsBatch.StrData = append(m.MetricsBatch.StrData, s)
		idx = uint32(int(livekit.MetricLabel_METRIC_LABEL_PREDEFINED_MAX_VALUE) + len(m.MetricsBatch.StrData) - 1)
		m.stringData[s] = idx
	}
	return idx
}
