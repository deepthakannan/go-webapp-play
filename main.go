package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Default With the Logger and Recovery middleware already attached
	// router := gin.Default()

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(func(context *gin.Context) {
		println(context.Request.Method + " " + context.Request.RequestURI)
	})

	hello := router.Group("hellogin")
	hello.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	hello.GET("/user/:name/*action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")
		message := name + " is " + action
		context.String(http.StatusOK, message)
	})

	hello.GET("/welcome", func(context *gin.Context) {
		firstName := context.DefaultQuery("firstName", "Guest")
		queryParam := context.Request.URL.Query()
		print(queryParam)
		lastName, pass := context.GetQuery("lastname")
		if pass {
			context.String(http.StatusOK, "Welcome "+firstName+" "+lastName)
		} else {
			context.String(http.StatusOK, "Welcome "+firstName)
		}
	})

	router.Run(":3000")
}
