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

//func userProfileHandler(writer http.ResponseWriter, req *http.Request) {
//	if req.Method != http.MethodGet {
//		fmt.Fprintf(writer, `{"error": "invalid method"}`)
//	}
//
//	authSvc := authservice.New(JwtSignKey, AccessTokenSubject, RefreshTokenSubject,
//		AccessTokenExpireDuration, RefreshTokenExpireDuration)
//
//	authToken := req.Header.Get("Authorization")
//	claims, err := authSvc.ParseToken(authToken)
//	if err != nil {
//		fmt.Fprintf(writer, `{"error": "invalid token"}`)
//	}
//
//	mysqlRepo := mysql.New()
//	userSvc := userservice.New(authSvc, mysqlRepo)
//
//	response, err := userSvc.GetProfile(userservice.ProfileRequest{UserID: claims.UserID})
//	if err != nil {
//		writer.Write([]byte(
//			fmt.Sprintf(`{"error": "%s"}`, err.Error()),
//		))
//
//		return
//	}
//
//	data, err := json.Marshal(response)
//	if err != nil {
//		writer.Write([]byte(
//			fmt.Sprintf(`{"error": "%s"}`, err.Error()),
//		))
//
//		return
//	}
//
//	writer.Write(data)
//
//}
//
//func healthCheckHandler(writer http.ResponseWriter, req *http.Request) {
//	fmt.Fprintf(writer, `{"message": "everything is ok"}`)
//}
//
//func userLoginHandler(writer http.ResponseWriter, req *http.Request) {
//	if req.Method != http.MethodPost {
//		fmt.Fprintf(writer, `{"error": "invalid method"}`)
//	}
//
//	data, err := io.ReadAll(req.Body)
//	if err != nil {
//		writer.Write([]byte(
//			fmt.Sprintf(`{"error": "%s"}`, err.Error()),
//		))
//
//		return
//	}
//
//	var lReq userservice.LoginRequest
//	err = json.Unmarshal(data, &lReq)
//	if err != nil {
//		writer.Write([]byte(
//			fmt.Sprintf(`{"error": "%s"}`, err.Error()),
//		))
//
//		return
//	}
//
//	authSvc := authservice.New(JwtSignKey, AccessTokenSubject, RefreshTokenSubject,
//		AccessTokenExpireDuration, RefreshTokenExpireDuration)
//
//	mysqlRepo := mysql.New()
//	userSvc := userservice.New(authSvc, mysqlRepo)
//
//	response, lErr := userSvc.Login(lReq)
//	if lErr != nil {
//		writer.Write([]byte(
//			fmt.Sprintf(`{"error": "%s"}`, lErr.Error()),
//		))
//
//		return
//	}
//
//	data, err = json.Marshal(response)
//	if err != nil {
//		writer.Write([]byte(
//			fmt.Sprintf(`{"error": "%s"}`, err.Error()),
//		))
//
//		return
//	}
//
//	writer.Write(data)
//
//}

func setupServices(cfg config.Config) (authservice.Service, userservice.Service) {
	authSvc := authservice.New(cfg.Auth)
	MySQLRepo := mysql.New(cfg.Mysql)
	userSvc := userservice.New(authSvc, MySQLRepo)

	return authSvc, userSvc
}
