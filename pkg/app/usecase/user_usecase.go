package usecase

import (
	"html"
	"strings"

	"github.com/satyamvatstyagi/UserManagementService/pkg/app/domain"
	"github.com/satyamvatstyagi/UserManagementService/pkg/app/models"
	"github.com/satyamvatstyagi/UserManagementService/pkg/common/jwt"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepository models.UserRepository
}

func NewUserUsecase(userRepository models.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (u *userUsecase) RegisterUser(registerUserRequest *domain.RegisterUserRequest) (*domain.RegisterUserResponse, error) {
	// Encrypt the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerUserRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Remove the space from the username
	registerUserRequest.UserName = html.EscapeString(strings.TrimSpace(registerUserRequest.UserName))

	// Call the repository
	userID, err := u.userRepository.RegisterUser(registerUserRequest.UserName, string(hashedPassword))
	if err != nil {
		return nil, err
	}

	return &domain.RegisterUserResponse{
		UserID: userID,
	}, nil
}

func (u *userUsecase) LoginUser(loginUserRequest *domain.LoginUserRequest) (*domain.LoginUserResponse, error) {
	// Remove the space from the username
	loginUserRequest.UserName = html.EscapeString(strings.TrimSpace(loginUserRequest.UserName))

	// Call the repository
	user, err := u.userRepository.GetUserByUserName(loginUserRequest.UserName)
	if err != nil {
		return nil, err
	}

	// Compare the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUserRequest.Password)); err != nil {
		return nil, err
	}

	// Generate the JWT token
	token, err := jwt.GenerateToken(user.UUID.String(), user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &domain.LoginUserResponse{
		Token: token,
	}, nil
}

func (u *userUsecase) GetUserByUserName(getUserByUserNameRequest *domain.GetUserByUserNameRequest) (*domain.GetUserByUserNameResponse, error) {
	// Remove the space from the username
	getUserByUserNameRequest.UserName = html.EscapeString(strings.TrimSpace(getUserByUserNameRequest.UserName))

	// Call the repository
	user, err := u.userRepository.GetUserByUserName(getUserByUserNameRequest.UserName)
	if err != nil {
		return nil, err
	}

	return &domain.GetUserByUserNameResponse{
		ID:        user.UUID.String(),
		UserName:  user.UserName,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}
