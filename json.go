package fen

const JsonResponseKey = "__fen_JsonResponseKey{}"

type JSON struct {
	Code       int       `json:"code,omitempty"`
	Message    string    `json:"message,omitempty"`
	Error      JsonError `json:"error,omitempty"`
	ErrorStack []string  `json:"error_stack,omitempty"`
}

type JsonError struct {
	Status  string           `json:"status,omitempty"`
	Details []JsonErrorField `json:"details,omitempty"`
}

type JsonErrorField struct {
	Field       string `json:"field,omitempty"`
	Description string `json:"description,omitempty"`
}

type JsonResponse interface {
	SetCode(int) JsonResponse
	SetMessage(string) JsonResponse
	SetError(JsonError) JsonResponse
	SetErrorStack([]string) JsonResponse
}

func (j *JSON) SetCode(code int) JsonResponse {
	j.Code = code
	return j
}

func (j *JSON) SetMessage(message string) JsonResponse {
	j.Message = message
	return j
}

func (j *JSON) SetError(err JsonError) JsonResponse {
	j.Error = err
	return j
}

func (j *JSON) SetErrorStack(stack []string) JsonResponse {
	j.ErrorStack = stack
	return j
}
