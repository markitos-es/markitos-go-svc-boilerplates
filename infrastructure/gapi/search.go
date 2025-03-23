package gapi

import (
	context "context"

	"github.com/markitos-es/markitos-svc-boilerplates/internal/services"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

func (s *Server) SearchBoilerplates(ctx context.Context, in *SearchBoilerplatesRequest) (*SearchBoilerplatesResponse, error) {
	var service services.BoilerplateSearchService = services.NewBoilerplateSearchService(s.repository)
	var request services.BoilerplateSearchRequest = services.BoilerplateSearchRequest{
		SearchTerm: in.SearchTerm,
		PageNumber: int(in.PageNumber),
		PageSize:   int(in.PageSize),
	}

	response, err := service.Do(request)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	domainBoilerplates := response.Data
	grpcCollection := s.ToProtos(domainBoilerplates)

	return &SearchBoilerplatesResponse{
		Boilerplates: grpcCollection,
	}, nil
}
