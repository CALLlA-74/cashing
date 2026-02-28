package controllers

import (
	"github.com/CALLlA-74/cashing/internal/config"
	"github.com/CALLlA-74/cashing/internal/domain"
)
import "github.com/gin-gonic/gin"

type ChangerUC interface {
	ChangeMoney(req domain.ChangeMoneyReq) (domain.ChangingResult, error)
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
