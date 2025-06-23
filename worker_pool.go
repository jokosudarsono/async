package async

import (
	"context"
	"fmt"
	"sync"
)

// --------------------------------------------------------------------------
// WORKER_POOL: WORKER MANAGER
// --------------------------------------------------------------------------
type workerPool[V comparable] struct {
	service     string
	totalWorker int
	tasksChan   chan *task[V]

	context context.Context
	wg      sync.WaitGroup

	recoveryHandlerFunc func(string) error
}

func NewWorkerPool[data comparable](ctx context.Context, service string, worker int) *workerPool[data] {
	return &workerPool[data]{
		context:     ctx,
		service:     fmt.Sprintf("%s: %s", ErrWorkerPool, service),
		totalWorker: worker,

		tasksChan: make(chan *task[data]),

		recoveryHandlerFunc: Recover,
	}
}

// custom recovery handler function hook
func (p *workerPool[V]) RecoveryHanlderFunc(fn func(string) error) *workerPool[V] {
	p.recoveryHandlerFunc = fn
	return p
}

func (p *workerPool[V]) Run() *workerPool[V] {
	for i := 0; i < p.totalWorker; i++ {
		go p.work()
	}

	return p
}

func (p *workerPool[V]) Dispatch(data V, fn func(ctx context.Context, data V)) {
	p.wg.Add(1)
	p.tasksChan <- newTask(data, fn)
}

func (p *workerPool[V]) Wait() error {
	p.wg.Wait()
	close(p.tasksChan)
	return nil
}

func (p *workerPool[V]) work() {
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
type task[V comparable] struct {
	service string
	data    V

	fn                  func(ctx context.Context, data V)
	recoveryHandlerFunc func(string) error
}

func newTask[V comparable](data V, fn func(ctx context.Context, data V)) *task[V] {
	return &task[V]{
		fn:                  fn,
		data:                data,
		recoveryHandlerFunc: Recover,
	}
}

// custom recovery handler function hook
func (t *task[V]) RecoveryHanlderFunc(fn func(string) error) *task[V] {
	t.recoveryHandlerFunc = fn
	return t
}

func (t *task[V]) Run(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	defer t.recoveryHandlerFunc(t.service)

	t.fn(ctx, t.data)
}
