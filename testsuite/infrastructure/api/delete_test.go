package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/markitos/markitos-svc-boilerplate/internal/domain"
	"github.com/markitos/markitos-svc-boilerplate/internal/services"
	"github.com/markitos/markitos-svc-boilerplate/testsuite/infrastructure/testdb"
	"github.com/stretchr/testify/assert"
)

func TestBoilerDeleteHandler_Success(t *testing.T) {
	var boiler *domain.Boilerplate = domain.NewRandomBoilerplate()
	testdb.GetRepository().Create(boiler)

	recorder := httptest.NewRecorder()
	requestBody, _ := json.Marshal(services.BoilerplateDeleteRequest{
		Id: boiler.Id,
	})
	request, _ := http.NewRequest(http.MethodDelete, "/v1/boilerplates/"+boiler.Id, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	RESTRouter().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}
