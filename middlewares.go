package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func BodyMiddleware(context *gin.Context) {
	println("ContentLength:" + strconv.FormatInt(context.Request.ContentLength, 10))
}

func ReqResLoggerMiddleware(context *gin.Context) {
	println(context.Request.Method + " " + context.Request.RequestURI)
	context.Next()
	println("Response for " + context.Request.Method + " " + context.Request.RequestURI + " : " + strconv.Itoa(context.Writer.Status()))
}
