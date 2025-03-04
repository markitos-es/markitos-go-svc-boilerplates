package service_test

import (
	"testing"

	"github.com/markitos/markitos-svc-boilerplate/internal/domain"
	"github.com/markitos/markitos-svc-boilerplate/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestCanCreateAUser(t *testing.T) {
	var name *domain.BoilerplateName = domain.GetBoilerplateNameRandom()
	var request services.BoilerplateCreateRequest = services.BoilerplateCreateRequest{
		Name: name.Value(),
	}

	var service services.BoilerplateCreateService = services.NewBoilerplateCreateService(repository)
	response, err := service.Do(request)

	assert.Nil(t, err)
	assert.True(t, repository.CreateHaveBeenCalledWith(&request.Name))
	assert.Equal(t, response.Name, request.Name)
	assert.NotEmpty(t, response.Id)
}
