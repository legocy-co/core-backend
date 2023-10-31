package errors

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"legocy-go/internal/domain/errors"
	"net/http"
)

type HttpResponseError struct {
	Status  int
	Message gin.H
}

func (e HttpResponseError) Error() string {
	return fmt.Sprintf("Error %v", e.Status)
}

func NewHttpResponseError(status int, message string) HttpResponseError {
	return HttpResponseError{Status: status, Message: gin.H{"error": message}}
}

func (HttpResponseError) FromAppError(e errors.AppError) HttpResponseError {
	var status int = http.StatusInternalServerError

	switch et := e.GetErrorType(); et {
	case errors.ConflictError:
		status = http.StatusConflict
	case errors.PermissionError:
		status = http.StatusForbidden
	case errors.NotFoundError:
		status = http.StatusNotFound
	case errors.ValidationError:
		status = http.StatusUnprocessableEntity
	}

	return HttpResponseError{Status: status, Message: gin.H{"error": e.GetErrorMessage()}}
}
