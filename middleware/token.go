package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var token = "123456"

func TokenCheck(c *gin.Context) {
	accessToken := c.GetHeader("access_token")
	if accessToken != token {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "token检验失败",
		})
		//中断后续中间件的执行
		//c.Abort()
		c.AbortWithError(http.StatusInternalServerError, errors.New("token check fail"))
	}
	c.Set("user_name", "nick")
	c.Set("user_id", "100")
	c.Next()
}
