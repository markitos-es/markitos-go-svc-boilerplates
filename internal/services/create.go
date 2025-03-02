package services

import "github.com/markitos/markitos-svc-boilerplate/internal/domain"

type BoilerplateCreateRequest struct {
	Name string
}

type BoilerplateCreateResponse struct {
	Id   string
	Name string
}

type BoilerplateCreateService struct {
	Repository domain.BoilerplateRepository
}

func NewBoilerplateCreateService(repository domain.BoilerplateRepository) BoilerplateCreateService {
	return BoilerplateCreateService{
		Repository: repository,
	}
}

func (s BoilerplateCreateService) Do(request BoilerplateCreateRequest) (*BoilerplateCreateResponse, error) {
	var id string = domain.UUIDv4()
	var boiler *domain.Boilerplate = domain.NewBoilerplate(id, request.Name)

	if err := s.Repository.Create(boiler); err != nil {
		return nil, err
	}

	return &BoilerplateCreateResponse{
		Id:   id,
		Name: request.Name,
	}, nil
}
