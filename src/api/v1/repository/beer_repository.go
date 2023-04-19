package repository

import (
	"karhub.backend.developer.test/src/api/v1/model"
)

type BeerRepository interface {
	GetByID(id int) (model.Beer, error)
	GetAll() ([]model.Beer, error)
	Create(data model.Beer) (model.Beer, error)
	Update(id int, data model.Beer) (model.Beer, error)
	Delete(id int) error
}
