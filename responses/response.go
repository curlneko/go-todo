package responses

import (
	"github.com/gin-gonic/gin"

	errResponse "gin-todo/responses/errors"
	"gin-todo/responses/successes"
)

func HandleError(c *gin.Context, err error) {
	if appErr, ok := err.(*errResponse.ErrorResponse); ok {
		c.JSON(appErr.Status, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
		})
		return
	}

	c.JSON(errResponse.ErrInternal.Status, gin.H{
		"code":    errResponse.ErrInternal.Code,
		"message": errResponse.ErrInternal.Message,
	})
}

func HandleSuccess(c *gin.Context, status int, data any) {
	success := successes.NewSuccessResponse(data)
	c.JSON(status, success)
}
