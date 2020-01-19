package service

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"username_across_platforms/server/client"
)

var (
	getRequestFunc func(url string) (*http.Response, error)
)
type clientMock struct {}

//mocking the client call, so we dont hit the real endpoint:
func (cm *clientMock) Get(url string) (*http.Response, error) {
	return getRequestFunc(url)
}


func TestUsernameCheck_Success(t *testing.T) {
	urls := []string{
		"http://twitter.com/stevensunflash",
		"http://instagram.com/stevensunflash",
		"http://dev.to/stevensunflash",
	}

	getRequestFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode:       http.StatusOK,
		}, nil
	}
	client.ClientCall = &clientMock{}

	result := UsernameService.UsernameCheck(urls)

	assert.NotNil(t, result)
	assert.EqualValues(t, len(result), 3)
}


func TestUsernameCheck_No_Match(t *testing.T) {
	urls := []string{
		"http://twitter.com/no_match_username",
		"http://instagram.com/no_match_username",
		"http://dev.to/no_match_username",
	}
	getRequestFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode:       http.StatusNotFound, //it can be 404, 422 or 500 depending the api
		}, nil
	}
	client.ClientCall = &clientMock{}

	result := UsernameService.UsernameCheck(urls)

	assert.Nil(t, result)
	assert.EqualValues(t, len(result), 0)
}