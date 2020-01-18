package service

import (
	"fmt"
	"testing"
)

var (
	getProviderFunc func(url string, c chan string)
)

type checkerMock struct{}

func (cm *checkerMock) CheckUrl(url string, c chan string) {}


func TestUsernameCheck_Success(t *testing.T) {
	//getProviderFunc = func(url string, c chan string){}
	//provider.Checker = &checkerMock{}

		urls := []string{
			"http://twitter.com/stevensunflash",
			"http://github.com/stevensunflash",
			"http://instagram.com/stevensunflash",
			"http://bitbucket.com/stevensunflash",
			"http://dev.to/stevensunflash",
		}

	result := UsernameService.UsernameCheck(urls)

	fmt.Println("the result: ", result)

}

//we are mocking the checkUrls method
//func (cm *getCheckerMock) checkUrls(url string, c chan string){}

//func TestUsernameCheck(t *testing.T) {
//
//	urls := []string{
//		"http://twitter.com/stevensunflash",
//		"http://github.com/stevensunflash",
//		"http://instagram.com/stevensunflash",
//		"http://bitbucket.com/stevensunflash",
//		"http://dev.to/stevensunflash",
//	}
//}



//func TestCheckUrls_Success(t *testing.T) {
//	url := "https://twitter.com/stevensunflash"
//	ch := make(chan string)
//	go provider.Checker.CheckUrls(url, ch)
//	result := <-ch
//	assert.NotNil(t, result)
//	assert.EqualValues(t, "https://twitter.com/stevensunflash", result)
//}
//
//func TestCheckUrls_Not_Existent_Url(t *testing.T) {
//	url := "https://wrong.com/stevensunflash"
//	ch := make(chan string)
//	go provider.Checker.CheckUrls(url, ch)
//	result := <-ch
//	assert.NotNil(t, result)
//	assert.EqualValues(t, "cant_access_resource", result)
//}
//
//func TestCheckUrls_Username_Dont_Exist(t *testing.T) {
//	url := "https://twitter.com/random_xxxxx_yyyy"
//	ch := make(chan string)
//	go provider.Checker.CheckUrls(url, ch)
//	result := <-ch
//	assert.NotNil(t, result)
//	assert.EqualValues(t, "no_match", result)
//}

//func TestReposService_CreateRepoConcurrentInvalidRequest(t *testing.T) {
//	request := repositories.CreateRepoRequest{}
//	output := make(chan repositories.CreateRepositoriesResult)
//
//	service := reposService{}
//	go service.CreateRepoConcurrent(request, output)
//	result := <-output
//	assert.NotNil(t, result)
//	assert.Nil(t, result.Response)
//	assert.NotNil(t, result.Error)
//	assert.EqualValues(t, http.StatusBadRequest, result.Error.Status())
//	assert.EqualValues(t, "Invalid repository name", result.Error.Message())
//}