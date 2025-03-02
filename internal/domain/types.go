package domain

import "fmt"

type BoilerplateId struct {
	value string
}

func NewBoilerplateId(value string) (*BoilerplateId, error) {
	if IsUUIDv4(value) {
		return &BoilerplateId{value}, nil
	}

	return nil, fmt.Errorf("invalid boilerplate id, must be an uuid v4, received: %s", value)
}
