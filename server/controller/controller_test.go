package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"username_across_platforms/server/service"
)

var (
	getUsernameService func(urls []string) []string
)

type serviceMock struct {}

func (sm *serviceMock) UsernameCheck(urls []string) []string {
	return getUsernameService(urls)
}

func TestUsername_Success(t *testing.T) {
	service.UsernameService = &serviceMock{} //will now make our fake struct to implement the "usernameService" interface
	getUsernameService = func(urls []string) []string {
		return []string{
			"https://twitter.com/stevensunflash",
			"https://dev.to/stevensunflash",
			"https://instagram.com/stevensunflash",
		}
	}
	r := gin.Default()
	jsonBody := `["https://twitter.com/stevensunflash", "https://instagram.com/stevensunflash", "https://dev.to/stevensunflash"]`

	req, err := http.NewRequest(http.MethodPost, "/username", bytes.NewBufferString(jsonBody))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	r.POST("/username", Username)
	r.ServeHTTP(rr, req)

	var result []string
	err = json.Unmarshal(rr.Body.Bytes(), &result)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, http.StatusOK, rr.Code)
	assert.EqualValues(t, 3, len(result))
}

//func TestUsername_No_Match(t *testing.T) {
//	service.UsernameService = &serviceMock{} //will now make our fake struct to implement the "usernameService" interface
//	getUsernameService = func(urls []string) []string {
//		return []string{}
//	}
//	r := gin.Default()
//	jsonBody := `["https://twitter.com/not_valid_username_xxx", "https://instagram.com/not_valid_username_xxx", "https://github.com/not_valid_username_xxx"]`
//
//	req, err := http.NewRequest(http.MethodPost, "/username", bytes.NewBufferString(jsonBody))
//	if err != nil {
//		t.Errorf("this is the error: %v\n", err)
//	}
//	rr := httptest.NewRecorder()
//	r.POST("/username", Username)
//	r.ServeHTTP(rr, req)
//
//	var result []string
//	err = json.Unmarshal(rr.Body.Bytes(), &result)
//
//	assert.Nil(t, err)
//	assert.NotNil(t, result) //the result is empty but not nil
//	assert.EqualValues(t, http.StatusOK, rr.Code)
//	assert.EqualValues(t, 0, len(result))
//}


//Here, we dont need to mock the service since we will never get there.
func TestUsername_Invalid_Data(t *testing.T) {
	r := gin.Default()
	//instead of using array syntax, we used object
	jsonBody := `{"https://twitter.com/stevensunflash", "https://instagram.com/stevensunflash", "https://github.com/stevensunflash"}`

	req, err := http.NewRequest(http.MethodPost, "/username", bytes.NewBufferString(jsonBody))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	r.POST("/username", Username)
	r.ServeHTTP(rr, req)

	var result []string
	err = json.Unmarshal(rr.Body.Bytes(), &result)

	assert.NotNil(t, err)
	assert.Nil(t, result)
	assert.EqualValues(t, http.StatusUnprocessableEntity, rr.Code)
}