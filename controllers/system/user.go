package system

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gjssss/soybean-admin-go/models/system"
	"github.com/gjssss/soybean-admin-go/utils"
)

type UserController struct{}

// 获取所有用户（GET）
func (c *UserController) GetAllUsers(ctx *gin.Context) {
	page := utils.ParsePagination(ctx)
	users, count, err := SystemService.User.GetAllUsers(page)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("获取用户失败: "+err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(utils.NewPagination(users, page, count)))
}

// 获取用户信息（GET）
func (c *UserController) GetUserInfo(ctx *gin.Context) {
	uid, _ := ctx.Get("userID")
	user, err := SystemService.User.GetUserById(uid.(uint))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.NewErrorResponse("获取用户信息失败: "+err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(user))
}

// 创建用户（POST）
func (c *UserController) CreateUser(ctx *gin.Context) {
	var user system.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("请求参数无效: "+err.Error()))
		return
	}

	createdUser, err := SystemService.User.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("创建用户失败: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(createdUser))
}

// 更新用户密码（POST）
func (c *UserController) UpdateUserPassword(ctx *gin.Context) {
	var params struct {
		ID       uint   `json:"id" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("请求参数无效: "+err.Error()))
		return
	}

	err := SystemService.User.UpdateUserPassword(system.User{ID: params.ID, Password: params.Password})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("更新密码失败: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSuccessResponse("密码更新成功"))
}

// 删除用户（POST）
func (c *UserController) DeleteUser(ctx *gin.Context) {
	var user system.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("请求参数无效: "+err.Error()))
		return
	}

	if err := SystemService.User.DeleteUser(user); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("删除用户失败: "+err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse("用户删除成功"))
}

// 批量删除用户（POST）
func (c *UserController) BatchDeleteUser(ctx *gin.Context) {
	var ids []uint
	if err := ctx.ShouldBindJSON(&ids); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("请求参数无效: "+err.Error()))
		return
	}

	if err := SystemService.User.BatchDeleteUser(ids); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("批量删除用户失败: "+err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse("批量删除成功"))
}

// 更新用户角色（POST）
func (c *UserController) UpdateUserRoles(ctx *gin.Context) {
	var params struct {
		UserID  uint   `json:"id" binding:"required"`
		RoleIDs []uint `json:"roleIds" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("请求参数无效: "+err.Error()))
		return
	}

	if err := SystemService.User.UpdateUserRoles(params.UserID, params.RoleIDs); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("更新用户角色失败: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSuccessResponse("更新用户角色成功"))
}

// 登录（POST）
func (c *UserController) Login(ctx *gin.Context) {
	var user system.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("请求参数无效: "+err.Error()))
		return
	}
	token, err := SystemService.User.Login(user.UserName, user.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.NewErrorResponse("登录失败: "+err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(token))
}

// 刷新Token（POST）
func (c *UserController) RefreshToken(ctx *gin.Context) {
	var data struct {
		RefreshToken string `json:"refreshToken" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("请求参数无效: "+err.Error()))
		return
	}
	token, _ := ctx.Get("accessToken")
	t, err := SystemService.User.Refresh(token.(string), data.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.NewErrorResponse("刷新Token失败: "+err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(t))
}

// 检查用户名是否存在（GET）
func (c *UserController) CheckUserNameExists(ctx *gin.Context) {
	userName := ctx.Query("userName")
	if userName == "" {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("用户名不能为空"))
		return
	}
	exists := SystemService.User.CheckUserNameExists(userName)
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(gin.H{"exists": exists}))
}

// 获取用户角色列表（GET）
func (c *UserController) GetUserRoles(ctx *gin.Context) {
	userID, err := strconv.ParseUint(ctx.Query("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("用户ID不合法"))
		return
	}

	roles, err := SystemService.User.GetUserRoles(uint(userID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("获取用户角色失败: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(roles))
}
