package gapi_test

import (
	"testing"

	"github.com/markitos-es/markitos-svc-boilerplates/infrastructure/gapi"
	internal_test "github.com/markitos-es/markitos-svc-boilerplates/testsuite/internal"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestBoilerplateCanCreate(t *testing.T) {
	boiler := internal_test.NewRandomOnlyNameBoilerplate()

	resp, err := grpcClient.CreateBoilerplate(ctx, &gapi.CreateBoilerplateRequest{
		Name: boiler.Name,
	})

	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Id)
	assert.Equal(t, boiler.Name, resp.Name)

	deletePersistedRandomBoilerplate(resp.Id)
}

func TestBoilerplateCantCreateWithoutName(t *testing.T) {
	_, err := grpcClient.CreateBoilerplate(ctx, &gapi.CreateBoilerplateRequest{
		Name: "",
	})

	assert.Error(t, err)
	st, ok := status.FromError(err)
	assert.True(t, ok)
	assert.Equal(t, codes.Internal, st.Code())
}

func TestBoilerplateCantCreateWithoutValidName(t *testing.T) {
	_, err := grpcClient.CreateBoilerplate(ctx, &gapi.CreateBoilerplateRequest{
		Name: "!!!!!invalid!!!name!!!",
	})

	assert.Error(t, err)
	st, ok := status.FromError(err)
	assert.True(t, ok)
	assert.Equal(t, codes.Internal, st.Code())
}
