package client

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

//USING HTTP TRANSPORT
// RoundTripFunc .
type RoundTripFunc func(req *http.Request) (*http.Response, error)

// RoundTrip . //this method is from the RoundTripper interface
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

//NewFakeClient returns *http.Client with Transport replaced to avoid making real calls
func NewFakeClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

func TestGetWithRoundTripper_Success(t *testing.T) {
	client := NewFakeClient(func(req *http.Request) (*http.Response, error) {
		//return the response we want
		return &http.Response{
			StatusCode: 200,
			// Must be set to non-nil value or it panics
			Header:     make(http.Header),
		}, nil
	})
	api := clientCall{*client}
	url := "https://twitter.com/stevensunflash" //this url can be anything
	body, err := api.Get(url)
	assert.Nil(t, err)
	assert.NotNil(t, body)
	assert.EqualValues(t, http.StatusOK, body.StatusCode)
}

func TestGetWithRoundTripper_No_Match(t *testing.T) {
	client := NewFakeClient(func(req *http.Request) (*http.Response, error) {
		//return the response we want
		return &http.Response{
			StatusCode: 404, //the real api status code may be 404, 422, 500. But we dont care
			// Must be set to non-nil value or it panics
			Header:     make(http.Header),
		}, nil
	})
	api := clientCall{*client}
	url := "https://twitter.com/no_match_random" //we passed in a user that is not found
	body, err := api.Get(url)
	assert.Nil(t, err)
	assert.NotNil(t, body)
	assert.EqualValues(t, http.StatusNotFound, body.StatusCode)
}

func TestGetWithRoundTripper_Failure(t *testing.T) {
	client := NewFakeClient(func(req *http.Request) (*http.Response, error) {
		return nil, errors.New("we couldn't access the url provided") //the response we want
	})
	api := clientCall{*client}
	url := "https://invalid_url/stevensunflash" //we passed an invalid url
	body, err := api.Get(url)
	assert.NotNil(t, err)
	assert.Nil(t, body)
	assert.EqualValues(t, "Get https://invalid_url/stevensunflash: we couldn't access the url provided", err.Error())
}


//USING HTTPTEST SERVER
//func TestDoStuffWithTestServer(t *testing.T) {
//	// Start a local HTTP server
//	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
//		// Test request parameters
//		//equals(t, req.URL.String(), "/some/path")
//		// Send response to be tested
//		rw.Write([]byte(`OK`))
//	}))
//	// Close the server when test finishes
//	defer server.Close()
//
//	// Use Client & URL from our local test server
//	api := clientCall{server.Client()}
//	url := "https://twitter.com/stevensunflash" //this url can be anything
//	body, err := api.Get(url)
//	assert.Nil(t, err)
//	//assert.EqualValues(t, []byte("OK"), body)
//	assert.EqualValues(t, http.StatusOK, body.StatusCode)
//}
