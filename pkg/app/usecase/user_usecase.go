package usecase

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/satyamvatstyagi/UserManagementService/pkg/app/domain"
	"github.com/satyamvatstyagi/UserManagementService/pkg/app/models"
	"github.com/satyamvatstyagi/UserManagementService/pkg/common/apirequest"
	"github.com/satyamvatstyagi/UserManagementService/pkg/common/jwt"
	"github.com/satyamvatstyagi/UserManagementService/pkg/common/restclient"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepository models.UserRepository
	httpClient     restclient.HTTPClient
}

func NewUserUsecase(userRepository models.UserRepository, hc restclient.HTTPClient) domain.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		httpClient:     hc,
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

// Function to send request to http client server
func (u *userUsecase) SendRequestToServer(url string, requestJson []byte) (response []byte, err error) {
	// Create the request
	req, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(requestJson))
	if err != nil {
		return nil, err
	}

	// Set the headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Get the basic auth credentials from the env variables
	username := os.Getenv("BASIC_AUTH_USER")
	password := os.Getenv("BASIC_AUTH_PASSWORD")

	// Encode the credentials
	auth := username + ":" + password
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(auth))

	// Set the Authorization header
	req.Header.Set("Authorization", "Basic "+encodedAuth)

	// Call the http client
	res, err := u.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Check if the status code is 200
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("recieved %d code from client", res.StatusCode)
	}

	// Read the response
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(res.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Convert the response to bytes
	response = buf.Bytes()

	return response, nil
}

func (u *userUsecase) GetOrderByOrderUserName(ctx context.Context, getOrderByOrderUserNameRequest *domain.GetOrderByOrderUserNameRequest) (*domain.GetOrderByOrderUserNameResponse, error) {
	// Remove the space from the username
	getOrderByOrderUserNameRequest.UserName = html.EscapeString(strings.TrimSpace(getOrderByOrderUserNameRequest.UserName))

	// Check if the user exists
	user, err := u.userRepository.GetUserByUserName(getOrderByOrderUserNameRequest.UserName)
	if err != nil {
		return nil, err
	}

	// Call send request to server
	// response, err := u.SendRequestToServer("http://localhost:8081/order/"+user.UserName, nil)
	// if err != nil {
	// 	return nil, err
	// }

	//-------------------------------------------------------------------------------------------------------------------//

	resp, err := apirequest.APIRequest(ctx,
		apirequest.RequestParams{URL: "http://localhost:8081/order/q" + user.UserName,
			Method: http.MethodGet,
			Body:   nil,
			Headers: map[string]string{"Content-Type": "application/json",
				"Accept":        "application/json",
				"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte(os.Getenv("BASIC_AUTH_USER")+":"+os.Getenv("BASIC_AUTH_PASSWORD")))},
			SpanName: "GetUserOrder"})

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	response := buf.Bytes()

	//-------------------------------------------------------------------------------------------------------------------//

	// Convert the response to struct
	var orderResponse domain.GetOrderByOrderUserNameResponse
	err = json.Unmarshal(response, &orderResponse)
	if err != nil {
		return nil, err
	}

	return &orderResponse, nil
}
