package errors

import (
	"net/http"
)

var (
	ErrTodoNotFound = &AppError{
		Code:    "TODO_NOT_FOUND",
		Message: "todo not found",
		Status:  http.StatusNotFound,
	}

	ErrDuplicateTitle = &AppError{
		Code:    "DUPLICATE_TITLE",
		Message: "todo title already exists",
		Status:  http.StatusConflict,
	}
)
