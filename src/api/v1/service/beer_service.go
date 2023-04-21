package service

import (
	"karhub.backend.developer.test/src/api/v1/domain"
	"karhub.backend.developer.test/src/api/v1/errors/apierr"
	"karhub.backend.developer.test/src/api/v1/model"
	"karhub.backend.developer.test/src/api/v1/repository"
	"math"
)

type BeerService struct {
	beerRepository repository.BeerRepository
}

func NewBeerService(beerRepository repository.BeerRepository) BeerService {
	return BeerService{beerRepository: beerRepository}
}

func (s *BeerService) GetAll() ([]domain.Beer, error) {
	beers, err := s.beerRepository.GetAll()
	if err != nil {
		return nil, err
	}

	beersModel := model.Beers(beers)

	return beersModel.ToDomain(), nil
}

func (s *BeerService) Create(data domain.Beer) (domain.Beer, error) {
	beer := model.Beer{
		Style:          *data.Style,
		MinTemperature: *data.MinTemperature,
		MaxTemperature: *data.MaxTemperature,
	}

	if beer.MinTemperature >= beer.MaxTemperature {
		return domain.Beer{}, apierr.ErrInvalidTemperature
	}

	createdBeer, err := s.beerRepository.Create(beer)
	if err != nil {
		return domain.Beer{}, err
	}

	return createdBeer.ToDomain(), nil
}

func (s *BeerService) Update(id int, data domain.Beer) (domain.Beer, error) {
	beer, err := s.beerRepository.GetByID(id)
	if err != nil {
		return domain.Beer{}, err
	}

	if !isTemperatureValid(&beer, &data) {
		return domain.Beer{}, apierr.ErrInvalidTemperature
	}

	updateFields(&beer, &data)

	updatedBeer, err := s.beerRepository.Update(id, beer)
	if err != nil {
		return domain.Beer{}, err
	}

	return updatedBeer.ToDomain(), nil
}

func (s *BeerService) Delete(id int) error {
	_, err := s.beerRepository.GetByID(id)
	if err != nil {
		return err
	}

	return s.beerRepository.Delete(id)
}

func (s *BeerService) GetClosestBeerStyles(temperature float64) ([]string, error) {
	beersModel, err := s.beerRepository.GetAll()
	if err != nil {
		return nil, err
	}

	beers := model.Beers(beersModel).ToDomain()

	closestBeers := findClosestBeers(temperature, beers)

	return closestBeers, nil
}

func updateFields(beer *model.Beer, data *domain.Beer) {
	if data.Style != nil {
		beer.Style = *data.Style
	}

	if data.MinTemperature != nil {
		beer.MinTemperature = *data.MinTemperature
	}

	if data.MaxTemperature != nil {
		beer.MaxTemperature = *data.MaxTemperature
	}
}

func isTemperatureValid(beer *model.Beer, data *domain.Beer) bool {
	min, max := beer.MinTemperature, beer.MaxTemperature

	if data.MinTemperature != nil {
		min = *data.MinTemperature
	}

	if data.MaxTemperature != nil {
		max = *data.MaxTemperature
	}

	return min < max
}

func findClosestBeers(temperature float64, beers []domain.Beer) []string {
	var minDelta *float64
	var targetBeers = make([]string, 0)

	for _, beer := range beers {
		delta := math.Abs(((*beer.MinTemperature + *beer.MaxTemperature) / 2) - temperature)

		if minDelta == nil {
			minDelta = &delta
			targetBeers = append(targetBeers, *beer.Style)
			continue
		}

		if delta < *minDelta {
			minDelta = &delta
			targetBeers = []string{*beer.Style}
		}

		if delta == *minDelta && !contains(*beer.Style, targetBeers) {
			minDelta = &delta
			targetBeers = append(targetBeers, *beer.Style)
		}
	}

	return targetBeers
}

func contains(elem string, array []string) bool {
	for _, beer := range array {
		if beer == elem {
			return true
		}
	}

	return false
}
