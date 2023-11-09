package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCourseDetail(c *gin.Context) {
	id := c.Param("id")
	name := c.Query("name")
	c.JSON(http.StatusOK, gin.H{
		"id":   id,
		"name": name,
	})
}
func GetCourseView(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "GetCourseView",
	})
}
func AddCourse(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": " AddCourse",
	})
}
func PublishCourse(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "PublishCourse",
	})
}
func UploadCourse(c *gin.Context) {
	//file, _ := c.FormFile("file")
	//c.SaveUploadedFile(file, "upload/"+file.Filename)
	//c.JSON(http.StatusOK, gin.H{
	//	"file_name": file.Filename,
	//})
	form, _ := c.MultipartForm()
	files := form.File["files"]
	list := make([]string, len(files))
	for i, file := range files {
		c.SaveUploadedFile(file, "upload/"+file.Filename)
		list[i] = file.Filename
	}
	c.JSON(http.StatusOK, gin.H{
		"files": list,
	})
}
