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

	return nil, fmt.Errorf("invalid boilerplate id, must be an uuid v4, received: %s", value)
}

type BoilerplateName struct {
	value string
}

func NewBoilerplateName(value string) (*BoilerplateName, error) {
	if isValidBoilerplateName(value) {
		return &BoilerplateName{value}, nil
	}

	return nil, fmt.Errorf("invalid boilerplate name, only letters and spaces, start/end with letter, received: %s", value)
}

func (b *BoilerplateName) Value() string {
	return b.value
}

func GetBoilerplateNameRandom() *BoilerplateName {
	result, _ := NewBoilerplateName(RandomPersonalName())

	return result
}

func isValidBoilerplateName(value string) bool {
	pattern := `^[a-zA-Z]{1}[a-zA-Z ]+[a-zA-Z]{1,}$`
	matched, _ := regexp.MatchString(pattern, value)

	return matched
}
