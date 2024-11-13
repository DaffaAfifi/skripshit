package response

type ResponseError struct {
	Status  int
	Message string
}

func (e *ResponseError) Error() string {
	return e.Message
}

func NewResponseError(status int, message string) *ResponseError {
	return &ResponseError{
		Status:  status,
		Message: message,
	}
}
