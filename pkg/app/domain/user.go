package domain

type UserUsecase interface {
	RegisterUser(registerUserRequest *RegisterUserRequest) (registerUserResponse *RegisterUserResponse, err error)
	LoginUser(loginUserRequest *LoginUserRequest) (loginUserResponse *LoginUserResponse, err error)
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
