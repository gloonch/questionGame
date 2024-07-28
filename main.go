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

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/health-check", healthCheckHandler)
	mux.HandleFunc("/users/register", userRegisterHandler)
	mux.HandleFunc("/users/login", userLoginHandler)

	log.Println("server is listening on port 8080...")
	server := http.Server{Addr: ":8080", Handler: mux}
	log.Fatal(server.ListenAndServe())
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
	userSvc := userservice.New(mysqlRepo)

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
	userSvc := userservice.New(mysqlRepo)

	_, lErr := userSvc.Login(lReq)
	if lErr != nil {
		writer.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, lErr.Error()),
		))

		return
	}

	writer.Write([]byte(`{"message": "user credential is ok"}`))

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
