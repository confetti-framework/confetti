package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
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
		return http.StatusText(e.HttpStatus)
	}
	return e.Message
}

func (e UserError) GetHttpStatus() int {
	if e.HttpStatus != 0 {
		return e.HttpStatus
	}
	return http.StatusBadRequest
}

type SystemError struct {
	HttpStatus int
	Code       string
}

func NewSystemError(err error, code string) error {
	return errors.Join(err, SystemError{Code: code}, err)
}

func (e SystemError) Error() string {
	if e.Code == "" {
		return http.StatusText(e.HttpStatus)
	}
	return e.Code
}

func (e SystemError) GetHttpStatus() int {
	if e.HttpStatus != 0 {
		return e.HttpStatus
	}
	return http.StatusInternalServerError
}

// apiErrorHandler converts the error to a reponse. Feel free to modify this to your needs.
func apiErrorHandler(writer http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	publicMessage := ""
	internalMessage := ""
	reference := fmt.Sprintf("%d", time.Now().UnixMilli())
	var systemError SystemError
	var userError UserError
	if errors.As(err, &systemError) {
		// Handle system error
		status = systemError.GetHttpStatus()
		reference = fmt.Sprintf("%s/%s", systemError.Code, reference)
		publicMessage = "internal server error: " + reference
		internalMessage = err.Error()
	} else if errors.As(err, &userError) {
		// Handle user error
		status = userError.GetHttpStatus()
		publicMessage = userError.Message
	} else {
		// Handle unknown error as a system error
		// Please join the error with with a user error:
		// errors.Join(err, handler.UserError{HttpStatus: http.StatusUnauthorized})
		status = http.StatusInternalServerError
		publicMessage = "unhandled internal server error"
		internalMessage = err.Error()
	}

	// Set the response header to application/json
	writer.Header().Set("Content-Type", "application/json")
	// Write the status code as 500 Internal Server Error
	writer.WriteHeader(status)
	// Create a JSON response with the error publicMessage
	_ = json.NewEncoder(writer).Encode(map[string]string{
		"reference": reference,
		"message":   publicMessage,
	})

	// If it is a user error, we don't want to log that
	if errors.As(err, &userError) {
		return
	}
	_, err = fmt.Printf("Error report: %s, message: %s", reference, internalMessage)
	if err != nil {
		log.Panicln("can't print error: " + err.Error())
	}
}
