package gapi

import (
	context "context"

	"github.com/markitos-es/markitos-svc-boilerplates/internal/domain"
	"github.com/markitos-es/markitos-svc-boilerplates/internal/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteBoilerplate(ctx context.Context, in *DeleteBoilerplateRequest) (*DeleteBoilerplateResponse, error) {
	if _, err := domain.NewBoilerplateId(in.Id); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	request := services.BoilerplateDeleteRequest{Id: in.Id}

	var service services.BoilerplateDeleteService = services.NewBoilerplateDeleteService(s.repository)
	if err := service.Do(request); err != nil {
		return nil, status.Error(s.GetGRPCCode(err), err.Error())
	}

	return &DeleteBoilerplateResponse{
		Deleted: request.Id,
	}, nil
}
