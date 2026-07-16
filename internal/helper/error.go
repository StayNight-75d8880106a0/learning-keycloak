package helper

import "net/http"

type AppError struct {
	Code      int           `json:"code"`
	Message   string        `json:"message"`
	ErrorCode string        `json:"errorCode"`
	Details   []ErrorDetail `json:"details"`
}

type ErrorDetail struct {
	Message string `json:"detailMessage"`
}

func (e *AppError) Error() string {
	return e.Message
}

func NewBadRequestError(message string, details []ErrorDetail) *AppError {
	return &AppError{
		Code:      http.StatusBadRequest,
		ErrorCode: "BAD_REQUEST",
		Message:   message,
		Details:   details,
	}
}

func NewInternalServerError(message string, details []ErrorDetail) *AppError {
	return &AppError{
		Code:      http.StatusInternalServerError,
		ErrorCode: "INTERNAL_SERVER_ERROR",
		Message:   message,
		Details:   details,
	}
}

func NewUnauthorizedError(message string, details []ErrorDetail) *AppError {
	return &AppError{
		Code:      http.StatusUnauthorized,
		ErrorCode: "UNAUTHORIZED",
		Message:   message,
		Details:   details,
	}
}

func NewUnprocessableEntityError(message string, details []ErrorDetail) *AppError {
	return &AppError{
		Code:      http.StatusUnprocessableEntity,
		ErrorCode: "UNPROCESSABLE_ENTITY",
		Message:   message,
		Details:   details,
	}
}
