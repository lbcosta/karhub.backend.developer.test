package repository

import (
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
	"karhub.backend.developer.test/src/api/v1/errors/apierr"
	"karhub.backend.developer.test/src/api/v1/model"
	"karhub.backend.developer.test/src/api/v1/repository"
	"karhub.backend.developer.test/src/config/database"
)

const uniqueConstraintViolationCode = "23505"

type GormBeerRepository struct {
	db database.PostgresDatabase
}

func NewGormBeerRepository(db database.PostgresDatabase) repository.BeerRepository {
	return GormBeerRepository{
		db: db,
	}
}

func (r GormBeerRepository) GetByID(id int) (model.Beer, error) {
	var beer model.Beer

	if err := r.db.First(&beer, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Beer{}, apierr.ErrBeerNotFound
		}

		return model.Beer{}, err
	}

	return beer, nil
}

func (r GormBeerRepository) GetAll() ([]model.Beer, error) {
	beers := make([]model.Beer, 0)

	if err := r.db.Find(&beers).Error; err != nil {
		return nil, err
	}

	return beers, nil
}

func (r GormBeerRepository) Create(data model.Beer) (model.Beer, error) {
	err := r.db.Create(&data).Error

	if err != nil {
		pgErr, ok := err.(*pgconn.PgError)
		if ok && pgErr.Code == uniqueConstraintViolationCode {
			return model.Beer{}, apierr.ErrBeerAlreadyExists
		}

		return model.Beer{}, err
	}

	return data, nil
}

func (r GormBeerRepository) Update(id int, data model.Beer) (model.Beer, error) {
	err := r.db.Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return model.Beer{}, err
	}

	return data, nil
}

func (r GormBeerRepository) Delete(id int) error {
	err := r.db.Delete(&model.Beer{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
