package gapi_test

import (
	"testing"

	"github.com/markitos-es/markitos-svc-boilerplates/infrastructure/gapi"
	"github.com/markitos-es/markitos-svc-boilerplates/internal/domain"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCanUpdateABoilerplate(t *testing.T) {
	boiler := createPersistedRandomBoilerplate()
	updatedName := boiler.Name + " UPDATED"

	resp, err := grpcClient.UpdateBoilerplate(ctx, &gapi.UpdateBoilerplateRequest{
		Id:   boiler.Id,
		Name: updatedName,
	})

	assert.NoError(t, err)
	assert.Equal(t, boiler.Id, resp.Updated)

	getResp, _ := grpcClient.GetBoilerplate(ctx, &gapi.GetBoilerplateRequest{Id: boiler.Id})
	assert.Equal(t, updatedName, getResp.Name)

	deletePersistedRandomBoilerplate(boiler.Id)
}

func TestCantUpdateANonExistingBoilerplate(t *testing.T) {
	randomId := domain.UUIDv4()
	_, err := grpcClient.UpdateBoilerplate(ctx, &gapi.UpdateBoilerplateRequest{
		Id:   randomId,
		Name: domain.RandomPersonalName(),
	})

	assert.Error(t, err)
	st, ok := status.FromError(err)
	assert.True(t, ok)
	assert.Equal(t, codes.NotFound, st.Code())
}

func TestCantUpdateAnInvalidBoilerplateId(t *testing.T) {
	_, err := grpcClient.UpdateBoilerplate(ctx, &gapi.UpdateBoilerplateRequest{
		Id:   "an-invalid-id-format",
		Name: domain.RandomPersonalName(),
	})

	assert.Error(t, err)
	st, ok := status.FromError(err)
	assert.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, st.Code())
}
