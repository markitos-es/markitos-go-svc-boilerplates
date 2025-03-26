package gapi

import (
	context "context"

	"github.com/markitos-es/markitos-svc-boilerplates/internal/services"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateBoilerplate(ctx context.Context, in *CreateBoilerplateRequest) (*CreateBoilerplateResponse, error) {
	var request services.BoilerplateCreateRequest = services.BoilerplateCreateRequest{
		Name: in.Name,
	}

	var service services.BoilerplateCreateService = services.NewBoilerplateCreateService(s.repository)
	response, err := service.Do(request)
	if err != nil {
		return nil, status.Error(s.GetGRPCCode(err), err.Error())
	}

	return &CreateBoilerplateResponse{
		Id:   response.Id,
		Name: response.Name,
	}, nil
}
