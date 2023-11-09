package router

import (
	"ginchat/middleware"
	"ginchat/web"
	"github.com/gin-gonic/gin"
)

func InitCourse(r *gin.Engine) {
	// token登录检查
	course := r.Group("/course", middleware.TokenCheck)
	v1 := course.Group("/v1")
	// http://loaclhost:8080/course/v1/detail/10
	v1.GET("/detail/:id", web.GetCourseDetail)
	v1.GET("/view/:id", web.GetCourseView)
	//加权限
	admin := course.Group("/admin", middleware.AuthCheck)
	adminV1 := admin.Group("/v1")
	adminV1.POST("/add", web.AddCourse)
	adminV1.POST("/publish", web.PublishCourse)
	adminV1.POST("/upload", web.UploadCourse)

}
