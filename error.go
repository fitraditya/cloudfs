package cloudfs

import (
	"errors"
	"fmt"
)

// ErrNotSupported is returned when this operations is not supported.
var ErrNotSupported = errors.New("storage doesn't support this operation")

// ErrAlreadyOpened is returned when the file is already opened.
var ErrAlreadyOpened = errors.New("already opened")

// ErrInvalidSeek is returned when the seek operation is not doable.
var ErrInvalidSeek = errors.New("invalid seek offset")

type IOError struct {
	message string
}

func NewIOError(message string) error {
	return &IOError{
		message: message,
	}
}

func (e IOError) Error() string {
	return e.message
}

type TimeoutError struct {
	message string
}

func NewTimeoutError(path string) error {
	return &TimeoutError{
		message: fmt.Sprintf("cloudfs %s: lock waiting time exceeded", path),
	}
}

func (e TimeoutError) Error() string {
	return e.message
}

type ContextCanceled struct {
	message string
}

func NewContextCanceled(message string) error {
	return &ContextCanceled{
		message: message,
	}
}

func (e ContextCanceled) Error() string {
	return e.message
}

type ContextDone struct {
	message string
}

func NewContextDone(message string) error {
	return &ContextDone{
		message: message,
	}
}

func (e ContextDone) Error() string {
	return e.message
}
