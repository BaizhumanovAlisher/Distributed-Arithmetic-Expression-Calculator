package model

type ResponseData struct {
	StatusCode int    `json:"status_code"`
	Body       []byte `json:"body"`
}

func NewResponseData(statusCode int, body []byte) *ResponseData {
	return &ResponseData{StatusCode: statusCode, Body: body}
}
