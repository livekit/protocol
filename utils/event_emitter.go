package utils

import (
	"container/list"
	"sync"

	"golang.org/x/exp/maps"

	"github.com/livekit/protocol/logger"
)

const defaultQueueSize = 16

type EventEmitterParams struct {
	QueueSize int
	Logger    logger.Logger
}

type EventEmitter[K comparable, V any] struct {
	params    EventEmitterParams
	mu        sync.RWMutex
	observers map[K]*list.List
}

func NewEventEmitter[K comparable, V any](params EventEmitterParams) *EventEmitter[K, V] {
	return &EventEmitter[K, V]{
		params:    params,
		observers: map[K]*list.List{},
	}
}

func NewDefaultEventEmitter[K comparable, V any]() *EventEmitter[K, V] {
	return NewEventEmitter[K, V](EventEmitterParams{
		QueueSize: defaultQueueSize,
		Logger:    logger.GetLogger(),
	})
}

func (e *EventEmitter[K, V]) Emit(k K, v V) {
	e.mu.RLock()
	defer e.mu.RUnlock()

	l, ok := e.observers[k]
	if !ok {
		return
	}

	for le := l.Front(); le != nil; le = le.Next() {
		le.Value.(EventObserver[V]).emit(v)
	}
}

func (e *EventEmitter[K, V]) Observe(k K) EventObserver[V] {
	o := EventObserver[V]{
		logger: e.params.Logger,
		ch:     make(chan V, e.params.QueueSize),
	}

	e.mu.Lock()
	l, ok := e.observers[k]
	if !ok {
		l = list.New()
		e.observers[k] = l
	}
	le := l.PushBack(o)
	e.mu.Unlock()

	o.stop = func() { e.stopObserving(k, le) }

	return o
}

func (e *EventEmitter[K, V]) ObservedKeys() []K {
	e.mu.Lock()
	defer e.mu.Unlock()
	return maps.Keys(e.observers)
}

func (e *EventEmitter[K, V]) stopObserving(k K, le *list.Element) {
	e.mu.Lock()
	defer e.mu.Unlock()

	l, ok := e.observers[k]
	if !ok {
		return
	}

	l.Remove(le)
	if l.Len() == 0 {
		delete(e.observers, k)
	}
}

type EventObserver[V any] struct {
	logger logger.Logger
	stop   func()
	ch     chan V
}

func NewEventObserver[V any](stopFunc func()) (EventObserver[V], func(v V)) {
	o := EventObserver[V]{
		logger: logger.GetLogger(),
		stop:   stopFunc,
		ch:     make(chan V, defaultQueueSize),
	}
	return o, o.emit
}

func (o EventObserver[V]) emit(v V) {
	select {
	case o.ch <- v:
	default:
		o.logger.Warnw("could not add event to observer queue", nil)
	}
}

func (o EventObserver[V]) Stop() {
	o.stop()
}

func (o EventObserver[V]) Events() <-chan V {
	return o.ch
}
