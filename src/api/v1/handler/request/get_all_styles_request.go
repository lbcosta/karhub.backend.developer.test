package request

import "github.com/go-playground/validator/v10"

type GetAllStylesRequest struct {
	Temperature float64 `json:"temperature" validate:"required"`
}

func (r *GetAllStylesRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return validationErrors
	}

	return nil
}
