package controllers

import (
	"github.com/CALLlA-74/cashing/internal/config"
	"github.com/CALLlA-74/cashing/internal/domain/Cassette/dto"
)
import "github.com/gin-gonic/gin"

type ChangerUC interface {
	ChangeMoney(req dto.ChangeMoneyReq) (dto.ChangingResult, error)
}

type Server struct {
	uc ChangerUC
}

func NewServer(uc ChangerUC) *Server {
	return &Server{uc: uc}
}

func (s *Server) RegisterRoutes(router gin.IRouter) {
	router.StaticFile(config.RootPath, config.HtmlPath)
	router.POST(config.FindChangingPath, s.changeMoney)
}
