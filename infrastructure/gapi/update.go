package gapi

import (
	context "context"

	"github.com/markitos-es/markitos-svc-boilerplates/internal/services"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateBoilerplate(ctx context.Context, in *UpdateBoilerplateRequest) (*UpdateBoilerplateResponse, error) {
	var service services.BoilerplateUpdateService = services.NewBoilerplateUpdateService(s.repository)
	var request services.BoilerplateUpdateRequest = services.BoilerplateUpdateRequest{
		Id:   in.Id,
		Name: in.Name,
	}
	if err := service.Do(request); err != nil {
		return nil, status.Error(s.GetGRPCCode(err), err.Error())
	}

	return &UpdateBoilerplateResponse{
		Updated: request.Id,
	}, nil
}
