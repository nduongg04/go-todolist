package errors

// Common error types as APIError with status codes
var (
	InvalidCredentials = &APIError{Code: 401, Message: "invalid email or password"}
	UserNotFound       = &APIError{Code: 404, Message: "user not found"}
	EmailTaken         = &APIError{Code: 400, Message: "email already taken"}
	Unauthorized       = &APIError{Code: 401, Message: "unauthorized"}
	Forbidden          = &APIError{Code: 403, Message: "forbidden"}
	InternalServer     = &APIError{Code: 500, Message: "internal server error"}
)

// Custom error type for API errors
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *APIError) Error() string {
	return e.Message
}

// Error constructors
func NewUnauthorizedError(msg string) *APIError {
	if msg == "" {
		msg = "unauthorized"
	}
	return &APIError{Code: 401, Message: msg}
}

func NewNotFoundError(msg string) *APIError {
	if msg == "" {
		msg = "not found"
	}
	return &APIError{Code: 404, Message: msg}
}

func NewBadRequestError(msg string) *APIError {
	if msg == "" {
		msg = "bad request"
	}
	return &APIError{Code: 400, Message: msg}
}

func NewInternalServerError(msg string) *APIError {
	if msg == "" {
		msg = "internal server error"
	}
	return &APIError{Code: 500, Message: msg}
}
