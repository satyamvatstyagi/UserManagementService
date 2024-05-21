package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/satyamvatstyagi/UserManagementService/pkg/app/domain"
	"github.com/satyamvatstyagi/UserManagementService/pkg/common/cerr"
)

type UserController struct {
	UserUsecase domain.UserUsecase
}

// HealthCheck      godoc
//
//	@Summary		Health Check
//	@Description	Health Check will return a message indicating that the user management service is up and running
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	domain.Response	"Service is up and running"
//
//	@Router			/user/health [get]
//	@Tags			user management service
func (c *UserController) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, domain.Response{Message: "User Management Service v2.0 is up and running", Success: true})
}

// RegisterUser godoc
//
//	@Summary		Register a new user
//	@Description	Register a new user
//	@Accept			json
//	@Produce		json
//	@Param			user	body		domain.RegisterUserRequest	true	"User Details"
//	@Success		200		{object}	domain.RegisterUserResp		"User Registered Successfully"
//	@Failure		400		{object}	domain.ErrorResponse		"Invalid Request"
//	@Failure		401		{object}	domain.ErrorResponse		"Unauthorized"
//	@Failure		500		{object}	domain.ErrorResponse		"Internal Server Error"
//	@Router			/user/register [post]
//	@Tags			user management service
func (c *UserController) RegisterUser(ctx *gin.Context) {
	var req domain.RegisterUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Println("[UserController][RegisterUser] Error in ShouldBindJSON: ", err)
		ctx.JSON(http.StatusBadRequest, domain.Response{Message: "Invalid Request", Success: false})
		return
	}

	// Call the usecase
	res, err := c.UserUsecase.RegisterUser(ctx.Request.Context(), &req)
	if err != nil {
		log.Println("[UserController][RegisterUser] Error in RegisterUser: ", err)
		ctx.JSON(http.StatusBadRequest, domain.Response{Message: cerr.GetErrorMessage(err), Success: false})
		return
	}

	ctx.JSON(http.StatusOK, domain.Response{Message: "User Registered Successfully", Success: true, Data: *res})
}

// LoginUser godoc
//
//	@Summary		Login a user
//	@Description	Login a user
//	@Accept			json
//	@Produce		json
//	@Param			user	body		domain.LoginUserRequest	true	"User Details"
//	@Success		200		{object}	domain.LoginSuccessResp	"User Logged In Successfully"
//	@Failure		400		{object}	domain.ErrorResponse	"Invalid Request"
//	@Failure		401		{object}	domain.ErrorResponse	"Unauthorized"
//	@Failure		500		{object}	domain.ErrorResponse	"Internal Server Error"
//	@Router			/user/login [post]
//	@Tags			user management service
func (c *UserController) LoginUser(ctx *gin.Context) {
	var req domain.LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Println("[UserController][LoginUser] Error in ShouldBindJSON: ", err)
		ctx.JSON(http.StatusBadRequest, domain.Response{Message: "Invalid Request", Success: false})
		return
	}

	// Call the usecase
	res, err := c.UserUsecase.LoginUser(ctx.Request.Context(), &req)
	if err != nil {
		log.Println("[UserController][LoginUser] Error in LoginUser: ", err)
		ctx.JSON(http.StatusBadRequest, domain.Response{Message: cerr.GetErrorMessage(err), Success: false})
		return
	}

	ctx.JSON(http.StatusOK, domain.Response{Message: "User Logged In Successfully", Success: true, Data: *res})
}

// GetUserByUserName godoc
//
//	@Summary		Get user by username
//	@Description	Get user by username
//	@Accept			json
//	@Produce		json
//	@Param			username	path		string							true	"User Name"
//	@Success		200			{object}	domain.GetUserByUserNameResp	"User Fetched Successfully"
//	@Failure		400			{object}	domain.ErrorResponse			"Invalid Request"
//	@Failure		401			{object}	domain.ErrorResponse			"Unauthorized"
//	@Failure		500			{object}	domain.ErrorResponse			"Internal Server Error"
//	@Router			/user/{username} [get]
//	@Tags			user management service
func (c *UserController) GetUserByUserName(ctx *gin.Context) {
	var req domain.GetUserByUserNameRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		log.Println("[UserController][GetUserByUserName] Error in ShouldBindUri: ", err)
		ctx.JSON(http.StatusBadRequest, domain.Response{Message: "Invalid Request", Success: false})
		return
	}

	// Call the usecase
	res, err := c.UserUsecase.GetUserByUserName(ctx.Request.Context(), &req)
	if err != nil {
		log.Println("[UserController][GetUserByUserName] Error in GetUserByUserName: ", err)
		ctx.JSON(http.StatusBadRequest, domain.Response{Message: cerr.GetErrorMessage(err), Success: false})
		return
	}

	ctx.JSON(http.StatusOK, domain.Response{Message: "User Fetched Successfully", Success: true, Data: *res})
}

// GetFibonacci godoc
//
//	@Summary		Calculate Fibonacci
//	@Description	Calculate Fibonacci
//	@Accept			json
//	@Produce		json
//	@Param			n	query		int						true	"Fibonacci Number"
//	@Success		200	{object}	domain.FibonacciResp	"Fibonacci Calculated Successfully"
//	@Failure		400	{object}	domain.ErrorResponse	"Invalid Request"
//	@Failure		401	{object}	domain.ErrorResponse	"Unauthorized"
//	@Failure		500	{object}	domain.ErrorResponse	"Internal Server Error"
//	@Router			/user/fibonacci [get]
//	@Tags			user management service
func (c *UserController) GetFibonacci(ctx *gin.Context) {
	var req domain.FibonacciRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		log.Println("[UserController][GetFibonacci] Error in ShouldBindUri: ", err)
		ctx.JSON(http.StatusBadRequest, domain.Response{Message: "Invalid Request", Success: false})
		return
	}

	// Call the usecase
	res, err := c.UserUsecase.Fibonacci(ctx.Request.Context(), req.Number)
	if err != nil {
		log.Println("[UserController][GetFibonacci] Error in Fibonacci: ", err)
		ctx.JSON(http.StatusBadRequest, domain.Response{Message: cerr.GetErrorMessage(err), Success: false})
		return
	}

	ctx.JSON(http.StatusOK, domain.Response{Message: "Fibonacci Calculated Successfully", Success: true, Data: res})
}
