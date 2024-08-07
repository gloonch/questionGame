package config

import (
	"questionGame/repository/mysql"
	"questionGame/service/authservice"
)

type HTTPServer struct {
	Port int
}

type Config struct {
	Auth       authservice.Config
	HTTPServer HTTPServer
	Mysql      mysql.Config
}
