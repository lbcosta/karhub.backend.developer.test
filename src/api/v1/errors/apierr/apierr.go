package apierr

import "errors"

var (
	ErrBeerNotFound       = errors.New("beer of given id not found")
	ErrBeerAlreadyExists  = errors.New("beer already exists")
	ErrInvalidTemperature = errors.New("min temperature is higher than max temperature")
)
