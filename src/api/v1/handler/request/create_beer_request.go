package request

import (
	"github.com/go-playground/validator/v10"
	"karhub.backend.developer.test/src/api/v1/domain"
)

type CreateBeerRequest struct {
	Style          *string  `json:"style" validate:"required"`
	MinTemperature *float64 `json:"min_temperature" validate:"required"`
	MaxTemperature *float64 `json:"max_temperature" validate:"required"`
}

func (r *CreateBeerRequest) ToDomain() domain.Beer {
	return domain.Beer{
		Style:          r.Style,
		MinTemperature: r.MinTemperature,
		MaxTemperature: r.MaxTemperature,
	}
}

func (r *CreateBeerRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return validationErrors
	}

	return nil
}
