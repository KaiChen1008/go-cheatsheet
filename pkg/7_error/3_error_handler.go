package error

import (
	"context"
	"errors"
	"time"
)

type ErrorHandler struct {
	maxAttempts int
	baseWait    time.Duration
	maxWait     time.Duration
}

func NewErrorHandler(maxAttempts int) *ErrorHandler {
	return &ErrorHandler{
		maxAttempts: maxAttempts,
		baseWait:    time.Second,
		maxWait:     time.Minute,
	}
}

func (eh *ErrorHandler) RunWithRetry(ctx context.Context, inFn func() error) error {
	var err error
	if err = inFn(); !eh.shouldRetry(err) {
		return err
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for range eh.maxAttempts - 1 {
		select {
		case <-ticker.C:
			if err = inFn(); !eh.shouldRetry(err) {
				return err
			}

			// reset wait
			waitTime := eh.nextWaitTime(eh.maxAttempts)
			ticker.Reset(waitTime)

		case <-ctx.Done():
			return ctx.Err()
		}
	}
	return err
}

func (eh *ErrorHandler) shouldRetry(err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return false
	}
	if errors.Is(err, ErrTimeout) {
		return false
	}
	return true
}

func (eh *ErrorHandler) nextWaitTime(attempts int) time.Duration {
	wait := eh.baseWait << (time.Duration(attempts - 1))
	if wait > eh.maxWait {
		return eh.maxWait
	}
	return wait
}
