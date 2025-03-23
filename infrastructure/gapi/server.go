package gapi

import (
	"errors"

	"github.com/markitos-es/markitos-svc-boilerplates/internal/domain"
	codes "google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (s *Server) GetGRPCCode(err error) codes.Code {
	var code codes.Code = codes.Internal
	if errors.Is(err, domain.ErrBoilerplateNotFound) {
		code = codes.NotFound
	}

	return code
}

func (s *Server) ToProtos(domainBoilerplates []*domain.Boilerplate) []*Boilerplate {
	var protoBoilerplates []*Boilerplate
	for _, boilerplate := range domainBoilerplates {
		protoBoilerplates = append(protoBoilerplates, s.ToProto(boilerplate))
	}

	return protoBoilerplates
}

func (s *Server) ToProto(domainBoilerplate *domain.Boilerplate) *Boilerplate {
	return &Boilerplate{
		Id:        domainBoilerplate.Id,
		Name:      domainBoilerplate.Name,
		CreatedAt: timestamppb.New(domainBoilerplate.CreatedAt),
		UpdatedAt: timestamppb.New(domainBoilerplate.UpdatedAt),
	}
}
