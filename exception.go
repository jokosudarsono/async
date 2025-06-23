package async

import "errors"

var (
	ErrPanic error = errors.New("[PANIC_RECOVERY]")

	ErrWorkerPool             error = errors.New("[WORKER_POOL]")
	ErrWorkerPoolLite         error = errors.New("[WORKER_POOL_LITE]")
	ErrWorkerPoolWithReceiver error = errors.New("[WORKER_POOL_WITH_RECEIVER]")
)
