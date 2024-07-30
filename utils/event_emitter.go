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
	"sync"

	"golang.org/x/exp/maps"

	"github.com/livekit/protocol/logger"
)

const DefaultEventQueueSize = 16

type EventEmitterParams struct {
	QueueSize int
	Blocking  bool
	Logger    logger.Logger
}

func DefaultEventEmitterParams() EventEmitterParams {
	return EventEmitterParams{
		QueueSize: DefaultEventQueueSize,
		Logger:    logger.GetLogger(),
	}
}

type EventEmitter[K comparable, V any] struct {
	params    EventEmitterParams
	mu        sync.RWMutex
	observers map[K]*EventObserverList[V]
}

func NewEventEmitter[K comparable, V any](params EventEmitterParams) *EventEmitter[K, V] {
	return &EventEmitter[K, V]{
		params:    params,
		observers: map[K]*EventObserverList[V]{},
	}
}

func NewDefaultEventEmitter[K comparable, V any]() *EventEmitter[K, V] {
	return NewEventEmitter[K, V](DefaultEventEmitterParams())
}

func (e *EventEmitter[K, V]) Emit(k K, v V) {
	e.mu.RLock()
	l, ok := e.observers[k]
	e.mu.RUnlock()
	if !ok {
		return
	}

	l.Emit(v)
}

type eventEmitterObserver[K comparable, V any] struct {
	e *EventEmitter[K, V]
	k K
	EventObserver[V]
}

func (o *eventEmitterObserver[K, V]) Stop() {
	o.EventObserver.Stop()
	o.e.cleanUpObserverList(o.k)
}

func (e *EventEmitter[K, V]) Observe(k K) EventObserver[V] {
	e.mu.Lock()
	l, ok := e.observers[k]
	if !ok {
		l = NewEventObserverList[V](e.params)
		e.observers[k] = l
	}
	o := l.Observe()
	e.mu.Unlock()

	return &eventEmitterObserver[K, V]{e, k, o}
}

func (e *EventEmitter[K, V]) ObservedKeys() []K {
	e.mu.Lock()
	defer e.mu.Unlock()
	return maps.Keys(e.observers)
}

func (e *EventEmitter[K, V]) cleanUpObserverList(k K) {
	e.mu.Lock()
	defer e.mu.Unlock()

	l, ok := e.observers[k]
	if ok && l.Len() == 0 {
		delete(e.observers, k)
	}
}

type EventObserverList[V any] struct {
	params    EventEmitterParams
	mu        sync.RWMutex
	observers []*eventObserverListObserver[V]
}

func NewEventObserverList[V any](params EventEmitterParams) *EventObserverList[V] {
	return &EventObserverList[V]{
		params: params,
	}
}

func NewDefaultEventObserverList[V any]() *EventObserverList[V] {
	return NewEventObserverList[V](DefaultEventEmitterParams())
}

func (l *EventObserverList[V]) Len() int {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return len(l.observers)
}

type eventObserverListObserver[V any] struct {
	l *EventObserverList[V]
	EventObserver[V]
	index int
}

func (o *eventObserverListObserver[V]) Stop() {
	o.l.stopObserving(o)
	o.EventObserver.Stop()
}

func (l *EventObserverList[V]) Observe() EventObserver[V] {
	o := &eventObserverListObserver[V]{l: l}

	if l.params.Blocking {
		o.EventObserver = &blockingEventObserver[V]{
			done: make(chan struct{}),
			ch:   make(chan V, l.params.QueueSize),
		}
	} else {
		o.EventObserver = &nonblockingEventObserver[V]{
			logger: l.params.Logger,
			ch:     make(chan V, l.params.QueueSize),
		}
	}

	l.mu.Lock()
	o.index = len(l.observers)
	l.observers = append(l.observers, o)
	l.mu.Unlock()

	return o
}

func (l *EventObserverList[V]) Emit(v V) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	for _, o := range l.observers {
		o.emit(v)
	}
}

func (l *EventObserverList[V]) stopObserving(o *eventObserverListObserver[V]) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.observers[o.index] = l.observers[len(l.observers)-1]
	l.observers[o.index].index = o.index
	l.observers[len(l.observers)-1] = nil
	l.observers = l.observers[:len(l.observers)-1]
}

type EventObserver[V any] interface {
	emit(v V)
	Stop()
	Events() <-chan V
}

type eventObserver[V any] struct {
	stopFunc func()
	EventObserver[V]
}

func (o *eventObserver[V]) Stop() {
	o.stopFunc()
	o.EventObserver.Stop()
}

func NewEventObserver[V any](stopFunc func()) (EventObserver[V], func(v V)) {
	o := &nonblockingEventObserver[V]{
		logger: logger.GetLogger(),
		ch:     make(chan V, DefaultEventQueueSize),
	}
	return &eventObserver[V]{stopFunc, o}, o.emit
}

type nonblockingEventObserver[V any] struct {
	logger logger.Logger
	ch     chan V
}

func (o *nonblockingEventObserver[V]) emit(v V) {
	select {
	case o.ch <- v:
	default:
		o.logger.Warnw("could not add event to observer queue", nil)
	}
}

func (o *nonblockingEventObserver[V]) Stop() {}

func (o *nonblockingEventObserver[V]) Events() <-chan V {
	return o.ch
}

type blockingEventObserver[V any] struct {
	done chan struct{}
	ch   chan V
}

func (o *blockingEventObserver[V]) emit(v V) {
	select {
	case o.ch <- v:
	case <-o.done:
	}
}

func (o *blockingEventObserver[V]) Stop() {
	close(o.done)
}

func (o *blockingEventObserver[V]) Events() <-chan V {
	return o.ch
}
