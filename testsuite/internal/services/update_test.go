package services_test

import (
	"testing"

	"github.com/markitos/markitos-svc-boilerplate/internal/domain"
	"github.com/markitos/markitos-svc-boilerplate/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestCanUpdateABoiler(t *testing.T) {
	var request services.BoilerplateUpdateRequest = services.BoilerplateUpdateRequest{
		Id:   domain.UUIDv4(),
		Name: domain.RandomPersonalName(),
	}

	var service services.BoilerplateUpdateService = services.NewBoilerplateUpdateService(repository)
	err := service.Do(request)

	assert.Nil(t, err)
	assert.True(t, repository.UpdateHaveBeenCalledWith(request.Id, request.Name))
	assert.True(t, repository.UpdateHaveBeenCalledOneWith(request.Id))
}
