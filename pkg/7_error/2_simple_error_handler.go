package error

import (
	"context"
	"time"
)

type SimpleErrorHandler struct {
	maxAttampts int
}

func NewSimple(maxAttampts int) *SimpleErrorHandler {
	return &SimpleErrorHandler{
		maxAttampts: maxAttampts,
	}
}

func (eh *SimpleErrorHandler) RunWithRetry(ctx context.Context, inFn func() error) error {
	var err error
	if err = inFn(); err == nil {
		return nil
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for range eh.maxAttampts - 1 {
		select {
		case <-ticker.C:
			if err = inFn(); err == nil {
				return nil
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	return err
}
