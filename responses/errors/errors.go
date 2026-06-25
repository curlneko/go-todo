package errors

import (
	"net/http"
)

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// AppError専用の関数、レシーバ（receiver）
/*
	Goには標準でこういうルールがあります：

	type error interface {
		Error() string
	}

	Error() string を持っていれば error として扱える
*/
func (e *ErrorResponse) Error() string {
	return e.Message
}

var (
	// ErrDuplicateTitle という変数に *AppError 型の値を代入。代入時は &AppError{...} でポインタ型を作る
	ErrTodoNotFound = &ErrorResponse{
		Code:    "TODO_NOT_FOUND",
		Message: "todo not found",
		Status:  http.StatusNotFound,
	}

	ErrDuplicateTitle = &ErrorResponse{
		Code:    "DUPLICATE_TITLE",
		Message: "todo title already exists",
		Status:  http.StatusConflict,
	}

	ErrInvalidID = &ErrorResponse{
		Code:    "INVALID_ID",
		Message: "invalid todo id",
		Status:  http.StatusBadRequest,
	}

	ErrInvalidRequest = &ErrorResponse{
		Code:    "INVALID_REQUEST",
		Message: "invalid request payload",
		Status:  http.StatusBadRequest,
	}

	ErrInternal = &ErrorResponse{
		Code:    "INTERNAL_ERROR",
		Message: "internal server error",
		Status:  http.StatusInternalServerError,
	}
)
