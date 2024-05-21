package domain

// This file is only for swagger documentation purposes. THis contains all the models used as a request and responses for swagger doc.

// Success response structure for regiser user, intended only for Swagger documentation.
type RegisterUserResp struct {
	SuccessResponse
	Data RegisterUserResponse `json:"data"`
}

// Success response structure for login details, intended only for Swagger documentation.
type LoginSuccessResp struct {
	SuccessResponse
	Data LoginUserResponse `json:"data"`
}

// Success response structure for responses with no data, intended only for Swagger documentation.
type SuccessResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success" example:"true" description:"I am testing"`
}

// Success response structure for get user by username, intended only for Swagger documentation.
type GetUserByUserNameResp struct {
	SuccessResponse
	Data GetUserByUserNameResponse `json:"data"`
}

// Success response structure for get order by order username, intended only for Swagger documentation.
type GetOrderByOrderUserNameResp struct {
	SuccessResponse
	Data GetOrderByOrderUserNameResponse `json:"data"`
}

// Failure response structure, intended only for Swagger documentation.
type ErrorResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success" example:"false"`
}

// Fibonacci response structure, intended only for Swagger documentation.
type FibonacciResp struct {
	SuccessResponse
	Data int `json:"data"`
}
