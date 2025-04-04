package gapi

import (
	context "context"

	"markitos-svc-boilerplates/internal/services"

	status "google.golang.org/grpc/status"
)

func (s *Server) ListBoilerplates(ctx context.Context, in *ListBoilerplatesRequest) (*ListBoilerplatesResponse, error) {
	var service services.BoilerplateAllService = services.NewBoilerplateAllService(s.repository)
	response, err := service.Do()
	if err != nil {
		return nil, status.Error(s.GetGRPCCode(err), err.Error())
	}

	return &ListBoilerplatesResponse{
		Boilerplates: s.ToProtos(response.Data),
	}, nil
}
