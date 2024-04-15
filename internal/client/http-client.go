package client

import (
	"context"
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
	ctx, cancel := context.WithTimeout(context.Background(), app.timeout)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	return http.DefaultClient.Do(request)
}

func NewAppHattpClient(timeout time.Duration) *appHttpClient {
	return &appHttpClient{timeout}
}
