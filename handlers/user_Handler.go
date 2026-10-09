package handlers

import (
	"gorm-crud-project/models"
	"gorm-crud-project/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service *service.UserService
}

// func (h *UserHandler) GetUsersHandler(c *gin.Context) {
// 	c.JSON(200, gin.H{"message": "get users"})
// }

func (h *UserHandler) GetUserByIdAndEmail(c *gin.Context) {
	id := c.Param("id")
	email := c.Query("email")

	user, err := h.Service.GetUserByIdAndEmail(id, email)
	if err != nil {
		c.JSON(404, gin.H{"message": "user not found "})
		return
	}
	c.JSON(200, user)

}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid json body"})
		return
	}

	createdUser, err := h.Service.CreateUser(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, createdUser)
}
