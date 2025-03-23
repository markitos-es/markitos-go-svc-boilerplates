package gapi

import (
	context "context"

	"github.com/markitos-es/markitos-svc-boilerplates/internal/services"
	status "google.golang.org/grpc/status"
)

func (s *Server) DeleteBoilerplate(ctx context.Context, in *DeleteBoilerplateRequest) (*DeleteBoilerplateResponse, error) {
	request := services.BoilerplateDeleteRequest{Id: in.Id}

	var service services.BoilerplateDeleteService = services.NewBoilerplateDeleteService(s.repository)
	if err := service.Do(request); err != nil {
		return nil, status.Error(s.GetGRPCCode(err), err.Error())
	}

	return &DeleteBoilerplateResponse{}, nil
}
