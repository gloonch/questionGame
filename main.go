package main

import (
	"fmt"
	"questionGame/entity"
	"questionGame/repository/mysql"
)

func main() {
	testUserMysqlRepo()
}

func testUserMysqlRepo() {

	mysqlRepo := mysql.New()

	createdUser, err := mysqlRepo.Register(entity.User{
		ID:          0,
		PhoneNumber: "0912",
		Name:        "Mahdi Hadian",
	})
	if err != nil {
		fmt.Println("register user: ", err)
	} else {
		fmt.Println("created user: ", createdUser)
	}

	isUnique, err := mysqlRepo.IsPhoneNumberUnique(createdUser.PhoneNumber)
	if err != nil {
		fmt.Println("unique err")
	}
	fmt.Println("isUnique: ", isUnique)
}
