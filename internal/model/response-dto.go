package model

import "time"

type ResponseDTO struct {
	TotalDuration                 time.Duration `json:"totalDuration"`
	TotalCompletedRequests        int           `json:"totalCompletedRequests"`
	TotalHttpOkStatusCodeRequests int           `json:"totalHttpOkStatusCodeRequests"`
	HttpStatusReportSummary       map[int]int   `json:"httpStatusReportSummary"`
}
