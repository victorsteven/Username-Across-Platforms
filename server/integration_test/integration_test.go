package integration_test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"username_across_platforms/server/client"
	"username_across_platforms/server/controller"
)

//NOTE THE INTEGRATION TEST AIM IS NOT TO COVER THE NITTY-GRITTY OF INDIVIDUAL PACKAGES. TO TEST THESE PACKAGES, HEAD TO THEM AND RUN THEIR TESTS(WHICH OFFER INCREASED COVERAGE).
//THE ESSENCE OF THIS INTEGRATION TEST IS TO COUPLE ALL DONE AND RUNNING IN ONE TEST. WE ARE MORE INTERESTED IN THE TEST PASSING HERE THAN ANYTHING ELSE.


//INTEGRATION TEST FROM THE "CONTROLLER" TO THE "CLIENT" PACKAGE
//We have done unit tests in each package, where, methods are tested independent of others,
//Now lets test the from the CONTROLLER to the CLIENT, without mocking anything.
//We will be hitting the real API
//Note this test will take sometime since we are hitting the real API
func TestUsername_Success_Real_Endpoint(t *testing.T) {

	r := gin.Default()
	jsonBody := `["https://twitter.com/stevensunflash", "https://instagram.com/stevensunflash", "https://dev.to/stevensunflash"]`

	req, err := http.NewRequest(http.MethodPost, "/username", bytes.NewBufferString(jsonBody))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	r.POST("/username", controller.Username)
	r.ServeHTTP(rr, req)

	var result []string
	err = json.Unmarshal(rr.Body.Bytes(), &result)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, http.StatusOK, rr.Code)
	assert.EqualValues(t, 3, len(result))
}



//INTEGRATION TEST FROM THE "CONTROLLER" TO THE "PROVIDER" PACKAGE
//Here, we will fake hitting the Real API,
//The essence of this integration test is to test from the Controller to the Provider package, then faking the CLIENT
var (
	getRequestFunc func(url string) (*http.Response, error)
)
type clientMock struct {}
//mocking the client call:
func (cm *clientMock) GetValue(url string) (*http.Response, error) {
	return getRequestFunc(url)
}

func TestUsername_Success_Fake_Endpoint(t *testing.T) {
	getRequestFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode:       http.StatusOK,
		}, nil
	}
	client.ClientCall = &clientMock{}

	r := gin.Default()
	jsonBody := `["https://twitter.com/stevensunflash", "https://instagram.com/stevensunflash", "https://dev.to/stevensunflash"]`

	req, err := http.NewRequest(http.MethodPost, "/username", bytes.NewBufferString(jsonBody))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	r.POST("/username", controller.Username)
	r.ServeHTTP(rr, req)

	var result []string
	err = json.Unmarshal(rr.Body.Bytes(), &result)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, http.StatusOK, rr.Code)
	assert.EqualValues(t, 3, len(result))
}




//This test is already in the controller unit test file, this in order to increase coverage.
//The CLIENT is never reached, because bad json was supplied
func TestUsername_Invalid_Data(t *testing.T) {
	r := gin.Default()
	//instead of using array syntax, we used object
	jsonBody := `{"https://twitter.com/stevensunflash", "https://instagram.com/stevensunflash", "https://github.com/stevensunflash"}`

	req, err := http.NewRequest(http.MethodPost, "/username", bytes.NewBufferString(jsonBody))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	r.POST("/username", controller.Username)
	r.ServeHTTP(rr, req)

	var result []string
	err = json.Unmarshal(rr.Body.Bytes(), &result)

	assert.NotNil(t, err)
	assert.Nil(t, result)
	assert.EqualValues(t, http.StatusUnprocessableEntity, rr.Code)
}
