package controllers

import (
	"fmt"
	"net/http"
)

type ApiError interface {
	error
	GetMessage() string
	GetStatusCode() int
}

type apiError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Err        string `json:"error"`
}

func (e apiError) Error() string {
	return fmt.Sprintf("message: %s. status: %d. error: %s", e.Message, e.StatusCode, e.Err)
}

func (e apiError) GetMessage() string {
	return e.Message
}

func (e apiError) GetStatusCode() int {
	return e.StatusCode
}

func NewApiError(message, err string, statusCode int) ApiError {
	return apiError{
		Message:    message,
		StatusCode: statusCode,
		Err:        err,
	}
}

func NewNotImplementedApiError(message string) ApiError {
	return apiError{
		Message:    message,
		StatusCode: http.StatusNotImplemented,
		Err:        "not_implemented",
	}
}

func NewBadRequestApiError(message string) ApiError {
	return apiError{
		Message:    message,
		StatusCode: http.StatusBadRequest,
		Err:        "bad_request",
	}
}

func NewNotFoundApiError(message string) ApiError {
	return apiError{
		Message:    message,
		StatusCode: http.StatusNotFound,
		Err:        "not_found",
	}
}
