package response

import (
	"net/http"

	"gitee.com/chensyi/vblog/exception"
	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}

func Failed(c *gin.Context, err error) {
	defer c.Abort()
	var e *exception.Exception
	if v, ok := err.(*exception.Exception); ok {
		e = v
	} else {
		e = exception.NewException(http.StatusInternalServerError, err.Error())
	}
	c.JSON(e.Code, e.Message)
}
