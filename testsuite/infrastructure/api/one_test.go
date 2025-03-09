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

func TestBoilerplateOneHandler_Success(t *testing.T) {
	var boiler *domain.Boilerplate = createPersistedRandomBoilerplate()

	recorder := httptest.NewRecorder()
	requestBody, _ := json.Marshal(services.BoilerplateOneRequest{
		Id: boiler.Id,
	})
	request, _ := http.NewRequest(http.MethodGet, "/v1/boilerplates/"+boiler.Id, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	RESTRouter().ServeHTTP(recorder, request)

	var response map[string]any
	json.NewDecoder(recorder.Body).Decode(&response)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, response["name"].(string), boiler.Name)
	assert.Equal(t, response["id"].(string), boiler.Id)

	deletePersisteRandomBoilerplate(response["id"].(string))
}
