package model

type APIError struct {
	ApiErr string `json:"apiError"`
	Id     *int   `json:"id,omitempty"`
}

func NewAPIError(s string) *APIError {
	return &APIError{ApiErr: s}
}
