package database_test

import (
	"log"
	"testing"
	"time"

	"github.com/markitos/markitos-svc-boilerplate/infrastructure/configuration"
	"github.com/markitos/markitos-svc-boilerplate/infrastructure/database"
	"github.com/markitos/markitos-svc-boilerplate/internal/domain"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestBoilerCreate(t *testing.T) {
	db := setupTestDB()
	repository := database.NewBoilerPostgresRepository(db)

	boiler := domain.NewBoilerplate(domain.UUIDv4(), "Hello, World!")
	err := repository.Create(boiler)
	require.NoError(t, err)

	var result domain.Boilerplate
	err = db.First(&result, "id = ?", boiler.Id).Error
	require.NoError(t, err)
	require.Equal(t, boiler.Id, result.Id)
	require.Equal(t, boiler.Name, result.Name)
	require.WithinDuration(t, boiler.CreatedAt, result.CreatedAt, time.Second)
	require.WithinDuration(t, boiler.UpdatedAt, result.UpdatedAt, time.Second)

	db.Delete(&result)
}

func setupTestDB() *gorm.DB {
	config, err := configuration.LoadConfiguration("../../..")
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open(postgres.Open(config.DatabaseDsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&domain.Boilerplate{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
