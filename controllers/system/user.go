package system

import (
	"net/http"

	"github.com/gjssss/soybean-admin-go/models/system"
	"github.com/gjssss/soybean-admin-go/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	page := utils.ParsePagination(ctx)
	users, count, err := SystemService.User.GetAllUsers(page)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(utils.NewPagination(users, page, count)))
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user system.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	createdUser, err := SystemService.User.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(createdUser))
}

func (c *UserController) UpdateUserPassword(ctx *gin.Context) {
	var user system.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	updatedUser, err := SystemService.User.UpdateUserPassword(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(updatedUser))
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	var user system.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	if err := SystemService.User.DeleteUser(user); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(""))
}

func (c *UserController) Login(ctx *gin.Context) {
	var user = system.User{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}
	token, err := SystemService.User.Login(user.UserName, user.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.NewErrorResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(token))
}

func (c *UserController) RefreshToken(ctx *gin.Context) {
	var data = struct {
		RefreshToken string `json:"refreshToken"`
	}{}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}
	token, _ := ctx.Get("accessToken")
	t, err := SystemService.User.Refresh(token.(string), data.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.NewErrorResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(t))
}

func (c *UserController) GetUserInfo(ctx *gin.Context) {
	uid, _ := ctx.Get("userID")
	user, err := SystemService.User.GetUserById(uid.(uint))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.NewErrorResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(user))
}
