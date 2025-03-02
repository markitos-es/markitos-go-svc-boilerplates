package api_test

import (
	"log"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/markitos/markitos-svc-boilerplate/infrastructure/api"
	"github.com/markitos/markitos-svc-boilerplate/infrastructure/configuration"
	"github.com/markitos/markitos-svc-boilerplate/infrastructure/database"
	"github.com/markitos/markitos-svc-boilerplate/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var boilerApiServer *api.Server
var boilerRepository domain.BoilerplateRepository

func TestMain(m *testing.M) {
	boilerApiServer = setupTestServer()
	boilerRepository = boilerApiServer.Repository()

	os.Exit(m.Run())
}

func setupTestServer() *api.Server {
	gin.SetMode(gin.TestMode)

	db := setupTestDB()
	return api.NewServer(":8080", database.NewBoilerPostgresRepository(db))
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
