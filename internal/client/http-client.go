package client

import (
	"net/http"
	"time"
)

type HttpClient interface {
	Get(url string) (*http.Response, error)
}

type appHttpClient struct {
	timeout time.Duration
}

func (app appHttpClient) Get(url string) (*http.Response, error) {
	client := http.DefaultClient
	client.Timeout = app.timeout

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return client.Do(request)
}

func NewAppHattpClient(timeout time.Duration) *appHttpClient {
	return &appHttpClient{timeout}
}
