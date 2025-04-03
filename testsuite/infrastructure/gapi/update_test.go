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
	boilerplate := createPersistedRandomBoilerplate()
	updatedName := boilerplate.Name + " UPDATED"

	resp, err := grpcClient.UpdateBoilerplate(ctx, &gapi.UpdateBoilerplateRequest{
		Id:   boilerplate.Id,
		Name: updatedName,
	})

	assert.NoError(t, err)
	assert.Equal(t, boilerplate.Id, resp.Updated)

	getResp, _ := grpcClient.GetBoilerplate(ctx, &gapi.GetBoilerplateRequest{Id: boilerplate.Id})
	assert.Equal(t, updatedName, getResp.Name)

	deletePersistedRandomBoilerplate(boilerplate.Id)
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
