package utils

import (
	"errors"
)

type CustomError struct {
	StatusCode int
	Err        error
}

func (r *CustomError) Error() string {
	return r.Err.Error()
}

func CreateError(code int, message string) error {
	return &CustomError{
		StatusCode: code,
		Err:        errors.New(message),
	}
}
