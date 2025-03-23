package api

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func (s Server) getQueryParam(ctx *gin.Context, paramName string) (*string, error) {
	response := ctx.Param(paramName)
	if response == "" {
		return nil, errors.New(paramName + " is required")

	}

	return &response, nil
}
