package routes

import (
	"gorm-crud-project/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, userHandler *handlers.UserHandler) {
	// router.GET("/users", userHandler.GetUserByIdAndEmail)
	router.GET("/users/:id", userHandler.GetUserByIdAndEmail)
	router.POST("users", userHandler.CreateUser)
}
