package main

import (
	"gorm-crud-project/config"
	"gorm-crud-project/handlers"
	"gorm-crud-project/kafka"
	"gorm-crud-project/repository"
	"gorm-crud-project/routes"
	"gorm-crud-project/service"

	"github.com/gin-gonic/gin"
)

func main() {

	db := config.ConnectDatabase()

	userRepo := &repository.UserRepository{
		DB: db,
	}

	userService := &service.UserService{
		Repo: userRepo,
	}

	userHandler := &handlers.UserHandler{
		Service: userService,
	}

	router := gin.Default()

	routes.RegisterUserRoutes(router, userHandler)
	go kafka.StartUserCreatedConsumer()
	router.Run(":8080")

	// var user models.User
	// result := db.First(&user, 5)
	// fmt.Println(user)
	// fmt.Printf("%+v\n", user)

	// var users []models.User
	// result := db.Find(&users)
	// fmt.Printf("%+v\n", users)

	// var user models.User
	// result := db.Where("email=?", "lakshmi.reddy@gmail.com").First(&user)
	// fmt.Printf("%+v \n", user)

	// var users []models.User
	// result := db.Where("age>?", 22).Find(&users)
	// fmt.Printf("%+v", users)
	// result := db.Where("age > ? and Name = ?", 21, "YashVerma").First(&users)
	// fmt.Printf("%+v", users)

	// result := db.Limit(5).Find(&users)
	// fmt.Printf("%+v \n", users)

	// var count int64
	// result := db.Model(&models.User{}).Count(&count)
	// fmt.Println(count)

	// var users []models.User
	// db.Select("name", "email").Find(&users)
	// fmt.Printf("%+v\n", users)

	// var emails []string
	// db.Model(&models.User{}).Pluck("email", &emails)
	// fmt.Printf("%+v \n ", emails)

}
