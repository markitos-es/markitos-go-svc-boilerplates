package gapi

import (
	context "context"

	"markitos-svc-boilerplates/internal/domain"
	"markitos-svc-boilerplates/internal/services"

	"google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

func (s *Server) GetBoilerplate(ctx context.Context, in *GetBoilerplateRequest) (*GetBoilerplateResponse, error) {
	if _, err := domain.NewBoilerplateId(in.Id); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

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
