package worker_pool

import (
	"context"
	"sync"
)

type WorkerFunc func(ctx context.Context)

type WorkerPool struct {
	wg        sync.WaitGroup
	ctx       context.Context
	cancel    context.CancelFunc
	jobs      chan WorkerFunc
	maxWorker int
}

func NewWorkerPool(ctx context.Context, maxWorker, maxCacheJob int) *WorkerPool {
	cancel, cancelFunc := context.WithCancel(ctx)

	wp := &WorkerPool{
		ctx:       cancel,
		cancel:    cancelFunc,
		jobs:      make(chan WorkerFunc, maxCacheJob),
		maxWorker: maxWorker,
	}

	wp.startWorker()

	return wp
}

func (w *WorkerPool) startWorker() {
	for i := 0; i < w.maxWorker; i++ {
		w.wg.Add(1)
		go w.processWorker()
	}
}

func (w *WorkerPool) processWorker() {
	defer w.wg.Done()
	for {
		select {
		case <-w.ctx.Done():
			return
		case job, ok := <-w.jobs:
			if !ok {
				return
			}

			job(w.ctx)
		}
	}
}

func (w *WorkerPool) Add(f WorkerFunc) {
	w.jobs <- f
}

func (w *WorkerPool) TryAdd(f WorkerFunc) bool {
	select {
	case w.jobs <- f:
		return true
	default:
		return false
	}
}

func (w *WorkerPool) Close() {
	w.cancel()
}

func (w *WorkerPool) WaitAndClose() {
	close(w.jobs)
	w.wg.Wait()
	w.Close()
}
