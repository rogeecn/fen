package fen

type JSON struct {
	Code       int         `json:"code,omitempty"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	ErrorStack []string    `json:"error_stack,omitempty"`
}
