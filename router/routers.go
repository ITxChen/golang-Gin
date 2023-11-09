package router

import (
	"ginchat/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	//引入全局中间件
	r.Use(middleware.LogInput)
	//引入
	InitApi(r)
	InitCourse(r)
}
