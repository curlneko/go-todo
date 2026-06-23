package errors

type AppError struct {
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
func (e *AppError) Error() string {
	return e.Message
}
