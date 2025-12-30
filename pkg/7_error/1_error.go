package error

import (
	"errors"
	"fmt"
)

// define some common errors
var (
	ErrTimeout = errors.New("timeout")
)

// wrap
func Wrap(message string, err error) error {
	return fmt.Errorf("%v: %w", message, err)
}

// unwrap
func Unwrap(err error) error {
	return errors.Unwrap(err)
}

// errors.Is
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// errors.As
func As(err error, target any) bool {
	return errors.As(err, target)
}

// a self-defined error
type MyError struct {
	Message string `json:"message"`
	Err     error  `json:"inner_error,omitempty"`
}

func New(msg string, err error) *MyError {
	return &MyError{
		Message: msg,
		Err:     err,
	}
}

// error is an interface, so we just need to implement Error() func
func (e *MyError) Error() string {
	return e.Message + ": " + e.Err.Error()
}

func (e *MyError) Unwrap() error {
	return e.Err
}
