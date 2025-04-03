package api_test

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/markitos-es/markitos-svc-boilerplates/infrastructure/api"
	"github.com/markitos-es/markitos-svc-boilerplates/infrastructure/database"
	"github.com/markitos-es/markitos-svc-boilerplates/internal/domain"
	"github.com/markitos-es/markitos-svc-boilerplates/testsuite/infrastructure/testdb"
	internal_test "github.com/markitos-es/markitos-svc-boilerplates/testsuite/internal"
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
	var boilerplate *domain.Boilerplate = internal_test.NewRandomBoilerplate()
	testdb.GetRepository().Create(boilerplate)

	return boilerplate
}

func persistBoilerplate(boilerplate *domain.Boilerplate) {
	testdb.GetRepository().Create(boilerplate)
}

func deletePersisteRandomBoilerplate(boilerplateId string) {
	repository := database.NewBoilerplatePostgresRepository(testdb.GetDB())
	id, err := domain.NewBoilerplateId(boilerplateId)
	if err != nil {
		panic(err)
	}

	repository.Delete(id)
}
