package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/markitos/markitos-svc-boilerplate/internal/domain"
	"github.com/markitos/markitos-svc-boilerplate/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestBoilerCreateHandler_Success(t *testing.T) {
	recorder := httptest.NewRecorder()
	boiler := domain.NewRandomBoilerplate()
	requestBody, _ := json.Marshal(services.BoilerplateCreateRequest{
		Name: boiler.Name,
	})
	request, _ := http.NewRequest(http.MethodPost, "/v1/boilerplates", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	RESTRouter().ServeHTTP(recorder, request)

	var response map[string]any
	json.NewDecoder(recorder.Body).Decode(&response)
	responseId := response["id"].(string)
	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.Equal(t, boiler.Name, response["name"])
	assert.NotEmpty(t, responseId)

	deletePersisteRandomBoilerplate(responseId)
}
