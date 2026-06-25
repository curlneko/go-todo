package utils

import (
	"github.com/gin-gonic/gin"

	appErrors "gin-todo/errors"
)

func HandleError(c *gin.Context, err error) {
	// err が *AppError 型かどうか。*appErrors.AppError＝AppErrorのポインタ型
	if appErr, ok := err.(*appErrors.AppError); ok {
		c.JSON(appErr.Status, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
		})
		return
	}

	c.JSON(appErrors.ErrInternal.Status, gin.H{
		"code":    appErrors.ErrInternal.Code,
		"message": appErrors.ErrInternal.Message,
	})
}
