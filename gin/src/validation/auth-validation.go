package validation

import (
	"gin-project/src/model"
	"gin-project/src/response"
)

func ValidateLogin(request model.LoginUserRequest) error {
	err := validate.Struct(request)
	if err != nil {
		return response.NewResponseError(400, FormatValidationError(err))
	}

	return nil
}
