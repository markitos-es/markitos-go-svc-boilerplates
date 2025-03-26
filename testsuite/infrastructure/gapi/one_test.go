package gapi_test

import (
	"testing"

	"github.com/markitos-es/markitos-svc-boilerplates/infrastructure/gapi"
	"github.com/markitos-es/markitos-svc-boilerplates/internal/domain"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestBoilerplateCanGetOne(t *testing.T) {
	boiler := createPersistedRandomBoilerplate()

	resp, err := grpcClient.GetBoilerplate(ctx, &gapi.GetBoilerplateRequest{
		Id: boiler.Id,
	})

	assert.NoError(t, err)
	assert.Equal(t, boiler.Name, resp.Name)
	assert.Equal(t, boiler.Id, resp.Id)

	deletePersistedRandomBoilerplate(resp.Id)
}

func TestBoilerplateCantGetInvalidId(t *testing.T) {
	_, err := grpcClient.GetBoilerplate(ctx, &gapi.GetBoilerplateRequest{
		Id: "an-invalid-id",
	})

	assert.Error(t, err)
	st, ok := status.FromError(err)
	assert.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, st.Code())
}

func TestBoilerplateCantGetValidIdButNonExistingResource(t *testing.T) {
	_, err := grpcClient.GetBoilerplate(ctx, &gapi.GetBoilerplateRequest{
		Id: domain.UUIDv4(),
	})

	assert.Error(t, err)
	st, ok := status.FromError(err)
	assert.True(t, ok)
	assert.Equal(t, codes.NotFound, st.Code())
}
