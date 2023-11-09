package router

import (
	"ginchat/web"
	"github.com/gin-gonic/gin"
)

func InitApi(r *gin.Engine) {
	//没有权限的
	api := r.Group("api")
	//接口进行版本控制
	v1 := api.Group("/v1")
	v1.GET("/ping", web.Ping)
	v1.POST("/login", web.Login)
	v1.POST("/register", web.Register)
}
