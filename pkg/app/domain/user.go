package domain

type UserUsecase interface {
	RegisterUser(registerUserRequest *RegisterUserRequest) (registerUserResponse *RegisterUserResponse, err error)
	LoginUser(loginUserRequest *LoginUserRequest) (loginUserResponse *LoginUserResponse, err error)
	GetUserByUserName(getUserByUserNameRequest *GetUserByUserNameRequest) (getUserByUserNameResponse *GetUserByUserNameResponse, err error)
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
