package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"log"
)

func LogInput(ctx *gin.Context) {
	//将body输入流读出后，body将会被清空，因此需要重新赋值
	requestBody, _ := ctx.GetRawData()
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))

	mp := make(map[string]interface{})
	mp["request_url"] = ctx.Request.RequestURI
	mp["status"] = ctx.Writer.Status()
	mp["method"] = ctx.Request.Method
	mp["access_token"] = ctx.GetHeader("accexx_token")
	//存在文件时不要打印文件内容
	mp["body"] = string(requestBody)
	log.Println(mp)
	ctx.Next()
}
