package model

import (
	"gorm.io/gorm"
	"karhub.backend.developer.test/src/api/v1/domain"
)

type Beer struct {
	gorm.Model
	Style          string  `json:"style" gorm:"uniqueIndex"`
	MinTemperature float64 `json:"min_temperature"`
	MaxTemperature float64 `json:"max_temperature"`
}

func (b Beer) ToDomain() domain.Beer {
	return domain.Beer{
		ID:             int(b.ID),
		Style:          &b.Style,
		MinTemperature: &b.MinTemperature,
		MaxTemperature: &b.MaxTemperature,
	}
}

type Beers []Beer

func (b Beers) ToDomain() []domain.Beer {
	beers := make([]domain.Beer, 0)
	for _, beer := range b {
		beers = append(beers, beer.ToDomain())
	}

	return beers
}
