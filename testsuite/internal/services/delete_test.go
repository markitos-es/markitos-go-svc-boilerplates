package services_test

import (
	"testing"

	"github.com/markitos/markitos-svc-boilerplate/internal/domain"
	"github.com/markitos/markitos-svc-boilerplate/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestCanDeleteAUser(t *testing.T) {
	var request services.BoilerplateDeleteRequest = services.BoilerplateDeleteRequest{
		Id: domain.UUIDv4(),
	}

	var service services.BoilerplateDeleteService = services.NewBoilerplateDeleteService(repository)
	err := service.Do(request)
	assert.Nil(t, err)
	assert.True(t, repository.DeleteHaveBeenCalledWith(&request.Id))
}
