package components

import "strings"

type OpenAIError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Param   string `json:"param"`
	Code    any    `json:"code"`
}

type OpenAIErrorWithCode struct {
	Error      OpenAIError `json:"error"`
	StatusCode int         `json:"status_code"`
	LocalError bool
}

// OpenAIErrorWrapper wraps an error into an OpenAIErrorWithStatusCode
func OpenAIErrorWrapper(text string, code string, statusCode int) *OpenAIErrorWithCode {
	lowerText := strings.ToLower(text)
	if strings.Contains(lowerText, "post") || strings.Contains(lowerText, "dial") || strings.Contains(lowerText, "http") {
		text = "请求上游失败"
	}
	openAIError := OpenAIError{
		Message: text,
		Type:    "http_api_error",
		Code:    code,
	}
	return &OpenAIErrorWithCode{
		Error:      openAIError,
		StatusCode: statusCode,
	}
}
