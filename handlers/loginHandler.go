package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Login struct {
	User string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginHandler(context *gin.Context) {
	var login Login
	if err := context.ShouldBindJSON(&login); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{ "error": err.Error()})
		return
	}
	
	context.JSON(http.StatusOK, gin.H{"status": login.User + ", you are logged in"})
}