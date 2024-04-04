package model

type APIError struct {
	ApiErr string `json:"apiError"`
}

func NewAPIError(s string) *APIError {
	return &APIError{ApiErr: s}
}

type IdRespond struct {
	Id int64 `json:"id"`
}

func NewIdRespond(id int64) *IdRespond {
	return &IdRespond{Id: id}
}

type TokenRespond struct {
	Token string `json:"token"`
}

func NewTokenRespond(token string) *TokenRespond {
	return &TokenRespond{Token: token}
}
