package global_errors

import "net/http"

type GlobalErrors struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes,omitempty"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *GlobalErrors) Error() string {
	return r.Message
}

func NewGlobalErrors(message string, err string, code int, causes []Causes) *GlobalErrors {
	return &GlobalErrors{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *GlobalErrors {
	return &GlobalErrors{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *GlobalErrors {
	return &GlobalErrors{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *GlobalErrors {
	return &GlobalErrors{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *GlobalErrors {
	return &GlobalErrors{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *GlobalErrors {
	return &GlobalErrors{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}

func NewBadGateway(message string) *GlobalErrors {
	return &GlobalErrors{
		Message: message,
		Err:     "bad_gateway",
		Code:    http.StatusBadGateway,
	}
}

func NewNoContent(message string) *GlobalErrors {
	return &GlobalErrors{
		Message: message,
		Err:     "no_content",
		Code:    http.StatusNoContent,
	}
}
