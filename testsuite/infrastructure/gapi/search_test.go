package gapi_test

import (
	"testing"

	"markitos-svc-boilerplates/infrastructure/gapi"
	"markitos-svc-boilerplates/internal/domain"
	"markitos-svc-boilerplates/testsuite/infrastructure/testdb"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCanSearchWithPattern(t *testing.T) {
	pattern := domain.RandomString(10)

	var ids []string
	for i := 0; i < 5; i++ {
		id := domain.UUIDv4()
		ids = append(ids, id)
		name := pattern + domain.RandomPersonalName()
		boilerplate, _ := domain.NewBoilerplate(id, name)

		testdb.GetRepository().Create(boilerplate)
	}

	resp, err := grpcClient.SearchBoilerplates(ctx, &gapi.SearchBoilerplatesRequest{
		SearchTerm: pattern,
		PageNumber: 1,
		PageSize:   6,
	})

	assert.NoError(t, err)
	assert.Equal(t, 5, len(resp.Boilerplates))

	foundCount := 0
	for _, b := range resp.Boilerplates {
		for _, id := range ids {
			if b.Id == id {
				foundCount++
				break
			}
		}
	}
	assert.Equal(t, 5, foundCount)

	for _, id := range ids {
		deletePersistedRandomBoilerplate(id)
	}
}

func TestCantSearchWithInvalidOptionalPage(t *testing.T) {
	_, err := grpcClient.SearchBoilerplates(ctx, &gapi.SearchBoilerplatesRequest{
		PageNumber: -1,
		PageSize:   10,
	})

	assert.Error(t, err)
	st, ok := status.FromError(err)
	assert.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, st.Code())
}
