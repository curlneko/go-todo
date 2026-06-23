package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"

	appErrors "gin-todo/errors"
)

func HandleError(c *gin.Context, err error) {
	if appErr, ok := err.(*appErrors.AppError); ok {
		c.JSON(appErr.Status, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
		})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"code":    "INTERNAL_ERROR",
		"message": err.Error(),
	})
}
