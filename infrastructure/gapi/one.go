package gapi

import (
	context "context"

	"github.com/markitos-es/markitos-svc-boilerplates/internal/services"
	status "google.golang.org/grpc/status"
)

func (s *Server) GetBoilerplate(ctx context.Context, in *GetBoilerplateRequest) (*GetBoilerplateResponse, error) {
	request := services.BoilerplateOneRequest{Id: in.Id}

	var service services.BoilerplateOneService = services.NewBoilerplateOneService(s.repository)
	response, err := service.Do(request)
	if err != nil {
		return nil, status.Error(s.GetGRPCCode(err), err.Error())

	}

	return &GetBoilerplateResponse{
		Id:   response.Data.Id,
		Name: response.Data.Name,
	}, nil
}
