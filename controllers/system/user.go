package system

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gjssss/soybean-admin-go/models/system"
	"github.com/gjssss/soybean-admin-go/utils"
)

type UserController struct{}

// @Summary 获取用户列表
// @Description 获取分页的用户列表
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param pageSize query int false "每页条数" default(10)
// @Success 200 {object} utils.Response[utils.Pagination[[]system.User]] "成功"
// @Failure 400 {object} utils.Response[string] "错误"
// @Security ApiKeyAuth
// @Router /users [get]
func (c *UserController) GetAllUsers(ctx *gin.Context) {
	page := utils.ParsePagination(ctx)
	users, count, err := SystemService.User.GetAllUsers(page)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("获取用户失败: "+err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(utils.NewPagination(users, page, count)))
}

// @Summary 获取用户信息
// @Description 获取当前登录用户的详细信息
// @Tags 认证
// @Accept json
// @Produce json
// @Success 200 {object} utils.Response[system.User] "成功"
// @Failure 401 {object} utils.Response[string] "认证失败"
// @Security ApiKeyAuth
// @Router /auth/getUserInfo [get]
func (c *UserController) GetUserInfo(ctx *gin.Context) {
	uid, _ := ctx.Get("userID")
	user, err := SystemService.User.GetUserById(uid.(uint))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.NewErrorResponse("获取用户信息失败: "+err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(user))
}

// @Summary 创建用户
// @Description 创建新用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param user body system.User true "用户信息"
// @Success 200 {object} utils.Response[system.User] "成功"
// @Failure 400 {object} utils.Response[string] "错误"
// @Security ApiKeyAuth
// @Router /users [post]
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

// @Summary 更新用户密码
// @Description 更新用户密码
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param passwordInfo body object{id=uint,password=string} true "密码信息"
// @Success 200 {object} utils.Response[string] "成功"
// @Failure 400 {object} utils.Response[string] "错误"
// @Security ApiKeyAuth
// @Router /users/password [post]
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

// @Summary 删除用户
// @Description 删除指定用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param user body system.User true "用户信息"
// @Success 200 {object} utils.Response[string] "成功"
// @Failure 400 {object} utils.Response[string] "错误"
// @Security ApiKeyAuth
// @Router /users/delete [post]
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

// @Summary 批量删除用户
// @Description 批量删除多个用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param ids body []uint true "用户ID列表"
// @Success 200 {object} utils.Response[string] "成功"
// @Failure 400 {object} utils.Response[string] "错误"
// @Security ApiKeyAuth
// @Router /users/batchDelete [post]
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

// @Summary 更新用户角色
// @Description 更新用户的角色列表
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param roleInfo body object{id=uint,roleIds=[]uint} true "用户角色信息"
// @Success 200 {object} utils.Response[string] "成功"
// @Failure 400 {object} utils.Response[string] "错误"
// @Security ApiKeyAuth
// @Router /users/roles [post]
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

// @Summary 用户登录
// @Description 用户登录接口，获取token
// @Tags 认证
// @Accept json
// @Produce json
// @Param user body object{userName=string,password=string} true "登录信息"
// @Success 200 {object} utils.Response[system.Token] "成功"
// @Failure 401 {object} utils.Response[string] "认证失败"
// @Router /auth/login [post]
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

// @Summary 刷新Token
// @Description 使用刷新令牌获取新的访问令牌
// @Tags 认证
// @Accept json
// @Produce json
// @Param refreshInfo body object{refreshToken=string} true "刷新令牌"
// @Success 200 {object} utils.Response[system.Token] "成功"
// @Failure 401 {object} utils.Response[string] "认证失败"
// @Security ApiKeyAuth
// @Router /auth/refreshToken [post]
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

// @Summary 检查用户名是否存在
// @Description 检查用户名是否已被使用
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param userName query string true "用户名"
// @Success 200 {object} utils.Response[utils.ExistsResult] "成功"
// @Failure 400 {object} utils.Response[string] "错误"
// @Security ApiKeyAuth
// @Router /users/checkUsername [get]
func (c *UserController) CheckUserNameExists(ctx *gin.Context) {
	userName := ctx.Query("userName")
	if userName == "" {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("用户名不能为空"))
		return
	}
	exists := SystemService.User.CheckUserNameExists(userName)
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(utils.ExistsResult{Exists: exists}))
}

// @Summary 获取用户角色
// @Description 获取指定用户的角色列表
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param id query uint true "用户ID"
// @Success 200 {object} utils.Response[[]system.Role] "成功"
// @Failure 400 {object} utils.Response[string] "错误"
// @Security ApiKeyAuth
// @Router /users/roles [get]
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
