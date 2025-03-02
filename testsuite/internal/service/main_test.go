package service_test

import (
	"os"
	"testing"

	"github.com/markitos/markitos-svc-boilerplate/internal/domain"
)

type MockSpyBoilerRepository struct {
	LastCreatedBoilerName *string
}

func NewMockSpyBoilerRepository() *MockSpyBoilerRepository {
	return &MockSpyBoilerRepository{
		LastCreatedBoilerName: nil,
	}
}

func (m *MockSpyBoilerRepository) Create(boiler *domain.Boilerplate) error {
	m.LastCreatedBoilerName = &boiler.Name

	return nil
}

func (m *MockSpyBoilerRepository) CreateHaveBeenCalledWith(boilerName *string) bool {
	var result bool = m.LastCreatedBoilerName != nil && *m.LastCreatedBoilerName == *boilerName

	m.LastCreatedBoilerName = nil

	return result
}

var repository *MockSpyBoilerRepository

func TestMain(m *testing.M) {
	repository = NewMockSpyBoilerRepository()

	os.Exit(m.Run())
}
