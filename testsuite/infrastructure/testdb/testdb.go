package testdb

import (
	"log"
	"sync"

	"github.com/markitos-es/markitos-svc-boilerplates/infrastructure/configuration"
	"github.com/markitos-es/markitos-svc-boilerplates/infrastructure/database"
	"github.com/markitos-es/markitos-svc-boilerplates/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbInstance   *gorm.DB
	dbOnce       sync.Once
	repoInstance domain.BoilerplateRepository
	repoOnce     sync.Once
)

func GetDB() *gorm.DB {
	dbOnce.Do(func() {
		config, err := configuration.LoadConfiguration("../../..")
		if err != nil {
			log.Fatalf("Failed to load configuration: %v", err)
		}

		db, err := gorm.Open(postgres.Open(config.DatabaseDsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		dbInstance = db
		dbInstance.AutoMigrate(&domain.Boilerplate{})
	})

	return dbInstance
}

func GetRepository() domain.BoilerplateRepository {
	repoOnce.Do(func() {
		db := GetDB()
		repo := database.NewBoilerPostgresRepository(db)
		repoInstance = &repo
	})

	return repoInstance
}
