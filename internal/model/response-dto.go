package model

type ResponseDTO struct {
	TotalDuration                 string      `json:"totalDuration"`
	TotalCompletedRequests        int         `json:"totalCompletedRequests"`
	TotalHttpOkStatusCodeRequests int         `json:"totalHttpOkStatusCodeRequests"`
	HttpStatusReportSummary       map[int]int `json:"httpStatusReportSummary"`
}
