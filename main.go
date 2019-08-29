package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/user/:name/*action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")
		message := name + " is " + action
		context.String(http.StatusOK, message);
	})

	router.GET("/welcome", func(context * gin.Context) {
		firstName := context.DefaultQuery("firstName", "Guest")
		queryParam := context.Request.URL.Query()
		print(queryParam)
		lastName, pass := context.GetQuery("lastname")
		if (pass) {
			context.String(http.StatusOK, "Welcome " + firstName + " " + lastName);
		} else {
			context.String(http.StatusOK, "Welcome " + firstName);
		}
	})


	router.Run(":3000")
}