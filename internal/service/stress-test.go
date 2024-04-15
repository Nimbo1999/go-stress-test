package service

import (
	"github.com/nimbo1999/go-stress-test/internal/client"
	"github.com/nimbo1999/go-stress-test/internal/model"
)

type StressTestService interface {
	RunTest(url string) model.TestResult
}

type defaultStressTestService struct {
	Client client.HttpClient
}

func (service defaultStressTestService) RunTest(url string) model.TestResult {
	response, err := service.Client.Get(url)

	if err != nil {
		return model.TestResult{
			Error:      err,
			HttpStatus: 500,
		}
	}

	return model.TestResult{
		Error:      err,
		HttpStatus: response.StatusCode,
	}
}

func NewStressTestService(client client.HttpClient) *defaultStressTestService {
	return &defaultStressTestService{client}
}
