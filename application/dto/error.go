package dto

type ErrorDto struct {
	Error ErrorDetailsDto `json:"error"`
}

type ErrorDetailsDto struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewErrorDto(code int, message string, data interface{}) *ErrorDto {
	return &ErrorDto{Error: ErrorDetailsDto{Code: code, Message: message, Data: data}}
}
