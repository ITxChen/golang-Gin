package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthCheck(c *gin.Context) {
	userId, _ := c.Get("user_id")
	userName, _ := c.Get("user_name")
	fmt.Printf("auth check called,userId:%v,userName:%s\n", userId, userName)
	c.JSON(http.StatusInternalServerError, gin.H{
		"userId":   userId,
		"userName": userName,
	})
}
