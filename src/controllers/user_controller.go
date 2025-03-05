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

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService}
}

// CreateUser godoc
//
//	@Summary		Create a new user
//	@Description	Create a new user with the provided information
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		UserRequest	true	"User information"
//	@Success		200		{object}	models.ApiResponse{data=UserResponse}
//	@Failure		400		{object}	models.ApiResponse
//	@Router			/api/v1/users [post]
func (ctr *UserController) CreateUser(c *gin.Context) {
	var newUser UserRequest

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, models.Response.BadRequest(err))
		return
	}
	user, err := ctr.userService.CreateUser(c, &models.User{
		Username: newUser.Username,
		Email:    newUser.Email,
	}, newUser.Password)

	if err != nil {
		c.JSON(400, models.Response.BadRequest(err))
		return
	}

	c.JSON(200, models.Response.Success(toResponse(&user)))
}

// GetUser godoc
//
//	@Summary		Get a user by ID
//	@Description	Get user details by user ID
//	@Tags			users
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	models.ApiResponse{data=UserResponse}
//	@Failure		400	{object}	models.ApiResponse
//	@Failure		404	{object}	models.ApiResponse
//	@Router			/api/v1/users/{id} [get]
func (ctr *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")

	// Convert string to int
	userId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, models.Response.BadRequest(err))
		return
	}

	user, err := ctr.userService.GetUser(c, userId)
	if err != nil {
		c.JSON(404, models.Response.NotFound(err))
		return
	}

	c.JSON(200, models.Response.Success(toResponse(&user)))
}

// GetUsers godoc
//
//	@Summary		List all users
//	@Description	Get a list of all users
//	@Tags			users
//	@Produce		json
//	@Success		200	{object}	models.ApiResponse{data=[]UserResponse}
//	@Failure		400	{object}	models.ApiResponse
//	@Router			/api/v1/users [get]
func (ctr *UserController) GetUsers(c *gin.Context) {
	users, err := ctr.userService.GetUsers(c)
	if err != nil {
		c.JSON(400, models.Response.BadRequest(err))
		return
	}

	usersResponse := make([]UserResponse, 0)
	for _, user := range users {
		usersResponse = append(usersResponse, toResponse(&user))
	}

	c.JSON(200, models.Response.Success(usersResponse))
}

// UpdateUser godoc
//
//	@Summary		Update a user
//	@Description	Update user information by user ID
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int			true	"User ID"
//	@Param			user	body		UserRequest	true	"Updated user information"
//	@Success		200		{object}	models.ApiResponse{data=UserResponse}
//	@Failure		400	{object}	models.ApiResponse
//	@Router			/api/v1/users/{id} [put]
func (ctr *UserController) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	// Convert string to int
	userId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, models.Response.BadRequest(err))
		return
	}

	var user UserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, models.Response.BadRequest(err))
		return
	}

	updatedUser, err := ctr.userService.UpdateUser(c, &models.User{
		ID:       uint(userId),
		Username: user.Username,
		Email:    user.Email,
	}, user.Password)

	if err != nil {
		c.JSON(400, models.Response.BadRequest(err))
		return
	}

	c.JSON(200, models.Response.Success(toResponse(&updatedUser)))
}

// DeleteUser godoc
//
//	@Summary		Delete a user
//	@Description	Delete a user by ID
//	@Tags			users
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	models.ApiResponse{data=bool}
//	@Failure		400	{object}	models.ApiResponse
//	@Router			/api/v1/users/{id} [delete]
func (ctr *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	// Convert string to int
	userId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, models.Response.BadRequest(err))
		return
	}

	err = ctr.userService.DeleteUser(c, userId)
	if err != nil {
		c.JSON(400, models.Response.BadRequest(err))
		return
	}

	c.JSON(200, models.Response.Success(true))
}

func toResponse(u *models.User) UserResponse {
	return UserResponse{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
	}
}
