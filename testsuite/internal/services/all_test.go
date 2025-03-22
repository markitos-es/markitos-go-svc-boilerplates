package services_test

import (
	"testing"

	"github.com/markitos/markitos-svc-boilerplate/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestCanGetAllResources(t *testing.T) {
	var service services.BoilerplateAllService = services.NewBoilerplateAllService(repository)
	boiler, err := service.Do()

	assert.Nil(t, err)
	assert.IsType(t, services.BoilerplateAllResponse{}, *boiler)
	assert.True(t, repository.AllHaveBeenCalled())
}
