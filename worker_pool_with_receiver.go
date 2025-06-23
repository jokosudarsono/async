package async

import (
	"context"
	"fmt"
	"sync"
)

// --------------------------------------------------------------------------
// WORKER_POOL: WORKER MANAGER
// --------------------------------------------------------------------------
type workerPoolWithReceiver[V, K comparable] struct {
	service     string
	totalWorker int

	tasksChan chan *taskWithReceiver[V, K]
	receiver  chan K
	reader    chan []K

	context    context.Context
	wg         sync.WaitGroup
	receiverWg sync.WaitGroup

	recoveryHandlerFunc func(string) error
}

func NewWorkerPoolWithReceiver[input, receiver comparable](
	ctx context.Context,
	service string,
	worker,
	queueCapacity int,
) *workerPoolWithReceiver[input, receiver] {
	var (
		capacity = 25
	)

	if queueCapacity > capacity {
		capacity = queueCapacity
	}

	return &workerPoolWithReceiver[input, receiver]{
		context:     ctx,
		service:     fmt.Sprintf("%s: %s", ErrWorkerPoolWithReceiver, service),
		totalWorker: worker,

		tasksChan: make(chan *taskWithReceiver[input, receiver]),
		receiver:  make(chan receiver, capacity),
		reader:    make(chan []receiver, 1),

		recoveryHandlerFunc: Recover,
	}
}

// custom recovery handler function hook
func (p *workerPoolWithReceiver[V, K]) RecoveryHanlderFunc(fn func(string) error) *workerPoolWithReceiver[V, K] {
	p.recoveryHandlerFunc = fn
	return p
}

func (p *workerPoolWithReceiver[V, K]) Run() *workerPoolWithReceiver[V, K] {
	for i := 0; i < p.totalWorker; i++ {
		go p.work()
	}

	p.receiverWg.Add(1)

	go func() {
		defer p.receiverWg.Done()
		defer close(p.reader)

		defer p.recoveryHandlerFunc(p.service)

		var data = []K{}

		for rcv := range p.receiver {
			data = append(data, rcv)
		}

		p.reader <- data
	}()

	return p
}

func (p *workerPoolWithReceiver[V, K]) Dispatch(data V, fn func(ctx context.Context, data V, receiver chan<- K)) {
	p.wg.Add(1)
	p.tasksChan <- newTaskWithReceiver(data, fn)
}

func (p *workerPoolWithReceiver[V, K]) Result() []K {
	return <-p.reader
}

func (p *workerPoolWithReceiver[V, K]) Wait() error {
	p.waitTask()

	close(p.receiver)
	p.receiverWg.Wait()

	return nil
}

func (p *workerPoolWithReceiver[V, K]) waitTask() error {
	p.wg.Wait()
	close(p.tasksChan)

	return nil
}

func (p *workerPoolWithReceiver[V, K]) work() {
	defer p.recoveryHandlerFunc(p.service)

	for task := range p.tasksChan {
		select {
		case <-p.context.Done():
			return
		default:
			task.RecoveryHanlderFunc(p.recoveryHandlerFunc).Run(p.context, &p.wg, p.receiver)
		}
	}
}

// --------------------------------------------------------------------------
// WORKER_POOL: TASKS MANAGER
// --------------------------------------------------------------------------
type taskWithReceiver[V, K comparable] struct {
	service string
	data    V

	fn                  func(ctx context.Context, data V, receiver chan<- K)
	recoveryHandlerFunc func(string) error
}

func newTaskWithReceiver[V, K comparable](data V, fn func(ctx context.Context, data V, receiver chan<- K)) *taskWithReceiver[V, K] {
	return &taskWithReceiver[V, K]{
		fn:                  fn,
		data:                data,
		recoveryHandlerFunc: Recover,
	}
}

// custom recovery handler function hook
func (t *taskWithReceiver[V, K]) RecoveryHanlderFunc(fn func(string) error) *taskWithReceiver[V, K] {
	t.recoveryHandlerFunc = fn
	return t
}

func (t *taskWithReceiver[V, K]) Run(ctx context.Context, wg *sync.WaitGroup, receiver chan<- K) {
	defer wg.Done()
	defer t.recoveryHandlerFunc(t.service)

	t.fn(ctx, t.data, receiver)
}
