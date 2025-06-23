package async

import (
	"context"
	"fmt"
)

func Go(ctxt context.Context, service string, fn func(ctx context.Context)) {
	newCtx := CopyContext(ctxt)

	go func() {
		defer Recover(fmt.Sprintf("async-go: %s", service))

		newCtx, cancel := context.WithCancel(newCtx)
		defer cancel()

		fn(newCtx)
	}()
}
