package app

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"os"
	"username_across_platforms/server/controller"
	"username_across_platforms/server/middleware"
)

var (
	router = gin.Default()
)

func StartApp() {

	router.Use(middleware.CORSMiddleware()) //to enable api request between client and server
	router.Use(static.Serve("/", static.LocalFile("./web", true))) //for the vue app

	router.POST("/username", controller.Username)
	router.GET("/ping", controller.Ping)


	port := os.Getenv("PORT") //using heroku host
	if port == "" {
		port = "8888" //localhost
	}

	router.Run(":"+port)
}