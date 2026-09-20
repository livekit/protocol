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

package configutil

// Derive narrows an observable onto a subtree of itself, so a package can
// define the config it owns, have the app embed it, and observe just that
// subtree without importing the app's config package.
//
// project selects; it must not compute. It runs on every Load, so it has to be
// cheap, and it must return a pointer into its argument rather than a freshly
// built value — callers (NewAtomicPointer among them) compare what Load
// returns by identity. For a projection that has to build something, use the
// NewAtomic* helpers instead: they cache the result and recompute it on reload.
//
// The result holds no state and needs no cleanup: it subscribes to src only
// while something is subscribed to it, and forwards src's emit semantics
// unchanged.
func Derive[Src, Dst any](src Observable[Src], project func(*Src) *Dst) Observable[Dst] {
	return derived[Src, Dst]{src: src, project: project}
}

type derived[Src, Dst any] struct {
	src     Observable[Src]
	project func(*Src) *Dst
}

func (d derived[Src, Dst]) Observe(cb func(*Dst)) func() {
	return d.src.Observe(func(c *Src) { cb(d.project(c)) })
}

func (d derived[Src, Dst]) Load() *Dst {
	return d.project(d.src.Load())
}
