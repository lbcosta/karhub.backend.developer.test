package service

import (
	"karhub.backend.developer.test/src/api/v1/domain"
	"karhub.backend.developer.test/src/api/v1/errors/apierr"
	"karhub.backend.developer.test/src/api/v1/model"
	"karhub.backend.developer.test/src/api/v1/repository"
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

	if isTemperatureValid(&beer, &data) {
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

func (s *BeerService) GetAllStyles(temperature float64) ([]domain.Beer, error) {
	//styles, err := s.beerRepository.GetAllStyles(temperature)
	//if err != nil {
	//	return nil, err
	//}
	//
	//domainStyles := model.Beers(styles)

	return nil, nil
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
	return data.MinTemperature != nil && (*data.MinTemperature >= beer.MaxTemperature) || data.MaxTemperature != nil && (*data.MaxTemperature <= beer.MinTemperature)
}
