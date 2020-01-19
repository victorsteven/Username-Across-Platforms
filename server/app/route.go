package app

import (
	"username_across_platforms/server/controller"
	"username_across_platforms/server/middleware"
)

func route() {
	router.Use(middleware.CORSMiddleware()) //to enable api request between client and server

	router.POST("/username", controller.Username)
}