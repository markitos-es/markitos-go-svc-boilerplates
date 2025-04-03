package gapi_test

import (
	"testing"

	"github.com/markitos-es/markitos-svc-boilerplates/infrastructure/gapi"
	"github.com/markitos-es/markitos-svc-boilerplates/internal/domain"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestBoilerplateCanDelete(t *testing.T) {
	boilerplate := createPersistedRandomBoilerplate()

	resp, err := grpcClient.DeleteBoilerplate(ctx, &gapi.DeleteBoilerplateRequest{
		Id: boilerplate.Id,
	})

	assert.NoError(t, err)
	assert.Equal(t, boilerplate.Id, resp.Deleted)

	_, err = grpcClient.GetBoilerplate(ctx, &gapi.GetBoilerplateRequest{Id: boilerplate.Id})
	assert.Error(t, err)
}

func TestBoilerplateCantDeleteValidButNonExistingId(t *testing.T) {
	_, err := grpcClient.DeleteBoilerplate(ctx, &gapi.DeleteBoilerplateRequest{
		Id: domain.UUIDv4(),
	})

	assert.Error(t, err)
	st, ok := status.FromError(err)
	assert.True(t, ok)
	assert.Equal(t, codes.NotFound, st.Code())
}

func TestBoilerplateCantDeleteInvalidBoilerplateId(t *testing.T) {
	_, err := grpcClient.DeleteBoilerplate(ctx, &gapi.DeleteBoilerplateRequest{
		Id: "an-invalid-id-format",
	})

	assert.Error(t, err)
	st, ok := status.FromError(err)
	assert.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, st.Code())
}
