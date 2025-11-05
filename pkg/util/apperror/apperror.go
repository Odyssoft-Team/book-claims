package apperror

import (
	"fmt"
	"net/http"
	"runtime"
	"runtime/debug"
)

type AppError struct {
	Code       int
	Message    string
	Underlying error
	File       string
	Line       int
	StackTrace string
}

func (e *AppError) Error() string {
	if e.Underlying != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Underlying)
	}
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Underlying
}

func captureCaller(skip int) (string, int) {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "unknown", 0
	}
	return file, line
}

// Helpers

func NewBadRequest(msg string) *AppError {
	file, line := captureCaller(2)
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: msg,
		File:    file,
		Line:    line,
	}
}

func NewInternalError(msg string, err error) *AppError {
	file, line := captureCaller(2)
	return &AppError{
		Code:       http.StatusInternalServerError,
		Message:    msg,
		Underlying: err,
		File:       file,
		Line:       line,
		StackTrace: string(debug.Stack()),
	}
}

func NewNotFoundError(msg string) *AppError {
	file, line := captureCaller(2)
	return &AppError{
		Code:    http.StatusNotFound,
		Message: msg,
		File:    file,
		Line:    line,
	}
}

func NewUnauthorized(msg string) *AppError {
	file, line := captureCaller(2)
	return &AppError{
		Code:    http.StatusUnauthorized,
		Message: msg,
		File:    file,
		Line:    line,
	}
}
