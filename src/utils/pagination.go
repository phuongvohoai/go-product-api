package utils

import (
	"phuong/go-product-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaginationOptions struct {
	DefaultPage     int
	DefaultPageSize int
	MaxPageSize     int
	ValidSortFields map[string]bool
	DefaultSortBy   string
	DefaultSortDir  string
}

func CreateDefaultPaginationOptions() PaginationOptions {
	return PaginationOptions{
		DefaultPage:     1,
		DefaultPageSize: 10,
		MaxPageSize:     100,
		ValidSortFields: map[string]bool{
			"id":         true,
			"created_at": true,
		},
		DefaultSortBy:  "id",
		DefaultSortDir: "asc",
	}
}

func ParsePaginationQuery(c *gin.Context, options PaginationOptions) *models.Pagination {
	pageStr := c.DefaultQuery("page", strconv.Itoa(options.DefaultPage))
	pageSizeStr := c.DefaultQuery("page_size", strconv.Itoa(options.DefaultPageSize))
	sortBy := c.DefaultQuery("sort_by", options.DefaultSortBy)
	sortDir := c.DefaultQuery("sort_dir", options.DefaultSortDir)

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = options.DefaultPage
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 || pageSize > options.MaxPageSize {
		pageSize = options.DefaultPageSize
	}

	if !options.ValidSortFields[sortBy] {
		sortBy = options.DefaultSortBy
	}

	if sortDir != "asc" && sortDir != "desc" {
		sortDir = options.DefaultSortDir
	}

	return &models.Pagination{
		Page:     page,
		PageSize: pageSize,
		SortBy:   sortBy,
		SortDir:  sortDir,
	}
}

func CreatePaginationResponse(page, pageSize, total int) map[string]interface{} {
	return map[string]interface{}{
		"page":      page,
		"page_size": pageSize,
		"total":     total,
	}
}
