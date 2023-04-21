package handler

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/suite"
	"io"
	"karhub.backend.developer.test/src/api/v1/domain"
	"karhub.backend.developer.test/src/api/v1/middleware"
	"karhub.backend.developer.test/src/api/v1/model"
	gormRepositories "karhub.backend.developer.test/src/api/v1/repository/gorm"
	httpRepositories "karhub.backend.developer.test/src/api/v1/repository/spotify"
	"karhub.backend.developer.test/src/api/v1/service"
	"karhub.backend.developer.test/src/config/database"
	"net/http/httptest"
	"strings"
	"testing"
)

type BeerHandlerTestSuite struct {
	suite.Suite
	SomeError error
	App       *fiber.App
	db        *database.PostgresDatabase
}

func (suite *BeerHandlerTestSuite) SetupTest() {
	db := database.PostgresDatabase(database.NewTestDatabase(database.TestDatabaseName))
	beerRepository := gormRepositories.NewGormBeerRepository(db)
	beerService := service.NewBeerService(beerRepository)

	playlistRepo := httpRepositories.NewSpotifyPlaylistRepository()
	playlistService := service.NewPlaylistService(playlistRepo)

	beerHandler := NewBeerHandler(beerService, playlistService)

	app := fiber.New()
	app.Get("/", beerHandler.HandleGetAll)
	app.Post("/", beerHandler.HandleCreate)
	app.Put("/:id", beerHandler.HandleUpdate)
	app.Delete("/:id", beerHandler.HandleDelete)

	app.Use(middleware.Authenticate)

	app.Get("/style", beerHandler.HandleGetClosestBeerStyles)

	suite.SomeError = errors.New("some error")
	suite.App = app
	suite.db = &db
}

func (suite *BeerHandlerTestSuite) TearDownTest() {
	database.TestDatabase{}.Destroy(database.TestDatabaseName)
}

func (suite *BeerHandlerTestSuite) Test_HandleGetAll_Success() {
	request := httptest.NewRequest("GET", "/", nil)
	request.Header.Set("Content-Type", "application/json")

	response, err := suite.App.Test(request, -1)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	var beers []domain.Beer
	err = json.Unmarshal(responseBody, &beers)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	suite.Equal(fiber.StatusOK, response.StatusCode)
	suite.Equal(9, len(beers))

	var beersInDb []model.Beer
	suite.db.Find(&beersInDb)

	suite.Equal(9, len(beersInDb))
}

func (suite *BeerHandlerTestSuite) Test_HandleGetAll_Error() {
	suite.db.Exec("DROP TABLE beers")

	request := httptest.NewRequest("GET", "/", nil)
	request.Header.Set("Content-Type", "application/json")

	response, err := suite.App.Test(request, -1)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	suite.Equal(fiber.StatusUnprocessableEntity, response.StatusCode)
}

func (suite *BeerHandlerTestSuite) Test_HandleCreate_Success() {
	reqBody := []byte(`{ "style": "teste", "min_temperature": 2, "max_temperature": 4 }`)

	request := httptest.NewRequest("POST", "/", strings.NewReader(string(reqBody)))
	request.Header.Set("Content-Type", "application/json")

	response, err := suite.App.Test(request, -1)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	var beer domain.Beer
	err = json.Unmarshal(responseBody, &beer)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	suite.Equal(fiber.StatusCreated, response.StatusCode)
	suite.Equal("teste", *beer.Style)
	suite.Equal(2.0, *beer.MinTemperature)
	suite.Equal(4.0, *beer.MaxTemperature)

	var beersInDb []model.Beer
	suite.db.Find(&beersInDb)

	suite.Equal(10, len(beersInDb))
}

func (suite *BeerHandlerTestSuite) Test_HandleCreate_BadRequest() {
	reqBody := []byte(`{ "style": "teste", "min_temperature": 2 }`)

	request := httptest.NewRequest("POST", "/", strings.NewReader(string(reqBody)))
	request.Header.Set("Content-Type", "application/json")

	response, err := suite.App.Test(request, -1)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	suite.Equal(fiber.StatusBadRequest, response.StatusCode)
}

func (suite *BeerHandlerTestSuite) Test_HandleCreate_AlreadyExists() {
	reqBody := []byte(`{ "style": "IPA", "min_temperature": -7, "max_temperature": 10 }`)

	request := httptest.NewRequest("POST", "/", strings.NewReader(string(reqBody)))
	request.Header.Set("Content-Type", "application/json")

	response, err := suite.App.Test(request, -1)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	suite.Equal(fiber.StatusUnprocessableEntity, response.StatusCode)
}

func (suite *BeerHandlerTestSuite) Test_HandleUpdate_Success() {
	reqBody := []byte(`{ "min_temperature": 2, "max_temperature": 4 }`)

	request := httptest.NewRequest("PUT", "/6", strings.NewReader(string(reqBody)))
	request.Header.Set("Content-Type", "application/json")

	response, err := suite.App.Test(request, -1)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	var beer domain.Beer
	err = json.Unmarshal(responseBody, &beer)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	suite.Equal(fiber.StatusOK, response.StatusCode)
	suite.Equal("IPA", *beer.Style)
	suite.Equal(2.0, *beer.MinTemperature)
	suite.Equal(4.0, *beer.MaxTemperature)

	var beersInDb model.Beer
	suite.db.First(&beersInDb, 6)

	suite.Equal("IPA", beersInDb.Style)
	suite.Equal(2.0, beersInDb.MinTemperature)
	suite.Equal(4.0, beersInDb.MaxTemperature)
}

func (suite *BeerHandlerTestSuite) Test_HandleUpdate_BadRequest() {
	reqBody := []byte(`{ "min_temperature": 5, "max_temperature": 4 }`)

	request := httptest.NewRequest("PUT", "/6", strings.NewReader(string(reqBody)))
	request.Header.Set("Content-Type", "application/json")

	response, err := suite.App.Test(request, -1)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	suite.Equal(fiber.StatusBadRequest, response.StatusCode)
}

func (suite *BeerHandlerTestSuite) Test_HandleUpdate_NotFound() {
	reqBody := []byte(`{ "min_temperature": 2, "max_temperature": 4 }`)

	request := httptest.NewRequest("PUT", "/100", strings.NewReader(string(reqBody)))
	request.Header.Set("Content-Type", "application/json")

	response, err := suite.App.Test(request, -1)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	suite.Equal(fiber.StatusNotFound, response.StatusCode)
}

func (suite *BeerHandlerTestSuite) Test_HandleUpdate_Error() {
	suite.db.Exec("DROP TABLE beers")

	reqBody := []byte(`{ "min_temperature": 2, "max_temperature": 4 }`)

	request := httptest.NewRequest("PUT", "/6", strings.NewReader(string(reqBody)))
	request.Header.Set("Content-Type", "application/json")

	response, err := suite.App.Test(request, -1)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	suite.Equal(fiber.StatusUnprocessableEntity, response.StatusCode)
}

func (suite *BeerHandlerTestSuite) Test_HandleDelete_Success() {
	request := httptest.NewRequest("DELETE", "/6", nil)

	response, err := suite.App.Test(request, -1)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	suite.Equal(fiber.StatusNoContent, response.StatusCode)

	var beersInDb []model.Beer
	suite.db.Find(&beersInDb)

	suite.Equal(8, len(beersInDb))
}

func (suite *BeerHandlerTestSuite) Test_HandleDelete_BadRequest() {
	request := httptest.NewRequest("DELETE", "/invalid", nil)

	response, err := suite.App.Test(request, -1)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	suite.Equal(fiber.StatusBadRequest, response.StatusCode)
}

func (suite *BeerHandlerTestSuite) Test_HandleDelete_NotFound() {
	request := httptest.NewRequest("DELETE", "/100", nil)

	response, err := suite.App.Test(request, -1)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	suite.Equal(fiber.StatusNotFound, response.StatusCode)
}

func (suite *BeerHandlerTestSuite) Test_HandleDelete_Error() {
	suite.db.Exec("DROP TABLE beers")

	request := httptest.NewRequest("DELETE", "/6", nil)

	response, err := suite.App.Test(request, -1)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	suite.Equal(fiber.StatusUnprocessableEntity, response.StatusCode)
}

func (suite *BeerHandlerTestSuite) Test_HandleGetClosestBeerStyles_Success() {
	reqBody := []byte(`{ "temperature": 3 }`)

	request := httptest.NewRequest("GET", "/style", strings.NewReader(string(reqBody)))
	request.Header.Set("Content-Type", "application/json")

	response, err := suite.App.Test(request, -1)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	var beerPlaylists []domain.BeerPlaylist
	err = json.Unmarshal(responseBody, &beerPlaylists)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	suite.Equal(fiber.StatusOK, response.StatusCode)
	suite.Equal(2, len(beerPlaylists))
	suite.Equal("IPA", beerPlaylists[0].BeerStyle)
	suite.Equal("Imperial Stouts", beerPlaylists[1].BeerStyle)
}

func (suite *BeerHandlerTestSuite) Test_HandleGetClosestBeerStyles_BadRequest() {
	reqBody := []byte(`{ "temperature": "invalid" }`)

	request := httptest.NewRequest("GET", "/style", strings.NewReader(string(reqBody)))
	request.Header.Set("Content-Type", "application/json")

	response, err := suite.App.Test(request, -1)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	suite.Equal(fiber.StatusBadRequest, response.StatusCode)
}

func (suite *BeerHandlerTestSuite) Test_HandleGetClosestBeerStyles_Error() {
	suite.db.Exec("DROP TABLE beers")

	reqBody := []byte(`{ "temperature": 3 }`)

	request := httptest.NewRequest("GET", "/style", strings.NewReader(string(reqBody)))
	request.Header.Set("Content-Type", "application/json")

	response, err := suite.App.Test(request, -1)
	if err != nil {
		suite.T().Fatalf("Failed to test: %s", err)
	}

	suite.Equal(fiber.StatusUnprocessableEntity, response.StatusCode)
}

func TestBeerHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(BeerHandlerTestSuite))
}
