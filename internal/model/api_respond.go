package model

type APIError struct {
	ApiErr string `json:"apiError"`
}

func NewAPIError(s string) *APIError {
	return &APIError{ApiErr: s}
}

type IdRespond struct {
	Id int `json:"id"`
}

func NewIdRespond(id int) *IdRespond {
	return &IdRespond{Id: id}
}
