package handler

var (
	// ErrNotFound - not found.
	ErrNotFound = NewError(404, "token.not_found")
	// ErrInvalidTimestamp - timestamp is has invalid format.
	ErrInvalidTimestamp = NewError(400, "timestamp.invalid_format")
	// ErrApiKey - Api-Key is not set or invalid.
	ErrApiKey = NewError(400, "apikey.required")
)

// Error represents a handler error. It provides methods for a HTTP status
// code and embeds the built-in error interface.
type Error interface {
	error
	Status() int
}

// StatusError represents error with http status code.
type StatusError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error allows StatusError to satisfy the error interface.
func (se StatusError) Error() string {
	return se.Message
}

// Status returns our HTTP status code.
func (se StatusError) Status() int {
	return se.Code
}

// NewError creates new error instance.
func NewError(code int, message string) Error {
	return &StatusError{
		Code:    code,
		Message: message,
	}
}
