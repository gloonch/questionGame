package main

import (
	"questionGame/config"
	"questionGame/delivery/httpserver"
	"questionGame/repository/mysql"
	"questionGame/service/authservice"
	"questionGame/service/userservice"
	"time"
)

const (
	JwtSignKey                 = "jwt_secret"
	AccessTokenSubject         = "ac"
	RefreshTokenSubject        = "rt"
	AccessTokenExpireDuration  = time.Hour * 24
	RefreshTokenExpireDuration = time.Hour * 24
)

func main() {

	cfg := config.Config{
		Auth: authservice.Config{
			SignKey:               JwtSignKey,
			AccessExpirationTime:  AccessTokenExpireDuration,
			RefreshExpirationTime: RefreshTokenExpireDuration,
			AccessSubject:         AccessTokenSubject,
			RefreshSubject:        RefreshTokenSubject,
		},
		HTTPServer: config.HTTPServer{Port: 8080},
		Mysql: mysql.Config{
			Username: "question",
			Password: "question7",
			Port:     3306,
			Host:     "localhost",
			DBName:   "question_db",
		},
	}

	authSvc, userSvc := setupServices(cfg)
	server := httpserver.New(cfg, authSvc, userSvc)
	server.Serve()

}

func setupServices(cfg config.Config) (authservice.Service, userservice.Service) {
	authSvc := authservice.New(cfg.Auth)
	MySQLRepo := mysql.New(cfg.Mysql)
	userSvc := userservice.New(authSvc, MySQLRepo)

	return authSvc, userSvc
}
