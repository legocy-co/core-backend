package errors

type AppErrorType string

const (
	ConflictError   = AppErrorType("conflict")
	NotFoundError   = AppErrorType("notFound")
	ValidationError = AppErrorType("validationFailed")
	PermissionError = AppErrorType("accessDenied")
	InternalError   = AppErrorType("internalError")
)

type AppError struct {
	errorType    AppErrorType
	errorMessage string
}

func NewAppError(errorType AppErrorType, errorMessage string) AppError {
	return AppError{
		errorType:    errorType,
		errorMessage: errorMessage,
	}
}

func (e AppError) Error() string {
	return e.GetErrorMessage()
}

func (e AppError) GetErrorType() AppErrorType {
	return e.errorType
}

func (e AppError) GetErrorMessage() string {
	return e.errorMessage
}
