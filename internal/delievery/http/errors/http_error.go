package errors

import (
	"github.com/gin-gonic/gin"
	"legocy-go/internal/domain/errors"
	"net/http"
)

type HttpResponseError struct {
	Status  int
	Message gin.H
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
