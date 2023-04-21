package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"karhub.backend.developer.test/src/api/v1/model"
	"log"
	"os"
)

const TestDatabaseName = "test.db"

type TestDatabase struct {
	*gorm.DB
}

func NewTestDatabase(databaseName string) TestDatabase {
	database := TestDatabase{}

	connection := database.connect(databaseName)

	connection.AutoMigrate(&model.Beer{})

	connection.Create([]model.Beer{
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
			Style:          "Weizenbier",
			MinTemperature: -4,
			MaxTemperature: 6,
		},
		{
			Style:          "Red ale",
			MinTemperature: -5,
			MaxTemperature: 5,
		},
		{
			Style:          "India pale ale",
			MinTemperature: -6,
			MaxTemperature: 7,
		},
		{
			Style:          "IPA",
			MinTemperature: -7,
			MaxTemperature: 10,
		},
		{
			Style:          "Dunkel",
			MinTemperature: -8,
			MaxTemperature: 2,
		},
		{
			Style:          "Imperial Stouts",
			MinTemperature: -10,
			MaxTemperature: 13,
		},
		{
			Style:          "Brown ale",
			MinTemperature: 0,
			MaxTemperature: 14,
		},
	})

	testDatabase := TestDatabase{
		DB: connection,
	}

	return testDatabase
}

func (TestDatabase) connect(dsn string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}

	return db
}

func (TestDatabase) Destroy(databaseName string) {
	os.Remove(databaseName)
}
