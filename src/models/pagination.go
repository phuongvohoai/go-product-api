package models

type Pagination struct {
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	SortBy   string `json:"sort_by"`
	SortDir  string `json:"sort_dir"`
}

type PaginationResponse struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Total    int `json:"total"`
}

type PaginatedListResponse struct {
	Items      interface{}        `json:"items"`
	Pagination PaginationResponse `json:"pagination"`
}

func NewPaginatedListResponse(items interface{}, page, pageSize, total int) PaginatedListResponse {
	return PaginatedListResponse{
		Items: items,
		Pagination: PaginationResponse{
			Page:     page,
			PageSize: pageSize,
			Total:    total,
		},
	}
}
