package controllers

import (
	"phuong/go-product-api/models"
	"phuong/go-product-api/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

type UserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService}
}

func (ctr *UserController) CreateUser(c *gin.Context) {
	var newUser UserRequest

	if err := c.ShouldBindJSON(&newUser); err != nil {
		return
	}
	user, err := ctr.userService.CreateUser(c, &models.User{
		Username: newUser.Username,
		Email:    newUser.Email,
	}, newUser.Password)

	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, user)
}

func (ctr *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")

	// Convert string to int
	userId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
	}

	user, err := ctr.userService.GetUser(c, userId)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, user)
}

func (ctr *UserController) GetUsers(c *gin.Context) {
	users, err := ctr.userService.GetUsers(c)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, users)
}

func (ctr *UserController) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	// Convert string to int
	userId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
	}

	var user UserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	updatedUser, err := ctr.userService.UpdateUser(c, &models.User{
		ID:       uint(userId),
		Username: user.Username,
		Email:    user.Email,
	}, user.Password)

	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, updatedUser)
}

func (ctr *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	// Convert string to int
	userId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
	}

	err = ctr.userService.DeleteUser(c, userId)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, gin.H{"message": "User deleted successfully"})
}
