package model

type ResponseData struct {
	StatusCode int `json:"status_code"`
	Body       any `json:"body"`
}

func NewResponseData(statusCode int, body any) *ResponseData {
	return &ResponseData{StatusCode: statusCode, Body: body}
}
