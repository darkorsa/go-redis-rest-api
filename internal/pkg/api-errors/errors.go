package apiErrors

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ApiErrorInterface interface {
	GetMessage() string
	GetStatus() int
	GetError() string
	GetInfo() []interface{}
}

type apiError struct {
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Error   string        `json:"error"`
	Info    []interface{} `json:"info"`
}

func (e apiError) GetMessage() string {
	return e.Message
}

func (e apiError) GetStatus() int {
	return e.Status
}

func (e apiError) GetError() string {
	return e.Error
}

func (e apiError) GetInfo() []interface{} {
	return e.Info
}

func NewError(message string, status int, err string, info []interface{}) ApiErrorInterface {
	return apiError{
		Message: message,
		Status:  status,
		Error:   err,
		Info:    info,
	}
}

func NewErrorFromBytes(bytes []byte) (ApiErrorInterface, error) {
	var apiErr ApiErrorInterface
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return apiErr, nil
}

func NewBadRequestError(message string) ApiErrorInterface {
	return apiError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFoundError(message string) ApiErrorInterface {
	return apiError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewUnauthorizedError(message string) ApiErrorInterface {
	return apiError{
		Message: message,
		Status:  http.StatusUnauthorized,
		Error:   "unauthorized",
	}
}

func NewInternalServerError(message string, err error) ApiErrorInterface {
	result := apiError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
	if err != nil {
		result.Info = append(result.Info, err.Error())
	}
	return result
}
