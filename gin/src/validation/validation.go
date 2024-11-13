package validation

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func InitValidator() {
	validate = validator.New()
}

func FormatValidationError(err error) string {
	var errors []string
	for _, err := range err.(validator.ValidationErrors) {
		var message string
		switch err.Tag() {
		case "required":
			message = fmt.Sprintf("%s harus diisi", err.Field())
		case "email":
			message = fmt.Sprintf("%s harus valid email", err.Field())
		case "max":
			message = fmt.Sprintf("%s maksimal %s", err.Field(), err.Param())
		case "min":
			message = fmt.Sprintf("%s minimal %s", err.Field(), err.Param())
		case "len":
			message = fmt.Sprintf("%s harus berisi %s karakter", err.Field(), err.Param())
		case "number":
			message = fmt.Sprintf("%s harus berupa angka", err.Field())
		default:
			message = fmt.Sprintf("%s tidak valid", err.Field())
		}
		errors = append(errors, message)
	}
	return strings.Join(errors, ", ")
}
