package controllers

import (
	"errors"
	"log"
	"phuong/go-product-api/models"
	"phuong/go-product-api/services"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userService  services.UserService
	emailService services.EmailService
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func NewAuthController(userService services.UserService, emailService services.EmailService) *AuthController {
	return &AuthController{userService, emailService}
}

// Login godoc
//
//	@Summary		Login
//	@Tags			auth
//	@Produce		json
//	@Param			body LoginRequest true "Login request"
//	@Router			/api/v1/auth/login [post]
func (ctr *AuthController) Login(c *gin.Context) {
	var loginRequest LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		_ = c.Error(err)
		return
	}

	user, err := ctr.userService.VerifyUser(c, loginRequest.Username, loginRequest.Password)
	if err != nil {
		_ = c.Error(err)
		return
	}

	tokenCh := make(chan string, 1)
	var wg sync.WaitGroup
	wg.Add(2)

	// Generate token and send email notification concurrently
	go ctr.generateToken(user.ID, user.Username, user.Email, &wg, tokenCh)
	go ctr.sendEmail(user.Email, "New login", "Your account has been logged in from a new device", &wg)

	wg.Wait()

	token := <-tokenCh
	close(tokenCh)

	c.JSON(200, models.Response.Success(&LoginResponse{
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}))
}

// Logout godoc
// @Summary			Logout
// @Tag				Auth
// @Produce			json
//
//	@Router			/api/v1/auth/logout [post]
func (ctr *AuthController) Logout(c *gin.Context) {
	claims, exists := c.Get("Claims")
	if !exists {
		c.JSON(400, models.Response.BadRequest(errors.New("MISSING CLAIMS")))
		return
	}

	// Type assertion to cast claims to service.Claims
	typedClaims, ok := claims.(services.Claims)
	if !ok {
		c.JSON(400, models.Response.BadRequest(errors.New("INVALID CLAIMS")))
		return
	}

	err := services.RevokeToken(typedClaims.StandardClaims.Id)
	if err != nil {
		c.JSON(400, models.Response.BadRequest(errors.New("TOKEN REVOKE FAILED")))
		return
	}

	c.JSON(200, models.Response.Success("Successfully logged out"))
}

func (ctr *AuthController) sendEmail(email string, subject string, body string, wg *sync.WaitGroup) {
	defer wg.Done()
	err := ctr.emailService.SendEmail(email, subject, body)
	if err != nil {
		log.Println("Error sending email: ", err)
	}
}

func (ctr *AuthController) generateToken(id uint, username string, email string, wg *sync.WaitGroup, tokenCh chan<- string) {
	defer wg.Done()
	token, err := services.GenerateToken(id, username, email)
	if err != nil {
		log.Println("Error generating token: ", err)
	}
	log.Println("Generated token: ", token)
	time.Sleep(1 * time.Second) // Simulate token generation delay
	tokenCh <- token
}
