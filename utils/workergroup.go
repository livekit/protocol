package utils

import "sync"

type WorkerGroup struct {
	wg sync.WaitGroup
}

func (w *WorkerGroup) Go(fn func()) {
	w.wg.Add(1)
	go func() {
		fn()
		w.wg.Done()
	}()
}

func (w *WorkerGroup) Wait() {
	w.wg.Wait()
}
