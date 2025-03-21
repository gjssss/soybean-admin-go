package system

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gjssss/soybean-admin-go/models/system"
	"github.com/gjssss/soybean-admin-go/utils"
)

type MenuController struct{}

// @Summary 获取所有菜单
// @Description 获取系统所有菜单列表
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Success 200 {object} utils.Response[[]system.Menu] "成功"
// @Failure 400 {object} utils.Response[string] "错误"
// @Security ApiKeyAuth
// @Router /menus [get]
func (c *MenuController) GetMenus(ctx *gin.Context) {
	menus, err := SystemService.Menu.GetMenus()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("获取菜单失败: "+err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(menus))
}

// @Summary 获取常量菜单
// @Description 获取系统常量菜单列表
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Success 200 {object} utils.Response[[]system.Menu] "成功"
// @Failure 400 {object} utils.Response[string] "错误"
// @Router /menus/constant [get]
func (c *MenuController) GetConstantMenus(ctx *gin.Context) {
	menus, err := SystemService.Menu.GetConstantMenu()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("获取常量菜单失败: "+err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(menus))
}

// @Summary 获取用户菜单
// @Description 获取当前用户的菜单列表
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Success 200 {object} utils.Response[[]system.Menu] "成功"
// @Failure 400 {object} utils.Response[string] "错误"
// @Security ApiKeyAuth
// @Router /menus/user [get]
func (c *MenuController) GetUserMenus(ctx *gin.Context) {
	uid, _ := ctx.Get("userID")
	menus, err := SystemService.Menu.GetMenusByUserId(uid.(uint))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("获取用户菜单失败: "+err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(menus))
}

// @Summary 获取角色菜单
// @Description 获取指定角色的菜单列表
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Param roleId query uint true "角色ID"
// @Success 200 {object} utils.Response[[]system.Menu] "成功"
// @Failure 400 {object} utils.Response[string] "错误"
// @Security ApiKeyAuth
// @Router /menus/role [get]
func (c *MenuController) GetMenusByRoleId(ctx *gin.Context) {
	rid := ctx.Query("roleId")
	if rid == "" {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("角色ID不能为空"))
		return
	}

	roleId, err := strconv.ParseUint(rid, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("角色ID必须为数字"))
		return
	}

	menus, err := SystemService.Menu.GetMenusByRoleId(uint(roleId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("获取角色菜单失败: "+err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(menus))
}

// @Summary 创建菜单
// @Description 创建新菜单
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Param menu body system.Menu true "菜单信息"
// @Success 200 {object} utils.Response[[]system.Menu] "成功"
// @Failure 400 {object} utils.Response[string] "错误"
// @Security ApiKeyAuth
// @Router /menus [post]
func (c *MenuController) CreateMenu(ctx *gin.Context) {
	var menu system.Menu
	if err := ctx.ShouldBindJSON(&menu); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("无效的菜单信息: "+err.Error()))
		return
	}

	if err := SystemService.Menu.CreateMenu(&menu); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("创建菜单失败: "+err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(menu))
}

// @Summary 更新菜单
// @Description 更新菜单信息
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Param menu body system.Menu true "菜单信息"
// @Success 200 {object} utils.Response[[]system.Menu] "成功"
// @Failure 400 {object} utils.Response[string] "错误"
// @Security ApiKeyAuth
// @Router /menus/update [post]
func (c *MenuController) UpdateMenu(ctx *gin.Context) {
	var menu system.Menu
	if err := ctx.ShouldBindJSON(&menu); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("无效的菜单信息: "+err.Error()))
		return
	}

	if menu.ID == 0 {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("菜单ID不能为空"))
		return
	}

	if err := SystemService.Menu.UpdateMenu(&menu); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("更新菜单失败: "+err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(menu))
}

// @Summary 删除菜单
// @Description 删除指定菜单
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Param id body object{id=uint} true "菜单ID"
// @Success 200 {object} utils.Response[string] "成功"
// @Failure 400 {object} utils.Response[string] "错误"
// @Security ApiKeyAuth
// @Router /menus/delete [post]
func (c *MenuController) DeleteMenu(ctx *gin.Context) {
	var params struct {
		ID uint `json:"id" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("请求参数无效: "+err.Error()))
		return
	}

	if err := SystemService.Menu.DeleteMenu(params.ID); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("删除菜单失败: "+err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse("删除成功"))
}

// @Summary 批量删除菜单
// @Description 批量删除多个菜单
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Param ids body object{ids=[]uint} true "菜单ID列表"
// @Success 200 {object} utils.Response[string] "成功"
// @Failure 400 {object} utils.Response[string] "错误"
// @Security ApiKeyAuth
// @Router /menus/batchDelete [post]
func (c *MenuController) BatchDeleteMenu(ctx *gin.Context) {
	var params struct {
		IDs []uint `json:"ids" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("请求参数无效: "+err.Error()))
		return
	}

	if len(params.IDs) == 0 {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("菜单ID列表不能为空"))
		return
	}

	if err := SystemService.Menu.BatchDeleteMenu(params.IDs); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("批量删除菜单失败: "+err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse("批量删除成功"))
}
