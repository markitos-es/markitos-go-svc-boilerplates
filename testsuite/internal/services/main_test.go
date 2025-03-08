package services_test

import (
	"os"
	"testing"

	"github.com/markitos/markitos-svc-boilerplate/internal/domain"
)

type MockSpyBoilerRepository struct {
	LastCreatedBoilerName *string
	LastDeleteBoilerId    *string
}

func NewMockSpyBoilerRepository() *MockSpyBoilerRepository {
	return &MockSpyBoilerRepository{
		LastCreatedBoilerName: nil,
		LastDeleteBoilerId:    nil,
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

func (m *MockSpyBoilerRepository) Delete(id *domain.BoilerplateId) error {
	value := id.Value()
	m.LastDeleteBoilerId = &value

	return nil
}

func (m *MockSpyBoilerRepository) DeleteHaveBeenCalledWith(boilerId *string) bool {
	var result bool = m.LastDeleteBoilerId != nil && *m.LastDeleteBoilerId == *boilerId

	m.LastDeleteBoilerId = nil

	return result
}

func (m *MockSpyBoilerRepository) Update(boiler *domain.Boilerplate) error {
	return nil
}

func (m *MockSpyBoilerRepository) One(id *domain.BoilerplateId) (*domain.Boilerplate, error) {
	return domain.NewRandomBoilerplate(), nil
}

var repository *MockSpyBoilerRepository

func TestMain(m *testing.M) {
	repository = NewMockSpyBoilerRepository()

	os.Exit(m.Run())
}
