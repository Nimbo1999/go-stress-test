package model

import (
	"sync"
)

type StressTestFinalResult struct {
	mu                            sync.Mutex
	TotalCompletedRequests        int
	TotalHttpOkStatusCodeRequests int
	HttpStatusReportSummary       map[int]int
}

func (result *StressTestFinalResult) IncrementCompletedRequests() {
	result.mu.Lock()
	defer result.mu.Unlock()
	result.TotalCompletedRequests += 1
}

func (result *StressTestFinalResult) IncrementHttpSuccessRequests() {
	result.mu.Lock()
	defer result.mu.Unlock()
	result.TotalHttpOkStatusCodeRequests += 1
}

func (result *StressTestFinalResult) AddStatusCodeToReport(statusCode int) {
	result.mu.Lock()
	defer result.mu.Unlock()
	_, ok := result.HttpStatusReportSummary[statusCode]
	if ok {
		result.HttpStatusReportSummary[statusCode] += 1
		return
	}
	result.HttpStatusReportSummary[statusCode] = 1
}

func (result *StressTestFinalResult) GetTotalCompletedRequests() int {
	result.mu.Lock()
	defer result.mu.Unlock()
	return result.TotalCompletedRequests
}
