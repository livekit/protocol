// Copyright 2026 LiveKit, Inc.
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
	"runtime"
	"sync"
	"unsafe"
)

// Synchronized guards a value behind a mutex whose Lock returns the guarded
// value and a Holder handle; Unlock is a method on the handle, so a release
// always matches its acquisition, even when the lock is handed off between
// goroutines. The embedded lockTracker registers in the shared scan registry,
// so stuck-lock diagnostics (ScanTrackedLocks, PopulateHolderStacks) work
// like they do for Mutex. The zero value is ready to use and guards the zero
// value of T.
type Synchronized[T any] struct {
	base  syncBase
	value T
}

func NewSynchronized[T any](value T) *Synchronized[T] {
	return &Synchronized[T]{value: value}
}

// Lock acquires the mutex and returns the guarded value with the Holder that
// releases it. The value must not be used after Holder.Unlock.
func (s *Synchronized[T]) Lock() (*T, Holder) {
	b := &s.base
	enabled := lockTrackerEnabled
	if enabled {
		b.t.trackWait()
	}
	b.mu.Lock()
	if enabled {
		if !b.registered {
			b.registered = true
			b.t.stack = make([]uintptr, lockTrackerMaxStackDepth)
			registerSyncBase(s, &b.t)
		}
		b.t.trackLock()
	}
	return &s.value, Holder{base: b}
}

// Holder releases one acquisition of a Synchronized lock. Treat it as
// move-only: copies (e.g. sending it to another goroutine for handoff) share
// the acquisition, which must be unlocked exactly once across all copies.
type Holder struct {
	base *syncBase
}

func (h *Holder) Unlock() {
	b := h.base
	if b == nil {
		panic("utils.Holder: unlocked twice")
	}
	h.base = nil
	b.t.trackUnlock()
	b.mu.Unlock()
}

type syncBase struct {
	mu         sync.Mutex
	registered bool // guarded by mu
	t          lockTracker
}

// registerSyncBase adds the lock to the scan registry on first Lock. The
// finalizer goes on the owner — the tracker is embedded in its allocation —
// and captures only the slot index, since capturing the tracker would root
// the owner and leak it.
func registerSyncBase(owner any, t *lockTracker) {
	ref := weakRefs.add(unsafe.Pointer(t))
	runtime.SetFinalizer(owner, func(any) {
		weakRefs.remove(ref)
	})
}
