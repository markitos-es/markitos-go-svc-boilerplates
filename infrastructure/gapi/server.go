package gapi

import (
	"errors"
	"strings"

	"markitos-svc-boilerplates/internal/domain"

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
	// Default error code
	var code codes.Code = codes.Internal

	// Map domain errors to appropriate gRPC status codes
	switch {
	case errors.Is(err, domain.ErrBoilerplateNotFound):
		code = codes.NotFound
	case errors.Is(err, domain.ErrInvalidBoilerplateId),
		errors.Is(err, domain.ErrInvalidBoilerplateName),
		errors.Is(err, domain.ErrInvalidPageNumber),
		errors.Is(err, domain.ErrInvalidPageSize),
		strings.Contains(err.Error(), "invalid"),
		strings.Contains(err.Error(), "Invalid"):
		code = codes.InvalidArgument
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
