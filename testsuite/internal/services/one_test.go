package services_test

import (
	"testing"

	"github.com/markitos/markitos-svc-boilerplate/internal/domain"
	"github.com/markitos/markitos-svc-boilerplate/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestCanGetAResource(t *testing.T) {
	var request services.BoilerplateOneRequest = services.BoilerplateOneRequest{
		Id: domain.UUIDv4(),
	}

	var service services.BoilerplateOneService = services.NewBoilerplateOneService(repository)
	boiler, err := service.Do(request)

	assert.Nil(t, err)
	assert.IsType(t, services.BoilerplateOneResponse{}, *boiler)
	assert.True(t, repository.OneHaveBeenCalledWith(&request.Id))
}
