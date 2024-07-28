package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"questionGame/repository/mysql"
	"questionGame/service/userservice"
)

const (
	JwtSignKey = "jwt_secret"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/health-check", healthCheckHandler)
	mux.HandleFunc("/users/register", userRegisterHandler)
	mux.HandleFunc("/users/login", userLoginHandler)
	mux.HandleFunc("/users/profile", userProfileHandler)

	log.Println("server is listening on port 8080...")
	server := http.Server{Addr: ":8080", Handler: mux}
	log.Fatal(server.ListenAndServe())
}

func userProfileHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		fmt.Fprintf(writer, `{"error": "invalid method"}`)
	}

	//jwtToken := sessionID := req.Header.Get("Authorization")
	// validate jtw token and retrieve userID from token payload

	pReq := userservice.ProfileRequest{UserID: 0}

	data, err := io.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error()),
		))
	}

	err = json.Unmarshal(data, &pReq)
	if err != nil {
		writer.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error()),
		))

		return
	}

	mysqlRepo := mysql.New()
	userSvc := userservice.New(mysqlRepo, JwtSignKey)

	response, err := userSvc.GetProfile(pReq)
	if err != nil {
		writer.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error()),
		))

		return
	}

	data, err = json.Marshal(response)
	if err != nil {
		writer.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error()),
		))

		return
	}

	writer.Write(data)

}

func userRegisterHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		fmt.Fprintf(writer, `{"error": "invalid method"}`)
	}

	data, err := io.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error()),
		))

		return
	}

	var uReq userservice.RegisterRequest
	err = json.Unmarshal(data, &uReq)
	if err != nil {
		writer.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error()),
		))

		return
	}

	mysqlRepo := mysql.New()
	userSvc := userservice.New(mysqlRepo, JwtSignKey)

	_, err = userSvc.Register(uReq)
	if err != nil {
		writer.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error()),
		))

		return
	}

	writer.Write([]byte(`{"message": "user created successfully"}`))

}

func healthCheckHandler(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, `{"message": "everything is ok"}`)
}

func userLoginHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		fmt.Fprintf(writer, `{"error": "invalid method"}`)
	}

	data, err := io.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error()),
		))

		return
	}

	var lReq userservice.LoginRequest
	err = json.Unmarshal(data, &lReq)
	if err != nil {
		writer.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error()),
		))

		return
	}

	mysqlRepo := mysql.New()
	userSvc := userservice.New(mysqlRepo, JwtSignKey)

	response, lErr := userSvc.Login(lReq)
	if lErr != nil {
		writer.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, lErr.Error()),
		))

		return
	}

	data, err = json.Marshal(response)
	if err != nil {
		writer.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error()),
		))

		return
	}

	writer.Write(data)

}

//func testUserMysqlRepo() {
//
//	mysqlRepo := mysql.New()
//
//	mysql.New()
//	mysql.MySQLDB{}
//	inja faqat db.go ro migire
//
//	vali az mysqlRepo mituni method haro call koni ke tuye user.go e
//	mysqlRepo.IsPhoneNumberUnique()
//
//	createdUser, err := mysqlRepo.Register(entity.User{
//		ID:          0,
//		PhoneNumber: "0912",
//		Name:        "Mahdi Hadian",
//	})
//	if err != nil {
//		fmt.Println("register user: ", err)
//	} else {
//		fmt.Println("created user: ", createdUser)
//	}
//
//	isUnique, err := mysqlRepo.IsPhoneNumberUnique(createdUser.PhoneNumber)
//	if err != nil {
//		fmt.Println("unique err")
//	}
//	fmt.Println("isUnique: ", isUnique)
//}
