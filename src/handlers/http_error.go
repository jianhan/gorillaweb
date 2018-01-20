package handlers

import "fmt"

type httpError struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}

func newHTTPError(code uint, message string) *httpError {
	return &httpError{
		Code:    code,
		Message: message,
	}
}

func (h *httpError) Error() string {
	return fmt.Sprintf("[%d] %s", h.Code, h.Message)
}
