package errors

import (
	"errors"
	"net/http"
)

type RestErr struct {
	Message string `json:"message"`
	Status int `json:"status"`
	Error string `json:"error"`
}

func NewError(msg string) error {
	return errors.New(msg)
}

func NewBadRequestError(message string) *RestErr  {
	return &RestErr{
		Message: message,
		Status: http.StatusBadRequest,
		Error: "bad request",
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status: http.StatusNotFound,
		Error: "not found",
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status: http.StatusInternalServerError,
		Error: "internal server error",
	}
}
