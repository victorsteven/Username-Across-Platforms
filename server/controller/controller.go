package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"username_across_platforms/server/service"
)


func Username(c *gin.Context) {
	var urls []string
	if err := c.ShouldBindJSON(&urls); err != nil {
		c.JSON(http.StatusUnprocessableEntity, errors.New("invalid json body"))
		return
	}
	matchedUrls := service.UsernameService.UsernameCheck(urls)

	c.JSON(http.StatusOK, matchedUrls)
}

//func Ping(c *gin.Context) {
//	c.JSON(200, gin.H{
//		"message": "pong",
//	})
//}
