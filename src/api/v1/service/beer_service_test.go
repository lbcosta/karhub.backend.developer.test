package service

import (
	"errors"
	"github.com/stretchr/testify/suite"
	"karhub.backend.developer.test/src/api/v1/domain"
	"karhub.backend.developer.test/src/api/v1/model"
	mocks "karhub.backend.developer.test/src/test/mocks/src/api/v1/repository"
	"testing"
)

type BeerServiceTestSuite struct {
	suite.Suite
	SomeError      error
	beerRepository *mocks.BeerRepository
	beerService    BeerService
}

func (suite *BeerServiceTestSuite) SetupTest() {
	suite.SomeError = errors.New("some error")
	suite.beerRepository = new(mocks.BeerRepository)
	suite.beerService = NewBeerService(suite.beerRepository)
}

func (suite *BeerServiceTestSuite) Test_GetAll_Success() {
	input := []model.Beer{
		{
			Style:          "Pilsen",
			MinTemperature: 4,
			MaxTemperature: 6,
		},
		{
			Style:          "IPA",
			MinTemperature: 5,
			MaxTemperature: 7,
		},
	}

	suite.beerRepository.On("GetAll").Return(input, nil)

	output, err := suite.beerService.GetAll()

	suite.NoError(err)
	suite.Equal(model.Beers(input).ToDomain(), output)

}

func (suite *BeerServiceTestSuite) Test_GetAll_Error() {
	suite.beerRepository.On("GetAll").Return(nil, suite.SomeError)

	output, err := suite.beerService.GetAll()

	suite.Error(err)
	suite.Nil(output)
}

func (suite *BeerServiceTestSuite) Test_Create_Success() {
	input := model.Beer{
		Style:          "Pilsen",
		MinTemperature: 4,
		MaxTemperature: 6,
	}

	suite.beerRepository.On("Create", input).Return(input, nil)

	output, err := suite.beerService.Create(input.ToDomain())

	suite.NoError(err)
	suite.Equal(input.ToDomain(), output)
}

func (suite *BeerServiceTestSuite) Test_Create_InvalidTemperature() {
	input := model.Beer{
		Style:          "Pilsen",
		MinTemperature: 4,
		MaxTemperature: 3,
	}

	output, err := suite.beerService.Create(input.ToDomain())

	suite.Error(err)
	suite.Equal(domain.Beer{}, output)
}

func (suite *BeerServiceTestSuite) Test_Create_Error() {
	input := model.Beer{
		Style:          "Pilsen",
		MinTemperature: 4,
		MaxTemperature: 6,
	}

	suite.beerRepository.On("Create", input).Return(model.Beer{}, suite.SomeError)

	output, err := suite.beerService.Create(input.ToDomain())

	suite.Error(err)
	suite.Equal(domain.Beer{}, output)
}

func (suite *BeerServiceTestSuite) Test_Update_Success() {
	const ID = 1
	inputModel := model.Beer{
		Style:          "Pilsen",
		MinTemperature: 4,
		MaxTemperature: 6,
	}
	inputModel.ID = ID

	updated := model.Beer{
		Style:          "IPA",
		MinTemperature: 4,
		MaxTemperature: 6,
	}
	updated.ID = ID

	suite.beerRepository.On("GetByID", ID).Return(inputModel, nil)
	suite.beerRepository.On("Update", ID, updated).Return(updated, nil)

	output, err := suite.beerService.Update(ID, updated.ToDomain())

	suite.NoError(err)
	suite.Equal(updated.ToDomain(), output)
}

func (suite *BeerServiceTestSuite) Test_Update_NotFound() {
	const ID = 1
	inputModel := model.Beer{
		Style:          "Pilsen",
		MinTemperature: 4,
		MaxTemperature: 6,
	}
	inputModel.ID = ID

	updated := model.Beer{
		Style:          "IPA",
		MinTemperature: 4,
		MaxTemperature: 6,
	}
	updated.ID = ID

	suite.beerRepository.On("GetByID", ID).Return(model.Beer{}, suite.SomeError)

	output, err := suite.beerService.Update(ID, updated.ToDomain())

	suite.Error(err)
	suite.Equal(domain.Beer{}, output)
}

func (suite *BeerServiceTestSuite) Test_Update_InvalidTemperature() {
	const ID = 1
	inputModel := model.Beer{
		Style:          "Pilsen",
		MinTemperature: 4,
		MaxTemperature: 6,
	}
	inputModel.ID = ID

	updated := model.Beer{
		Style:          "IPA",
		MinTemperature: 4,
		MaxTemperature: 3,
	}
	updated.ID = ID

	suite.beerRepository.On("GetByID", ID).Return(inputModel, nil)

	output, err := suite.beerService.Update(ID, updated.ToDomain())

	suite.Error(err)
	suite.Equal(domain.Beer{}, output)
}

func (suite *BeerServiceTestSuite) Test_Update_Error() {
	const ID = 1
	inputModel := model.Beer{
		Style:          "Pilsen",
		MinTemperature: 4,
		MaxTemperature: 6,
	}
	inputModel.ID = ID

	updated := model.Beer{
		Style:          "IPA",
		MinTemperature: 4,
		MaxTemperature: 6,
	}
	updated.ID = ID

	suite.beerRepository.On("GetByID", ID).Return(inputModel, nil)
	suite.beerRepository.On("Update", ID, updated).Return(model.Beer{}, suite.SomeError)

	output, err := suite.beerService.Update(ID, updated.ToDomain())

	suite.Error(err)
	suite.Equal(domain.Beer{}, output)
}

func (suite *BeerServiceTestSuite) Test_Delete_Success() {
	const ID = 1

	suite.beerRepository.On("GetByID", ID).Return(model.Beer{}, nil)
	suite.beerRepository.On("Delete", ID).Return(nil)

	err := suite.beerService.Delete(ID)

	suite.NoError(err)
}

func (suite *BeerServiceTestSuite) Test_Delete_NotFound() {
	const ID = 1

	suite.beerRepository.On("GetByID", ID).Return(model.Beer{}, suite.SomeError)

	err := suite.beerService.Delete(ID)

	suite.Error(err)
}

func (suite *BeerServiceTestSuite) Test_Delete_Error() {
	const ID = 1

	suite.beerRepository.On("GetByID", ID).Return(model.Beer{}, nil)
	suite.beerRepository.On("Delete", ID).Return(suite.SomeError)

	err := suite.beerService.Delete(ID)

	suite.Error(err)
}

func (suite *BeerServiceTestSuite) Test_GetClosestBeerStyles_Success() {
	const temperature = 5.0
	expected := []string{"Brown ale"}

	suite.beerRepository.On("GetAll").Return([]model.Beer{
		{
			Style:          "Pilsens",
			MinTemperature: -2,
			MaxTemperature: 4,
		},
		{
			Style:          "IPA",
			MinTemperature: -7,
			MaxTemperature: 10,
		},
		{
			Style:          "Brown ale",
			MinTemperature: 0,
			MaxTemperature: 14,
		},
	}, nil)

	output, err := suite.beerService.GetClosestBeerStyles(temperature)

	suite.NoError(err)
	suite.Equal(expected, output)
}

func (suite *BeerServiceTestSuite) Test_GetClosestBeerStyles_Error() {
	const temperature = 5.0

	suite.beerRepository.On("GetAll").Return([]model.Beer{}, suite.SomeError)

	output, err := suite.beerService.GetClosestBeerStyles(temperature)

	suite.Error(err)
	suite.Nil(output)
}

func (suite *BeerServiceTestSuite) Test_Create_SameTemperature_Error() {
	const ID = 1
	inputModel := model.Beer{
		Style:          "Pilsen",
		MinTemperature: 4,
		MaxTemperature: 4,
	}
	inputModel.ID = ID

	suite.beerRepository.On("Create", inputModel).Return(inputModel, nil)

	output, err := suite.beerService.Create(inputModel.ToDomain())

	suite.Error(err)
	suite.Equal(domain.Beer{}, output)
}

func (suite *BeerServiceTestSuite) Test_GetClosestBeerStyles_EmptyList() {
	const temperature = 5.0

	suite.beerRepository.On("GetAll").Return([]model.Beer{}, nil)

	output, err := suite.beerService.GetClosestBeerStyles(temperature)

	suite.NoError(err)
	suite.Empty(output)
}

func (suite *BeerServiceTestSuite) Test_GetClosestBeerStyles_MultipleBeers() {
	const temperature = 1.0
	expected := []string{"Weissbier", "Pilsens"}

	suite.beerRepository.On("GetAll").Return([]model.Beer{
		{
			Style:          "Weissbier",
			MinTemperature: -1,
			MaxTemperature: 3,
		},
		{
			Style:          "Pilsens",
			MinTemperature: -2,
			MaxTemperature: 4,
		},
		{
			Style:          "IPA",
			MinTemperature: -7,
			MaxTemperature: 10,
		},
	}, nil)

	output, err := suite.beerService.GetClosestBeerStyles(temperature)

	suite.NoError(err)
	suite.Equal(expected, output)
}

func TestBeerServiceTestSuite(t *testing.T) {
	suite.Run(t, new(BeerServiceTestSuite))
}
