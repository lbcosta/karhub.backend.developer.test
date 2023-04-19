package request

import (
	"github.com/go-playground/validator/v10"
	"karhub.backend.developer.test/src/api/v1/domain"
)

type UpdateBeerRequest struct {
	Style          *string  `json:"style,omitempty" validate:"omitempty"`
	MinTemperature *float64 `json:"min_temperature,omitempty" validate:"omitempty"`
	MaxTemperature *float64 `json:"max_temperature,omitempty" validate:"omitempty"`
}

func (r *UpdateBeerRequest) ToDomain() domain.Beer {
	return domain.Beer{
		Style:          r.Style,
		MinTemperature: r.MinTemperature,
		MaxTemperature: r.MaxTemperature,
	}
}

func (r *UpdateBeerRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return validationErrors
	}

	return nil
}
