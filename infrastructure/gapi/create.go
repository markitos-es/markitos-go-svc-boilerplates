package gapi

import (
	context "context"

	"github.com/markitos-es/markitos-svc-boilerplates/internal/services"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

func (s *Server) CreateBoilerplate(ctx context.Context, in *CreateBoilerplateRequest) (*CreateBoilerplateResponse, error) {
	var service services.BoilerplateCreateService = services.NewBoilerplateCreateService(s.repository)
	var request services.BoilerplateCreateRequest = services.BoilerplateCreateRequest{
		Name: in.Name,
	}
	response, err := service.Do(request)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create boilerplate: %s", err)
	}

	return &CreateBoilerplateResponse{
		Id:   response.Id,
		Name: response.Name,
	}, nil
}
