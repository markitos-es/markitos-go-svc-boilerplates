package domain

import (
	"fmt"
	"regexp"
)

type BoilerplateId struct {
	value string
}

func NewBoilerplateId(value string) (*BoilerplateId, error) {
	if IsUUIDv4(value) {
		return &BoilerplateId{value}, nil
	}

	return nil, ErrBoilerplateBadRequest
}

func (b *BoilerplateId) Value() string {
	return b.value
}

type BoilerplateName struct {
	value string
}

const BOILERPLATE_NAME_MAX_LENGTH = 100
const BOILERPLATE_NAME_MIN_LENGTH = 3

func NewBoilerplateName(value string) (*BoilerplateName, error) {
	if isValidBoilerplateName(value) {
		return &BoilerplateName{value}, nil
	}

	return nil, fmt.Errorf("invalid boilerplate name, only letters and spaces, start/end with letter, received: %s", value)
}

func isValidBoilerplateName(value string) bool {
	if len(value) > BOILERPLATE_NAME_MAX_LENGTH || len(value) < BOILERPLATE_NAME_MIN_LENGTH {
		return false
	}

	pattern := `^[a-zA-Z]{1}[a-zA-Z ]+[a-zA-Z]$|^[a-zA-Z]{1}$`
	matched, err := regexp.MatchString(pattern, value)
	if err != nil {
		return false
	}

	return matched
}
