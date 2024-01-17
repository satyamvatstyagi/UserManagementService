package domain

import "context"

type UserUsecase interface {
	RegisterUser(registerUserRequest *RegisterUserRequest) (registerUserResponse *RegisterUserResponse, err error)
	LoginUser(loginUserRequest *LoginUserRequest) (loginUserResponse *LoginUserResponse, err error)
	GetUserByUserName(ctx context.Context, getUserByUserNameRequest *GetUserByUserNameRequest) (getUserByUserNameResponse *GetUserByUserNameResponse, err error)
	SendRequestToServer(url string, requestJson []byte) (response []byte, err error)
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
