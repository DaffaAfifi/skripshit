package validation

import (
	"gin-project/src/model"
	"gin-project/src/response"
)

func ValidateCreateUser(request model.CreateUserRequest) error {
	err := validate.Struct(request)
	if err != nil {
		return response.NewResponseError(400, FormatValidationError(err))
	}

	return nil
}

func ValidateLoginUser(request model.LoginUserRequest) error {
	err := validate.Struct(request)
	if err != nil {
		return response.NewResponseError(400, FormatValidationError(err))
	}

	return nil
}

func ValidateUpdateUser(request model.UpdateUserRequest) error {
	err := validate.Struct(request)
	if err != nil {
		return response.NewResponseError(400, FormatValidationError(err))
	}

	return nil
}
