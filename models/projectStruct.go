package models

import "time"

type RequestParams struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	MinCount  int    `json:"minCount"`
	MaxCount  int    `json:"maxCount"`
}

type ResponseItem struct {
	Key        string    `bson:"key"`
	CreatedAt  time.Time `bson:"createdAt"`
	TotalCount int       `bson:"totalCount"`
}

type Response struct {
	Code    int            `json:"code"`
	Message string         `json:"msg"`
	Records []ResponseItem `json:"records"`
}

type Payload struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
