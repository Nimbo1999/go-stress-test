package usecase

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/nimbo1999/go-stress-test/internal/model"
	"github.com/nimbo1999/go-stress-test/internal/service"
)

type TestResult = model.TestResult

type RunStressTestUseCaseFunc func(url string, requests, concurrency int)

type RequestCompleted struct {
	mu          sync.Mutex
	RequestMade int
}

func (req *RequestCompleted) Increment() {
	req.mu.Lock()
	defer req.mu.Unlock()
	req.RequestMade += 1
}

type StressTestUseCase struct {
	Url                   string
	Requests, Concurrency int
	service               service.StressTestService
}

func (usecase *StressTestUseCase) Execute() model.ResponseDTO {
	fmt.Println("Initializing requests...")
	now := time.Now()
	concurrencyRequestsChannel := make(chan int, usecase.Concurrency)
	resultsChannel := make(chan TestResult, usecase.Concurrency)

	finalResult := model.StressTestFinalResult{
		TotalCompletedRequests:        0,
		TotalHttpOkStatusCodeRequests: 0,
		HttpStatusReportSummary:       make(map[int]int),
	}

	for i := 0; i < usecase.Concurrency; i++ {
		go usecase.requestInitializer(concurrencyRequestsChannel, resultsChannel, usecase.Url, usecase.service)
		go usecase.resultReceiver(resultsChannel, &finalResult)
	}

	for i := 0; i < usecase.Requests; i++ {
		concurrencyRequestsChannel <- i
	}
	close(concurrencyRequestsChannel)

	for {
		if completedRequests := finalResult.GetTotalCompletedRequests(); completedRequests == usecase.Requests {
			break
		}
	}
	close(resultsChannel)

	return model.ResponseDTO{
		TotalDuration:                 fmt.Sprintf("%1.f seconds", time.Since(now).Seconds()),
		TotalCompletedRequests:        finalResult.TotalCompletedRequests,
		TotalHttpOkStatusCodeRequests: finalResult.TotalHttpOkStatusCodeRequests,
		HttpStatusReportSummary:       finalResult.HttpStatusReportSummary,
	}
}

func (usecase *StressTestUseCase) requestInitializer(
	channel chan int,
	testResultChannel chan TestResult,
	url string,
	service service.StressTestService,
) {
	for range channel {
		testResultChannel <- service.RunTest(url)
	}
}

func (usecase *StressTestUseCase) resultReceiver(channel chan TestResult, finalResult *model.StressTestFinalResult) {
	for result := range channel {
		if result.Error != nil {
			if netErr, ok := result.Error.(net.Error); ok && netErr.Timeout() {
				finalResult.AddStatusCodeToReport(http.StatusRequestTimeout)
			} else {
				finalResult.AddStatusCodeToReport(http.StatusInternalServerError)
			}
			finalResult.IncrementCompletedRequests()
			return
		}

		if result.HttpStatus != 0 {
			finalResult.AddStatusCodeToReport(result.HttpStatus)
		}
		if result.HttpStatus == http.StatusOK {
			finalResult.IncrementHttpSuccessRequests()
		}
		finalResult.IncrementCompletedRequests()
	}
}

func NewStressTestUseCase(url string, requests, concurrency int, service service.StressTestService) *StressTestUseCase {
	return &StressTestUseCase{
		Url:         url,
		Requests:    requests,
		Concurrency: concurrency,
		service:     service,
	}
}
