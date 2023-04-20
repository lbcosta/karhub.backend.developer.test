package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"karhub.backend.developer.test/src/api/v1/model"
	"karhub.backend.developer.test/src/config"
	"log"
	"os"
	"time"
)

var PingPostgres func() error

type PostgresDatabase struct {
	*gorm.DB
}

func NewPostgresDatabase() PostgresDatabase {
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s timezone=%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		"America/Sao_Paulo")

	connection := PostgresDatabase{}.connect(dsn)

	db, _ := connection.DB()
	PingPostgres = db.Ping

	connection.AutoMigrate(&model.Beer{})

	postgresDb := PostgresDatabase{
		DB: connection,
	}

	return postgresDb
}

func (PostgresDatabase) connect(dsn string) *gorm.DB {
	env := config.GetEnvironment()

	gormLogger := logger.Default.LogMode(logger.Silent)

	if env == config.DevelopmentEnvironment {
		gormLogger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Info,
				Colorful:      true,
			},
		)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: gormLogger})
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}

	return db
}
