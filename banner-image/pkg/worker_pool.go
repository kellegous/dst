package pkg

import "sync"

type WorkerPool struct {
	ch chan func()
	wg sync.WaitGroup
}

func (w *WorkerPool) Submit(fn func()) {
	w.ch <- fn
}

func (w *WorkerPool) Drain() {
	close(w.ch)
	w.wg.Wait()
}

func StartWorkers(n int) *WorkerPool {
	ch := make(chan func())

	wp := &WorkerPool{
		ch: ch,
	}

	for i := 0; i < n; i++ {
		go func() {
			defer wp.wg.Done()
			for fn := range ch {
				fn()
			}
		}()
	}

	wp.wg.Add(n)
	return wp
}
