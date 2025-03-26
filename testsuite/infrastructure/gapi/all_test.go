package gapi_test

import (
	"testing"

	"github.com/markitos-es/markitos-svc-boilerplates/infrastructure/gapi"
	"github.com/stretchr/testify/assert"
)

func TestBoilerplateCanListAllResources(t *testing.T) {
	boiler1 := createPersistedRandomBoilerplate()
	boiler2 := createPersistedRandomBoilerplate()

	resp, err := grpcClient.ListBoilerplates(ctx, &gapi.ListBoilerplatesRequest{})

	assert.NoError(t, err)
	assert.NotNil(t, resp.Boilerplates)

	found1, found2 := false, false
	for _, b := range resp.Boilerplates {
		if b.Id == boiler1.Id {
			found1 = true
		}
		if b.Id == boiler2.Id {
			found2 = true
		}
	}
	assert.True(t, found1, "First boilerplate not found in response")
	assert.True(t, found2, "Second boilerplate not found in response")

	deletePersistedRandomBoilerplate(boiler1.Id)
	deletePersistedRandomBoilerplate(boiler2.Id)
}
