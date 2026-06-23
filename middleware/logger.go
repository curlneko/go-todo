package middleware

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		switch c.Request.Method {
		case "POST", "PUT", "PATCH":
			// リクエストボディを読み込む。[]byte型
			bodyBytes, err := io.ReadAll(c.Request.Body)

			if err != nil {
				fmt.Println("Failed to read request body:", err)
			} else {
				fmt.Println("Request Body:", string(bodyBytes))
			}

			// Request.Bodyは一度読むと空になるため、Body を元に戻す
			c.Request.Body = io.NopCloser(
				bytes.NewBuffer(bodyBytes),
			)
		}

		// Controller を実行
		c.Next()

		duration := time.Since(start)

		status := c.Writer.Status()

		logLevel := "[INFO]"

		if status >= 500 {
			logLevel = "[ERROR]"
		} else if status >= 400 {
			logLevel = "[WARN]"
		}

		fmt.Printf(
			"%s %s %s %d %v IP=%s UA=%s\n",
			logLevel,
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			duration,
			c.ClientIP(),
			c.Request.UserAgent(),
		)
	}
}
