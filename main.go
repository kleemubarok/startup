package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"startup/handler"
	"startup/user"
)

func main() {
	dsn := "mysql:mysql@tcp(127.0.0.1:3306)/startup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)

	router.Run()
	/*
		userInput := user.RegisterUserInput{}
		userInput.Name="Anak Baru"
		userInput.Occupation="Musisi"
		userInput.Email="blueband@minyak.co.id"
		userInput.Password="passwordgaalay"

		userService.RegisterUser(userInput)*/

}