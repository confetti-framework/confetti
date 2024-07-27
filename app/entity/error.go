package entity

import (
	"errors"
	net "net/http"
)

type UserError struct {
	HttpStatus int
	Message    string
}

func NewUserError(message string, httpStatus int) error {
	return UserError{Message: message, HttpStatus: httpStatus}
}

func (e UserError) Error() string {
	if e.Message == "" {
		return net.StatusText(e.HttpStatus)
	}
	return e.Message
}

func (e UserError) GetHttpStatus() int {
	if e.HttpStatus != 0 {
		return e.HttpStatus
	}
	return net.StatusBadRequest
}

type SystemError struct {
	HttpStatus int
	Message    string
}

func NewSystemError(err error, code string) error {
	return errors.Join(err, SystemError{Message: code}, err)
}

func (e SystemError) Error() string {
	if e.Message == "" {
		return net.StatusText(e.HttpStatus)
	}
	return e.Message
}

func (e SystemError) GetHttpStatus() int {
	if e.HttpStatus != 0 {
		return e.HttpStatus
	}
	return net.StatusInternalServerError
}
