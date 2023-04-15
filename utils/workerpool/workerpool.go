package workerpool

import (
	"hash/maphash"
	"sync"

	"github.com/gammazero/deque"
	"go.uber.org/atomic"
)

type task func()

type worker struct {
	tasks    *deque.Deque[task]
	taskLock sync.Mutex
	stopped  atomic.Bool
	drain    atomic.Bool
	notify   chan struct{}
	done     chan struct{}
}

func newWorker() *worker {
	w := &worker{
		tasks:  deque.New[task](),
		notify: make(chan struct{}, 10),
		done:   make(chan struct{}),
	}
	go w.process()
	return w
}

func (w *worker) submit(task task) {
	if w.stopped.Load() {
		return
	}
	w.taskLock.Lock()
	w.tasks.PushBack(task)
	n := w.tasks.Len()
	w.taskLock.Unlock()

	if n == 1 {
		w.notifyWorker()
	}
}

func (w *worker) stop() {
	w.stopped.Store(true)
	w.notifyWorker()

}

func (w *worker) stopDrain() {
	w.drain.Store(true)
	w.stop()
	<-w.done
}

func (w *worker) notifyWorker() {
	select {
	case w.notify <- struct{}{}:
	default:
	}
}

func (w *worker) notifyDone() {
	select {
	case w.done <- struct{}{}:
	default:
	}
}

func (w *worker) process() {
	for !w.stopped.Load() || w.drain.Load() {
		w.taskLock.Lock()
		if w.tasks.Len() == 0 {
			w.taskLock.Unlock()
			if w.drain.Load() {
				w.notifyDone()
				return
			}
			<-w.notify
			continue
		}
		t := w.tasks.PopFront()
		w.taskLock.Unlock()
		t()
	}
}

type OrderedWorkerPool struct {
	workers     []*worker
	hash        maphash.Hash
	workerCount int
	nilKeyCount atomic.Uint64
	lock        sync.Mutex
	stopped     atomic.Bool
}

func NewOrderedWorkerPool(workerCount int) *OrderedWorkerPool {
	if workerCount < 1 {
		workerCount = 1
	}

	wp := &OrderedWorkerPool{
		workers:     make([]*worker, workerCount),
		workerCount: workerCount,
	}

	for i := 0; i < workerCount; i++ {
		wp.workers[i] = newWorker()
	}

	return wp
}

func (p *OrderedWorkerPool) Submit(fn task, key string) {
	var keyHash uint64
	if key == "" {
		keyHash = p.nilKeyCount.Inc()
	} else {
		p.lock.Lock()
		p.hash.Reset()
		p.hash.WriteString(key)
		keyHash = p.hash.Sum64()
		p.lock.Unlock()
	}

	p.workers[keyHash%uint64(p.workerCount)].submit(fn)
}

func (p *OrderedWorkerPool) SubmitUnordered(fn task) {
	p.Submit(fn, "")
}

func (p *OrderedWorkerPool) Stop() {
	if p.stopped.Swap(true) {
		return
	}

	for _, w := range p.workers {
		w.stop()
	}
}

func (p *OrderedWorkerPool) StopDrain() {
	if p.stopped.Swap(true) {
		return
	}
	var wg sync.WaitGroup
	wg.Add(len(p.workers))

	for _, w := range p.workers {
		tmp := w
		waitFor := func() {
			defer wg.Done()
			tmp.stopDrain()
		}
		go waitFor()
	}
	wg.Wait()
}
