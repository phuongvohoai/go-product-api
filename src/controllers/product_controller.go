package controllers

import (
	"phuong/go-product-api/models"
	"phuong/go-product-api/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService *services.ProductService
}

type ProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
	CategoryID  int     `json:"category_id" binding:"required"`
}

type ProductResponse struct {
	ID          uint             `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Price       float64          `json:"price"`
	CategoryID  uint             `json:"category_id"`
	Category    CategoryResponse `json:"category"`
}

func NewProductController(productService *services.ProductService) *ProductController {
	return &ProductController{productService}
}

// CreateProduct godoc
//
//	@Summary		Create a new product
//	@Description	Create a new product with the provided information
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			product	body		ProductRequest	true	"Product information"
//	@Success		200		{object}	models.ApiResponse{data=ProductResponse}
//	@Failure		400		{object}	models.ApiResponse
//	@Router			/api/v1/products [post]
func (ctr *ProductController) CreateProduct(c *gin.Context) {
	var newProduct ProductRequest

	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(400, models.Response.BadRequest(err))
		return
	}

	product, err := ctr.productService.CreateProduct(c, &models.Product{
		Name:        newProduct.Name,
		Description: newProduct.Description,
		Price:       newProduct.Price,
		CategoryID:  uint(newProduct.CategoryID),
	})

	if err != nil {
		c.JSON(400, models.Response.BadRequest(err))
		return
	}

	c.JSON(201, models.Response.Success(toProductResponse(&product)))
}

// GetProduct godoc
//
//	@Summary		Get a product by ID
//	@Description	Get product details by product ID
//	@Tags			products
//	@Produce		json
//	@Param			id	path		int	true	"Product ID"
//	@Success		200	{object}	models.ApiResponse{data=ProductResponse}
//	@Failure		400	{object}	models.ApiResponse
//	@Failure		404	{object}	models.ApiResponse
//	@Router			/api/v1/products/{id} [get]
func (ctr *ProductController) GetProductById(c *gin.Context) {
	id := c.Param("id")

	productId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, models.Response.BadRequest(err))
		return
	}

	product, err := ctr.productService.GetProduct(c, productId)
	if err != nil {
		c.JSON(404, models.Response.NotFound(err))
		return
	}

	c.JSON(200, models.Response.Success(toProductResponse(&product)))
}

// GetProducts godoc
//
//	@Summary		List all products
//	@Description	Get a list of all products
//	@Tags			products
//	@Produce		json
//	@Success		200	{object}	models.ApiResponse{data=[]ProductResponse}
//	@Failure		400	{object}	models.ApiResponse
//	@Router			/api/v1/products [get]
func (ctr *ProductController) GetProducts(c *gin.Context) {
	// Check for optional category_id query parameter
	categoryIDStr := c.Query("category_id")

	var products []models.Product
	var err error

	if categoryIDStr != "" {
		categoryID, err := strconv.Atoi(categoryIDStr)
		if err != nil {
			c.JSON(400, models.Response.BadRequest(err))
			return
		}
		products, err = ctr.productService.GetProductsByCategory(c, categoryID)
	} else {
		products, err = ctr.productService.GetProducts(c)
	}

	if err != nil {
		c.JSON(400, models.Response.BadRequest(err))
		return
	}

	productsResponse := make([]ProductResponse, 0)
	for _, product := range products {
		productsResponse = append(productsResponse, toProductResponse(&product))
	}

	c.JSON(200, models.Response.Success(productsResponse))
}

// UpdateProduct godoc
//
//	@Summary		Update a product
//	@Description	Update product information by product ID
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"Product ID"
//	@Param			product	body		ProductRequest	true	"Updated product information"
//	@Success		200		{object}	models.ApiResponse{data=ProductResponse}
//	@Failure		400		{object}	models.ApiResponse
//	@Router			/api/v1/products/{id} [put]
func (ctr *ProductController) UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	productId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, models.Response.BadRequest(err))
		return
	}

	var productReq ProductRequest
	if err := c.ShouldBindJSON(&productReq); err != nil {
		c.JSON(400, models.Response.BadRequest(err))
		return
	}

	updatedProduct, err := ctr.productService.UpdateProduct(c, &models.Product{
		ID:          uint(productId),
		Name:        productReq.Name,
		Description: productReq.Description,
		Price:       productReq.Price,
		CategoryID:  uint(productReq.CategoryID),
	})

	if err != nil {
		c.JSON(400, models.Response.BadRequest(err))
		return
	}

	c.JSON(200, models.Response.Success(toProductResponse(&updatedProduct)))
}

// DeleteProduct godoc
//
//	@Summary		Delete a product
//	@Description	Delete a product by ID
//	@Tags			products
//	@Produce		json
//	@Param			id	path		int	true	"Product ID"
//	@Success		200	{object}	models.ApiResponse{data=bool}
//	@Failure		400	{object}	models.ApiResponse
//	@Router			/api/v1/products/{id} [delete]
func (ctr *ProductController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	productId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, models.Response.BadRequest(err))
		return
	}

	err = ctr.productService.DeleteProduct(c, productId)
	if err != nil {
		c.JSON(400, models.Response.BadRequest(err))
		return
	}

	c.JSON(200, models.Response.Success(true))
}

func toProductResponse(p *models.Product) ProductResponse {
	return ProductResponse{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		CategoryID:  p.CategoryID,
		Category: CategoryResponse{
			ID:          p.Category.ID,
			Name:        p.Category.Name,
			Description: p.Category.Description,
		},
	}
}
