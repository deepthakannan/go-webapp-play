package main

import (
	"strconv"
	"github.com/gin-gonic/gin"
)

func BodyMiddleware(context *gin.Context) {
	println("ContentLength:" + strconv.FormatInt(context.Request.ContentLength, 10))
	context.Next()
}

