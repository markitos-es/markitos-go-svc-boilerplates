package api_test

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/markitos/markitos-svc-boilerplate/infrastructure/api"
	"github.com/markitos/markitos-svc-boilerplate/infrastructure/database"
	"github.com/markitos/markitos-svc-boilerplate/internal/domain"
	"github.com/markitos/markitos-svc-boilerplate/testsuite/infrastructure/testdb"
)

var boilerplatesApiServer *api.Server

func TestMain(m *testing.M) {
	setupRESTServer()
	os.Exit(m.Run())
}

func RESTServer() *api.Server {
	return boilerplatesApiServer
}

func setupRESTServer() {
	gin.SetMode(gin.TestMode)
	boilerplatesApiServer = api.NewServer(":8080", testdb.GetRepository())
}

func RESTRouter() *gin.Engine {
	return RESTServer().Router()
}

func createPersistedRandomBoilerplate() *domain.Boilerplate {
	var boiler *domain.Boilerplate = domain.NewRandomBoilerplate()
	testdb.GetRepository().Create(boiler)

	return boiler
}

func deletePersisteRandomBoilerplate(boilerId string) {
	repository := database.NewBoilerPostgresRepository(testdb.GetDB())
	id, err := domain.NewBoilerplateId(boilerId)
	if err != nil {
		panic(err)
	}

	repository.Delete(id)
}
