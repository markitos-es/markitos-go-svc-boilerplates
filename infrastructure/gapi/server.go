package gapi

import (
	"github.com/markitos-es/markitos-svc-boilerplates/internal/domain"
)

type Server struct {
	UnimplementedBoilerplateServiceServer
	address    string
	repository domain.BoilerplateRepository
}

func NewServer(address string, repository domain.BoilerplateRepository) *Server {
	apiGRPC := &Server{
		address:    address,
		repository: repository,
	}

	return apiGRPC
}

func (s *Server) Repository() domain.BoilerplateRepository {
	return s.repository
}
