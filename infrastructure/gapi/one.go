package gapi

import (
	context "context"
)

func (s *Server) GetBoilerplate(ctx context.Context, in *GetBoilerplateRequest) (*GetBoilerplateResponse, error) {
	return &GetBoilerplateResponse{}, nil
}
