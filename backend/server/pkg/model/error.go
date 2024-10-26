package model

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	errFieldMissing       = errors.New("field missing")
	errFieldInvalidFormat = errors.New("invalid format")

	errConflict       = errors.New("conflict")
	errInternalServer = errors.New("internal server error")
)

type ApiError struct {
	statusCode int
	message    string
	err        error
}

func (e *ApiError) StatusCode() int {
	return e.statusCode
}

func (e *ApiError) Error() string {
	return e.message
}

func (e *ApiError) Is(target error) bool {
	return errors.Is(e.err, target)
}

func NewFieldMissingError(name string) *ApiError {
	return &ApiError{
		statusCode: http.StatusBadRequest,
		message:    fmt.Sprintf("%s: %s", errFieldMissing.Error(), name),
		err:        errFieldMissing,
	}
}

func NewFieldInvalidFormatError(detail string) *ApiError {
	return &ApiError{
		statusCode: http.StatusBadRequest,
		message:    fmt.Sprintf("%s: %s", errFieldInvalidFormat.Error(), detail),
		err:        errFieldInvalidFormat,
	}
}

func NewConflictError(detail string) *ApiError {
	return &ApiError{
		statusCode: http.StatusConflict,
		message:    fmt.Sprintf("%s: %s", errConflict.Error(), detail),
		err:        errConflict,
	}
}

func NewInternalServerError() *ApiError {
	return &ApiError{
		statusCode: http.StatusInternalServerError,
		message:    errInternalServer.Error(),
		err:        errInternalServer,
	}
}
