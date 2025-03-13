package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markitos/markitos-svc-boilerplate/internal/domain"
	"github.com/markitos/markitos-svc-boilerplate/internal/services"
)

func (s Server) create(ctx *gin.Context) {
	var request services.BoilerplateCreateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return
	}

	var service services.BoilerplateCreateService = services.NewBoilerplateCreateService(s.repository)
	response, err := service.Do(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (s Server) delete(ctx *gin.Context) {
	id, err := s.getQueryParam(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return
	}

	request := services.BoilerplateDeleteRequest{Id: *id}
	var service services.BoilerplateDeleteService = services.NewBoilerplateDeleteService(s.repository)
	if err := service.Do(request); err != nil {
		var code int = http.StatusBadRequest
		if errors.Is(err, domain.BoilerplateNotFoundError) {
			code = http.StatusNotFound
		}
		ctx.JSON(code, errorResonses(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (s Server) one(ctx *gin.Context) {
	id, err := s.getQueryParam(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return
	}

	request := services.BoilerplateOneRequest{Id: *id}
	var service services.BoilerplateOneService = services.NewBoilerplateOneService(s.repository)
	response, err := service.Do(request)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResonses(err))
		return
	}

	ctx.JSON(http.StatusOK, response.Data)
}

func (s Server) getQueryParam(ctx *gin.Context, paramName string) (*string, error) {
	response := ctx.Param(paramName)
	if response == "" {
		return nil, errors.New(paramName + " is required")

	}

	return &response, nil
}
