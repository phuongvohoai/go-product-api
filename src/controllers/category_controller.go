package controllers

import (
	"phuong/go-product-api/models"
	"phuong/go-product-api/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryService *services.CategoryService
}

type CategoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type CategoryResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewCategoryController(categoryService *services.CategoryService) *CategoryController {
	return &CategoryController{categoryService}
}

// CreateCategory godoc
//
//	@Summary		Create a new category
//	@Description	Create a new category with the provided information
//	@Tags			categories
//	@Accept			json
//	@Produce		json
//	@Param			category	body		CategoryRequest	true	"Category information"
//	@Success		200			{object}	models.ApiResponse{data=CategoryResponse}
//	@Failure		400			{object}	models.ApiResponse
//	@Router			/api/v1/categories [post]
func (ctr *CategoryController) CreateCategory(c *gin.Context) {
	var newCategory CategoryRequest

	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.JSON(400, models.Response.BadRequest(err))
		return
	}

	category, err := ctr.categoryService.CreateCategory(c, &models.Category{
		Name:        newCategory.Name,
		Description: newCategory.Description,
	})

	if err != nil {
		c.JSON(400, models.Response.BadRequest(err))
		return
	}

	c.JSON(201, models.Response.Success(toCategoryResponse(&category)))
}

// GetCategory godoc
//
//	@Summary		Get a category by ID
//	@Description	Get category details by category ID
//	@Tags			categories
//	@Produce		json
//	@Param			id	path		int	true	"Category ID"
//	@Success		200	{object}	models.ApiResponse{data=CategoryResponse}
//	@Failure		400	{object}	models.ApiResponse
//	@Failure		404	{object}	models.ApiResponse
//	@Router			/api/v1/categories/{id} [get]
func (ctr *CategoryController) GetCategoryById(c *gin.Context) {
	id := c.Param("id")

	// Convert string to uint
	categoryId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, models.Response.BadRequest(err))
		return
	}

	category, err := ctr.categoryService.GetCategory(c, categoryId)
	if err != nil {
		c.JSON(404, models.Response.NotFound(err))
		return
	}

	c.JSON(200, models.Response.Success(toCategoryResponse(&category)))
}

// GetCategories godoc
//
//	@Summary		List all categories
//	@Description	Get a list of all categories
//	@Tags			categories
//	@Produce		json
//	@Success		200	{object}	models.ApiResponse{data=[]CategoryResponse}
//	@Failure		400	{object}	models.ApiResponse
//	@Router			/api/v1/categories [get]
func (ctr *CategoryController) GetCategories(c *gin.Context) {
	categories, err := ctr.categoryService.GetCategories(c)
	if err != nil {
		c.JSON(400, models.Response.BadRequest(err))
		return
	}

	categoriesResponse := make([]CategoryResponse, 0)
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, toCategoryResponse(&category))
	}

	c.JSON(200, models.Response.Success(categoriesResponse))
}

// UpdateCategory godoc
//
//	@Summary		Update a category
//	@Description	Update category information by category ID
//	@Tags			categories
//	@Accept			json
//	@Produce		json
//	@Param			id			path		int				true	"Category ID"
//	@Param			category	body		CategoryRequest	true	"Updated category information"
//	@Success		200			{object}	models.ApiResponse{data=CategoryResponse}
//	@Failure		400			{object}	models.ApiResponse
//	@Router			/api/v1/categories/{id} [put]
func (ctr *CategoryController) UpdateCategory(c *gin.Context) {
	id := c.Param("id")

	// Convert string to uint
	categoryId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, models.Response.BadRequest(err))
		return
	}

	var categoryReq CategoryRequest
	if err := c.ShouldBindJSON(&categoryReq); err != nil {
		c.JSON(400, models.Response.BadRequest(err))
		return
	}

	updatedCategory, err := ctr.categoryService.UpdateCategory(c, &models.Category{
		ID:          uint(categoryId),
		Name:        categoryReq.Name,
		Description: categoryReq.Description,
	})

	if err != nil {
		c.JSON(400, models.Response.BadRequest(err))
		return
	}

	c.JSON(200, models.Response.Success(toCategoryResponse(&updatedCategory)))
}

// DeleteCategory godoc
//
//	@Summary		Delete a category
//	@Description	Delete a category by ID
//	@Tags			categories
//	@Produce		json
//	@Param			id	path		int	true	"Category ID"
//	@Success		200	{object}	models.ApiResponse{data=bool}
//	@Failure		400	{object}	models.ApiResponse
//	@Router			/api/v1/categories/{id} [delete]
func (ctr *CategoryController) DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	categoryId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, models.Response.BadRequest(err))
		return
	}

	err = ctr.categoryService.DeleteCategory(c, categoryId)
	if err != nil {
		c.JSON(400, models.Response.BadRequest(err))
		return
	}

	c.JSON(200, models.Response.Success(true))
}

func toCategoryResponse(c *models.Category) CategoryResponse {
	return CategoryResponse{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
	}
}
