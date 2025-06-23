// -------------------------------------------------------------
// AsyncGroup Module
//
// AsyncGroup module is tool to simplify complex goroutine process
// -------------------------------------------------------------

package async

import (
	"context"
	"sync"
)

type asyncGroup struct {
	ctxt   context.Context
	cancel func(error)

	wg sync.WaitGroup

	recoveryHandlerFunc func(string) error
	service             string
	err                 error
}

func AsyncGroup(ctx context.Context, service string) *asyncGroup {
	newCtx := CopyContext(ctx)
	cancellableCtx, cancel := context.WithCancelCause(newCtx)

	return &asyncGroup{
		cancel:              cancel,
		recoveryHandlerFunc: Recover,
		ctxt:                cancellableCtx,
	}
}

// custom recovery handler function hook
func (a *asyncGroup) RecoveryHanlderFunc(fn func(string) error) *asyncGroup {
	a.recoveryHandlerFunc = fn
	return a
}

func (a *asyncGroup) Wait() error {
	a.wg.Wait()

	if a.cancel != nil {
		a.cancel(a.err)
	}

	return a.err
}

// goroutine spawner
func (a *asyncGroup) Go(fn func(ctx context.Context)) {
	a.wg.Add(1)

	go func() {
		defer a.recoveryHandlerFunc(a.service)
		defer a.wg.Done()

		ctx, cancel := context.WithCancel(a.ctxt)
		defer cancel()

		fn(ctx)
	}()
}
