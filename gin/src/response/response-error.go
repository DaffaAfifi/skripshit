package response

// ResponseError adalah tipe yang digunakan untuk menyimpan informasi error dengan status dan pesan.
type ResponseError struct {
	Status  int
	Message string
}

// Error mengimplementasikan interface error untuk ResponseError, mengembalikan pesan error.
func (e *ResponseError) Error() string {
	return e.Message
}

// NewResponseError adalah konstruktor untuk membuat objek ResponseError baru dengan status dan pesan.
func NewResponseError(status int, message string) *ResponseError {
	return &ResponseError{
		Status:  status,
		Message: message,
	}
}
