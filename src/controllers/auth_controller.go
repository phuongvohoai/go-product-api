package controllers

import (
	"phuong/go-product-api/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userService services.UserService
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewAuthController(userService services.UserService) *AuthController {
	return &AuthController{userService}
}

// Login godoc
//
//	@Summary		Login
//	@Tags			auth
//	@Param			body body LoginRequest true "Login request"
//	@Router			/api/v1/auth/login [post]
func (ctr *AuthController) Login(c *gin.Context) {
	var loginRequest LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	user, err := ctr.userService.VerifyUser(c, loginRequest.Username, loginRequest.Password)
	if err != nil {
		c.JSON(404, gin.H{"error": "Invalid username or password"})
		return
	}

	c.JSON(200, user)
}
