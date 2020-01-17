package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"username_across_platforms/server/service"
)


func Username(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.New("issues from request"))
		return
	}
	var urls []string

	err = json.Unmarshal(body, &urls)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, errors.New("issues from request un marshalling"))
		return
	}
	fmt.Println("the array: ", urls)

	matchedUrls := service.UsernameCheck(urls)

	if len(matchedUrls) == 0 {
		c.JSON(http.StatusNoContent, "No item")
		return
	}
	c.JSON(http.StatusOK, matchedUrls)
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
