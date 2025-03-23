package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markitos-es/markitos-svc-boilerplates/internal/domain"
	"github.com/markitos-es/markitos-svc-boilerplates/internal/services"
)

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
		var code int = http.StatusBadRequest
		if errors.Is(err, domain.ErrBoilerplateNotFound) {
			code = http.StatusNotFound
		}
		ctx.JSON(code, errorResonses(err))
		return
	}

	ctx.JSON(http.StatusOK, response.Data)
}
