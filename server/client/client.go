package client

import (
	"net/http"
)

var (
	ClientCall HTTPClient = &clientCall{}
)

type HTTPClient interface {
	GetValue(url string) (*http.Response, error)
}
type clientCall struct {
	Client http.Client
}

func (ci *clientCall) GetValue(url string) (*http.Response, error) {
	res, err := ci.Client.Get(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}
