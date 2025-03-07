package system

import (
	"net/http"

	"github.com/gjssss/soybean-admin-go/models"
	"github.com/gjssss/soybean-admin-go/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := UserService.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(users))
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	createdUser, err := UserService.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(createdUser))
}

func (c *UserController) UpdateUserPassword(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	updatedUser, err := UserService.UpdateUserPassword(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(updatedUser))
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	if err := UserService.DeleteUser(user); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(""))
}

func (c *UserController) Login(ctx *gin.Context) {
	var user = models.User{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}
	token, err := UserService.Login(user.UserName, user.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.NewErrorResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(token))
}
