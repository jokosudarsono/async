package async

import (
	"fmt"
)

func Recover(service string) error {
	if err := recover(); err != nil {
		return fmt.Errorf("%s: %s: %v", ErrPanic, service, err)
	}

	return nil
}
