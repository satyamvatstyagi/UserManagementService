package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/satyamvatstyagi/UserManagementService/pkg/app/domain"
	"github.com/satyamvatstyagi/UserManagementService/pkg/common/cerr"
	"github.com/satyamvatstyagi/UserManagementService/pkg/common/logger"
)

type UserController struct {
	UserUsecase domain.UserUsecase
	Logger      *logger.MtnLogger
}

func (c *UserController) RegisterUser(ctx *gin.Context) {
	var req domain.RegisterUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
		return
	}

	// Call the usecase
	res, err := c.UserUsecase.RegisterUser(&req)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": cerr.GetErrorMessage(err)})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User Registered Successfully", "data": gin.H{"user_id": res.UserID}})
}

func (c *UserController) LoginUser(ctx *gin.Context) {
	var req domain.LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
		return
	}

	// Call the usecase
	res, err := c.UserUsecase.LoginUser(&req)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": cerr.GetErrorMessage(err)})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User Logged In Successfully", "data": gin.H{"token": res.Token}})
}

func (c *UserController) GetUserByUserName(ctx *gin.Context) {
	var req domain.GetUserByUserNameRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
		return
	}

	// Call the usecase
	res, err := c.UserUsecase.GetUserByUserName(&req)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": cerr.GetErrorMessage(err)})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User Fetched Successfully", "data": res})
}

func (c *UserController) GetOrderByOrderUserName(ctx *gin.Context) {
	var req domain.GetOrderByOrderUserNameRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
		return
	}

	// Call the usecase
	res, err := c.UserUsecase.GetOrderByOrderUserName(&req)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": cerr.GetErrorMessage(err)})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Order Fetched Successfully", "data": res})
}
