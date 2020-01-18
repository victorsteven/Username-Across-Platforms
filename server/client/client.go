package client

import (
	"net/http"
)

var (
	ClientCall HTTPClient = &clientCall{}
)

type HTTPClient interface {
	Get(url string) (*http.Response, error)
}
type clientCall struct {
	Client http.Client
}

func (ci *clientCall) Get(url string) (*http.Response, error) {
	//request, err := http.NewRequest(http.MethodGet, url, nil)
	//if err != nil {
	//	return nil, err
	//}
	//client := http.Client{}

	res, err := ci.Client.Get(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}
