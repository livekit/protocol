// Copyright 2024 LiveKit, Inc.
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

package configutil

import (
	"math"
	"sync/atomic"
	"time"
)

type ConfigObserver[T any] interface {
	Observe(cb func(*T)) func()
	Load() *T
}

// AtomicFloat32 stores a float32 in a sync/atomic.Uint32 via its IEEE-754 bit
// representation. Exists because sync/atomic has no Float32.
type AtomicFloat32 struct {
	v atomic.Uint32
}

func (a *AtomicFloat32) Load() float32   { return math.Float32frombits(a.v.Load()) }
func (a *AtomicFloat32) Store(v float32) { a.v.Store(math.Float32bits(v)) }
func (a *AtomicFloat32) Swap(v float32) float32 {
	return math.Float32frombits(a.v.Swap(math.Float32bits(v)))
}

// AtomicFloat64 stores a float64 in a sync/atomic.Uint64 via its IEEE-754 bit
// representation. Exists because sync/atomic has no Float64.
type AtomicFloat64 struct {
	v atomic.Uint64
}

func (a *AtomicFloat64) Load() float64   { return math.Float64frombits(a.v.Load()) }
func (a *AtomicFloat64) Store(v float64) { a.v.Store(math.Float64bits(v)) }
func (a *AtomicFloat64) Swap(v float64) float64 {
	return math.Float64frombits(a.v.Swap(math.Float64bits(v)))
}

// AtomicDuration stores a time.Duration in a sync/atomic.Int64 (Duration is int64).
type AtomicDuration struct {
	v atomic.Int64
}

func (a *AtomicDuration) Load() time.Duration   { return time.Duration(a.v.Load()) }
func (a *AtomicDuration) Store(v time.Duration) { a.v.Store(int64(v)) }
func (a *AtomicDuration) Swap(v time.Duration) time.Duration {
	return time.Duration(a.v.Swap(int64(v)))
}

// AtomicString stores a string atomically. Backed by atomic.Pointer to a copy,
// since strings can vary in length.
type AtomicString struct {
	v atomic.Pointer[string]
}

func (a *AtomicString) Load() string {
	if p := a.v.Load(); p != nil {
		return *p
	}
	return ""
}
func (a *AtomicString) Store(v string) { a.v.Store(&v) }

// AtomicTime stores a time.Time atomically via atomic.Pointer.
type AtomicTime struct {
	v atomic.Pointer[time.Time]
}

func (a *AtomicTime) Load() time.Time {
	if p := a.v.Load(); p != nil {
		return *p
	}
	return time.Time{}
}
func (a *AtomicTime) Store(v time.Time) { a.v.Store(&v) }

func NewAtomicBool[C any](config ConfigObserver[C], accessor func(*C) bool) *atomic.Bool {
	return newAtomic(func(v bool) *atomic.Bool { a := &atomic.Bool{}; a.Store(v); return a }, config, accessor)
}

func NewAtomicDuration[C any](config ConfigObserver[C], accessor func(*C) time.Duration) *AtomicDuration {
	return newAtomic(func(v time.Duration) *AtomicDuration { a := &AtomicDuration{}; a.Store(v); return a }, config, accessor)
}

func NewAtomicFloat32[C any](config ConfigObserver[C], accessor func(*C) float32) *AtomicFloat32 {
	return newAtomic(func(v float32) *AtomicFloat32 { a := &AtomicFloat32{}; a.Store(v); return a }, config, accessor)
}

func NewAtomicFloat64[C any](config ConfigObserver[C], accessor func(*C) float64) *AtomicFloat64 {
	return newAtomic(func(v float64) *AtomicFloat64 { a := &AtomicFloat64{}; a.Store(v); return a }, config, accessor)
}

func NewAtomicInt32[C any](config ConfigObserver[C], accessor func(*C) int32) *atomic.Int32 {
	return newAtomic(func(v int32) *atomic.Int32 { a := &atomic.Int32{}; a.Store(v); return a }, config, accessor)
}

func NewAtomicInt64[C any](config ConfigObserver[C], accessor func(*C) int64) *atomic.Int64 {
	return newAtomic(func(v int64) *atomic.Int64 { a := &atomic.Int64{}; a.Store(v); return a }, config, accessor)
}

func NewAtomicString[C any](config ConfigObserver[C], accessor func(*C) string) *AtomicString {
	return newAtomic(func(v string) *AtomicString { a := &AtomicString{}; a.Store(v); return a }, config, accessor)
}

func NewAtomicTime[C any](config ConfigObserver[C], accessor func(*C) time.Time) *AtomicTime {
	return newAtomic(func(v time.Time) *AtomicTime { a := &AtomicTime{}; a.Store(v); return a }, config, accessor)
}

func NewAtomicUint32[C any](config ConfigObserver[C], accessor func(*C) uint32) *atomic.Uint32 {
	return newAtomic(func(v uint32) *atomic.Uint32 { a := &atomic.Uint32{}; a.Store(v); return a }, config, accessor)
}

func NewAtomicUint64[C any](config ConfigObserver[C], accessor func(*C) uint64) *atomic.Uint64 {
	return newAtomic(func(v uint64) *atomic.Uint64 { a := &atomic.Uint64{}; a.Store(v); return a }, config, accessor)
}

func NewAtomicPointer[T, C any](config ConfigObserver[C], accessor func(*C) *T) *atomic.Pointer[T] {
	return newAtomic(func(v *T) *atomic.Pointer[T] { a := &atomic.Pointer[T]{}; a.Store(v); return a }, config, accessor)
}

type AtomicValue[T any] struct {
	v atomic.Value
}

func newAtomicValue[T any](_ T) *AtomicValue[T] {
	return &AtomicValue[T]{}
}

func (a *AtomicValue[T]) Store(v T) {
	a.v.Store(v)
}

func (a *AtomicValue[T]) Load() T {
	return a.v.Load().(T)
}

func NewAtomicValue[T, C any](config ConfigObserver[C], accessor func(*C) T) *AtomicValue[T] {
	return newAtomic(newAtomicValue[T], config, accessor)
}

func newAtomic[T, C any, A interface{ Store(T) }](ctor func(T) A, config ConfigObserver[C], accessor func(*C) T) A {
	var zero T
	v := ctor(zero)
	config.Observe(func(cc *C) { v.Store(accessor(cc)) })
	v.Store(accessor(config.Load()))
	return v
}
