package app

import (
	"fmt"
	"github.com/CALLlA-74/cashing/internal/config"
	"github.com/CALLlA-74/cashing/internal/controllers"
	"github.com/CALLlA-74/cashing/internal/domain"
	"github.com/gin-gonic/gin"
)

type App struct {
	cfg    *config.Config
	router *gin.Engine
	server *controllers.Server
}

func New(cfg *config.Config) *App {
	return &App{
		cfg:    cfg,
		router: gin.Default(),
		server: controllers.NewServer(domain.NewUC()),
	}
}

func (app *App) Start() {
	app.server.RegisterRoutes(app.router)
	_ = app.router.Run(fmt.Sprintf("%s:%s", app.cfg.HTTP.IP, app.cfg.HTTP.Port))
}
