package services

import "github.com/markitos/markitos-svc-boilerplate/internal/domain"

type BoilerplateUpdateRequest struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type BoilerplateUpdateService struct {
	Repository domain.BoilerplateRepository
}

func NewBoilerplateUpdateService(repository domain.BoilerplateRepository) BoilerplateUpdateService {
	return BoilerplateUpdateService{
		Repository: repository,
	}
}

func (s BoilerplateUpdateService) Do(request BoilerplateUpdateRequest) error {
	boilerId, err := domain.NewBoilerplateId(request.Id)
	if err != nil {
		return err
	}

	boiler, err := s.Repository.One(boilerId)
	if err != nil {
		return err
	}

	boiler.Name = request.Name

	return s.Repository.Update(boiler)
}
