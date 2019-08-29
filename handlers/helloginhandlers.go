package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func PingPongHandler(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "pong",
	})
}

func NameActionHandler(context *gin.Context) {
	name := context.Param("name")
	action := context.Param("action")
	message := name + " is " + action
	context.String(http.StatusOK, message)
}

func WelcomeHandler(context *gin.Context) {
	firstName := context.DefaultQuery("firstName", "Guest")
	queryParam := context.Request.URL.Query()
	print(queryParam)
	lastName, pass := context.GetQuery("lastname")
	if pass {
		context.String(http.StatusOK, "Welcome "+firstName+" "+lastName)
	} else {
		context.String(http.StatusOK, "Welcome "+firstName)
	}
}