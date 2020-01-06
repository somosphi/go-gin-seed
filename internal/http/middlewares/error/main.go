package error_middleware

import (
	"github.com/gin-gonic/gin"
)

func Middleware(c *gin.Context) {
	err := c.Errors.Last()
	if err == nil {
		return
	}

}
