package main

import (
	"./handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// Default With the Logger and Recovery middleware already attached
	// router := gin.Default()

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(ReqResLoggerMiddleware)

	hello := router.Group("hellogin")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	hello.Use(BodyMiddleware)

	println(hello)

	// note: scoping for readablity?
	{
		hello.POST("/logmein", func(context *gin.Context) {
			panic("login not working...try again after some time")
		})

		hello.POST("/login", handlers.LoginHandler)
	}

	hello.GET("/ping", handlers.PingPongHandler)

	hello.GET("/user/:name/*action", handlers.NameActionHandler)

	hello.GET("/welcome", handlers.WelcomeHandler)

	router.Run(":3000")
}
