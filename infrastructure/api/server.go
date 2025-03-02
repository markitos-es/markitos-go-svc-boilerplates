package api

import (
	"github.com/gin-gonic/gin"
	"github.com/markitos/markitos-svc-boilerplate/internal/domain"
)

type Server struct {
	address    string
	repository domain.BoilerplateRepository
	router     *gin.Engine
}

func NewServer(address string, repository domain.BoilerplateRepository) *Server {
	server := &Server{
		address:    address,
		repository: repository,
	}
	server.createRouter()

	return server
}

func (s *Server) Repository() domain.BoilerplateRepository {
	return s.repository
}

func (s *Server) Router() *gin.Engine {
	return s.router
}

func (s *Server) createRouter() {
	s.router = gin.Default()
	s.router.POST("/v1/boilerplates", s.create)
}

func (s *Server) Run() error {
	return s.router.Run(s.address)
}

func errorResonses(err error) gin.H {
	return gin.H{"error": err.Error()}
}
