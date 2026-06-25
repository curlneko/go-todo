package errors

import (
	"net/http"
)

var (
	// ErrDuplicateTitle という変数に *AppError 型の値を代入。代入時は &AppError{...} でポインタ型を作る
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

	ErrInvalidID = &AppError{
		Code:    "INVALID_ID",
		Message: "invalid todo id",
		Status:  http.StatusBadRequest,
	}

	ErrInvalidRequest = &AppError{
		Code:    "INVALID_REQUEST",
		Message: "invalid request payload",
		Status:  http.StatusBadRequest,
	}

	ErrInternal = &AppError{
		Code:    "INTERNAL_ERROR",
		Message: "internal server error",
		Status:  http.StatusInternalServerError,
	}
)
