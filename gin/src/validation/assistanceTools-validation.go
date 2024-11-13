package validation

import (
	"gin-project/src/model"
	"gin-project/src/response"
)

func ValidateCreateAssistanceTools(request model.CreateAssistanceToolsRequest) error {
	err := validate.Struct(request)
	if err != nil {
		return response.NewResponseError(400, FormatValidationError(err))
	}

	return nil
}
