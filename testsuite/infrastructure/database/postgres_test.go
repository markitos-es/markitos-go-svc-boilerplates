package database_test

import (
	"testing"
	"time"

	"github.com/markitos/markitos-svc-boilerplate/internal/domain"
	"github.com/markitos/markitos-svc-boilerplate/testsuite/infrastructure/testdb"
	"github.com/stretchr/testify/require"
)

func TestBoilerCreate(t *testing.T) {
	boiler, _ := domain.NewBoilerplate(domain.UUIDv4(), "Hello, World!")
	err := testdb.GetRepository().Create(boiler)
	require.NoError(t, err)

	var result domain.Boilerplate
	err = testdb.GetDB().First(&result, "id = ?", boiler.Id).Error
	require.NoError(t, err)
	require.Equal(t, boiler.Id, result.Id)
	require.Equal(t, boiler.Name, result.Name)
	require.WithinDuration(t, boiler.CreatedAt, result.CreatedAt, time.Second)
	require.WithinDuration(t, boiler.UpdatedAt, result.UpdatedAt, time.Second)

	testdb.GetDB().Delete(&result)
}
