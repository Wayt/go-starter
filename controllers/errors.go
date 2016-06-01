package controllers

import (
	"errors"
	"strings"
)

var (
	ErrInternalError  = errors.New("internal error")
	ErrDuplicateEntry = errors.New("duplicate entry")
	ErrNotFound       = errors.New("not found")
)

// ControllerError represente a controller error
type ControllerError struct {
	ErrType error // Should be the error type
	Err     error // orignal error
}

// Error interface compliance
func (ce ControllerError) Error() string {
	return ce.ErrType.Error()
}

// Type error
func (ce ControllerError) Type() error {
	return ce.ErrType
}

// Origin is the original error
func (ce ControllerError) Origin() error {
	return ce.Err
}

func wrap(err error) error {

	switch {
	case strings.Contains(err.Error(), "Duplicate entry"):
		return ControllerError{
			ErrType: ErrDuplicateEntry,
			Err:     err,
		}
	case strings.Contains(err.Error(), "record not found"):
		return ControllerError{
			ErrType: ErrNotFound,
			Err:     err,
		}
	}

	return ControllerError{
		ErrType: ErrInternalError,
		Err:     err,
	}
}
