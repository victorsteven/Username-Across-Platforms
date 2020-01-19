package app

import (
	"github.com/gin-gonic/contrib/static"
	"username_across_platforms/server/controller"
	"username_across_platforms/server/middleware"
)

func route() {
	router.Use(middleware.CORSMiddleware()) //to enable api request between client and server
	router.Use(static.Serve("/", static.LocalFile("./web", true))) //for the vue app

	router.POST("/username", controller.Username)
	//router.GET("/ping", controller.Ping)
}