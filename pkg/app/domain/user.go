package domain

import "context"

type UserUsecase interface {
	RegisterUser(ctx context.Context, registerUserRequest *RegisterUserRequest) (registerUserResponse *RegisterUserResponse, err error)
	LoginUser(ctx context.Context, loginUserRequest *LoginUserRequest) (loginUserResponse *LoginUserResponse, err error)
	GetUserByUserName(ctx context.Context, getUserByUserNameRequest *GetUserByUserNameRequest) (getUserByUserNameResponse *GetUserByUserNameResponse, err error)
	Fibonacci(ctx context.Context, n int) (int, error)
	SendRequestToServer(ctx context.Context, url string, requestJson []byte) (response []byte, err error)
	GetOrderByOrderUserName(ctx context.Context, getOrderByOrderUserNameRequest *GetOrderByOrderUserNameRequest) (getOrderByOrderUserNameResponse *GetOrderByOrderUserNameResponse, err error)
}

type RegisterUserRequest struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterUserResponse struct {
	UserID string `json:"user_id"`
}

type LoginUserRequest struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserResponse struct {
	Token string `json:"token"`
}

type GetUserByUserNameRequest struct {
	UserName string `uri:"username" binding:"required"`
}

type GetUserByUserNameResponse struct {
	ID        string `json:"id"`
	UserName  string `json:"user_name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type GetOrderByOrderUserNameRequest struct {
	UserName string `uri:"username" binding:"required"`
}

type GetOrderByOrderUserNameResponse struct {
	OrderID     string `json:"order_id"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
}
type FibonacciRequest struct {
	Number int `uri:"number" binding:"required,numeric,min=0"`
}

type Response struct {
	// Message is a string message returned in the response.
	Message string `json:"message"`
	// Success is a boolean value indicating whether the request was successful or not.
	Success bool `json:"success"`
	// ErrorCode is an integer value indicating the error code.
	ErrorCode int `json:"error_code,omitempty"`
	// Data is the data returned in the response.
	Data interface{} `json:"data,omitempty"`
}

type TokenValidationRequest struct {
	Token string `json:"token" binding:"required"`
}
