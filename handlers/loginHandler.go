package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	User     string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Zone     string `json:"Zone" binding:"-"`
}

func LoginHandler(context *gin.Context) {
	var login Login
	if err := context.ShouldBindJSON(&login); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// gin.H -> H is a shortcut for map[string]interface{} -> Dictionary<string, object>
	context.JSON(http.StatusOK, gin.H{"status": login.User + ", you are logged in " + login.Zone}) 
}
