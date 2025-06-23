package async

import (
	"context"
	"fmt"
	"sync"
)

// --------------------------------------------------------------------------
// WORKER_POOL: WORKER MANAGER
// --------------------------------------------------------------------------
type workerPoolLite struct {
	service     string
	totalWorker int
	tasksChan   chan *taskLite

	context context.Context
	wg      sync.WaitGroup

	recoveryHandlerFunc func(string) error
}

func NewWorkerPoolLite(ctx context.Context, service string, worker int) *workerPoolLite {
	return &workerPoolLite{
		context:     ctx,
		service:     fmt.Sprintf("%s: %s", ErrWorkerPoolLite, service),
		totalWorker: worker,

		tasksChan: make(chan *taskLite),

		recoveryHandlerFunc: Recover,
	}
}

// custom recovery handler function hook
func (p *workerPoolLite) RecoveryHanlderFunc(fn func(string) error) *workerPoolLite {
	p.recoveryHandlerFunc = fn
	return p
}

func (p *workerPoolLite) Run() *workerPoolLite {
	for i := 0; i < p.totalWorker; i++ {
		go p.work()
	}

	return p
}

func (p *workerPoolLite) Dispatch(fn func(ctx context.Context)) {
	defer p.recoveryHandlerFunc(p.service)

	p.wg.Add(1)
	p.tasksChan <- newTaskLite(fn)
}

func (p *workerPoolLite) Wait() error {
	p.wg.Wait()
	close(p.tasksChan)
	return nil
}

func (p *workerPoolLite) work() {
	defer p.recoveryHandlerFunc(p.service)

	for task := range p.tasksChan {
		select {
		case <-p.context.Done():
			return
		default:
			task.RecoveryHanlderFunc(p.recoveryHandlerFunc).Run(p.context, &p.wg)
		}
	}
}

// --------------------------------------------------------------------------
// WORKER_POOL: TASKS MANAGER
// --------------------------------------------------------------------------
type taskLite struct {
	service string

	fn                  func(ctx context.Context)
	recoveryHandlerFunc func(string) error
}

func newTaskLite(fn func(ctx context.Context)) *taskLite {
	return &taskLite{
		fn:                  fn,
		recoveryHandlerFunc: Recover,
	}
}

// custom recovery handler function hook
func (t *taskLite) RecoveryHanlderFunc(fn func(string) error) *taskLite {
	t.recoveryHandlerFunc = fn
	return t
}

func (t *taskLite) Run(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	defer t.recoveryHandlerFunc(t.service)

	t.fn(ctx)
}
