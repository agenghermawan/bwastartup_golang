package main

import (
	"bwastartup/user"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("connection is succesfully")

	var users []user.User
	db.Find(&users)

	length := len(users)
	fmt.Println(length)

	for _, user := range users {
		fmt.Println("name :", user.Name)
		fmt.Println("====================")
		fmt.Print("email :", user.Email)
	}
}
