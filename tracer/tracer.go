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

package tracer

import "context"

type Tracer interface {
	Start(ctx context.Context, spanName string, opts ...interface{}) (context.Context, Span)
}

type Span interface {
	RecordError(err error)
	End()
}

var tracer Tracer = &NoOpTracer{}

// Can be used for your own tracing (for example, with Lightstep)
func SetTracer(t Tracer) {
	tracer = t
}

func Start(ctx context.Context, spanName string, opts ...interface{}) (context.Context, Span) {
	return tracer.Start(ctx, spanName, opts...)
}

type NoOpTracer struct{}

func (t *NoOpTracer) Start(ctx context.Context, _ string, _ ...interface{}) (context.Context, Span) {
	return ctx, &NoOpSpan{}
}

type NoOpSpan struct{}

func (s *NoOpSpan) RecordError(_ error) {}

func (s *NoOpSpan) End() {}
