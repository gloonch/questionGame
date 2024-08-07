package httpserver

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"questionGame/config"
	"questionGame/service/authservice"
	"questionGame/service/userservice"
)

type Server struct {
	config  config.Config
	authSvc authservice.Service
	userSvc userservice.Service
}

func New(config config.Config, authSvc authservice.Service, userSvc userservice.Service) Server {
	return Server{
		config:  config,
		authSvc: authSvc,
		userSvc: userSvc,
	}
}

func (s Server) Serve() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health-check", s.healthCheck)

	e.POST("/users/register", s.userRegister)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.config.HTTPServer.Port)))

}
