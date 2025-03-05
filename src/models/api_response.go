package models

type ApiError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ApiResponse struct {
	Data      interface{} `json:"data,omitempty"`
	Errors    []ApiError  `json:"errors,omitempty"`
	IsSuccess bool        `json:"isSuccess"`
}

var Response = ApiResponse{}

func (r *ApiResponse) Success(data interface{}) ApiResponse {
	return ApiResponse{
		Data:      data,
		IsSuccess: true,
	}
}

func (r *ApiResponse) Failed(errors ...ApiError) ApiResponse {
	return ApiResponse{
		Errors:    errors,
		IsSuccess: false,
	}
}

func (r *ApiResponse) BadRequest(err error) ApiResponse {
	return Response.Failed(ApiError{Code: "bad_request", Message: err.Error()})
}

func (r *ApiResponse) NotFound(err error) ApiResponse {
	return Response.Failed(ApiError{Code: "not_found", Message: err.Error()})
}

func (r *ApiResponse) Unauthorized(err error) ApiResponse {
	return Response.Failed(ApiError{Code: "unauthorized", Message: err.Error()})
}
