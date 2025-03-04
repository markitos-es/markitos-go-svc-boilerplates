package domain_test

import (
	"testing"

	"github.com/markitos/markitos-svc-boilerplate/internal/domain"
)

func TestCanCreateValidBoilerplateName(t *testing.T) {
	validNames := []string{
		"ValidName",
		"AnotherValidName",
		"Valid Name With Spaces",
		"Short",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
		"InvalidNameWithMoreThanOneHundredCharactersInvalidNameWithMoreThanOneHundredCharactersInvalidNameWithMoreThanOneHundredCharacters"}
	for _, name := range validNames {
		if _, err := domain.NewBoilerplateName(name); err != nil {
			t.Errorf("Expected valid name, but got invalid: %s", name)
		}
	}

	invalidNames := []string{
		" InvalidName",
		"InvalidName ",
		"Invalid Name ",
		" Invalid Name",
		"Invalid@Name",
		"Invalid#Name",
		"Invalid123Name",
		"Invalid Name With Spaces ",
		" Invalid Name With Spaces",
		"Invalid Name With Spaces And Symbols!",
	}
	for _, name := range invalidNames {
		if _, err := domain.NewBoilerplateName(name); err == nil {
			t.Errorf("Expected valid name, but got invalid: %s", name)
		}
	}
}

func TestCanCreateARandomValidName(t *testing.T) {
	for range 20 {
		sut := domain.GetBoilerplateNameRandom()
		if sut == nil {
			t.Error("expected valid name, but got invalid value")
		}
	}
}
