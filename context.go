package async

import "context"

type ContextKey string

func CopyContext(parent context.Context) context.Context {
	newCtx := parent // Preserve cancellation, timeouts, etc.

	keys := GetConfig().GetCopyContext()
	for _, key := range keys {
		if val := parent.Value(key); val != nil {
			newCtx = context.WithValue(newCtx, key, val)
		}
	}

	return newCtx
}
