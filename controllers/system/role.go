package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gjssss/soybean-admin-go/models/system"
	"github.com/gjssss/soybean-admin-go/utils"
)

type RoleController struct{}

// @Summary 获取所有角色
// @Description 获取所有角色列表（不分页）
// @Tags 角色管理
// @Accept json
// @Produce json
// @Success 200 {object} utils.Response[[]system.Role] "成功"
// @Failure 400 {object} utils.Response[string] "错误"
// @Security ApiKeyAuth
// @Router /roles/all [get]
func (c *RoleController) GetAllRoles(ctx *gin.Context) {
	roles, err := SystemService.Role.GetAllRoles()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("获取所有角色失败: "+err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(roles))
}

// @Summary 获取分页角色列表
// @Description 获取分页的角色列表
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param current query int false "页码" default(1)
// @Param size query int false "每页条数" default(10)
// @Success 200 {object} utils.Response[utils.Pagination[system.Role]] "成功"
// @Failure 400 {object} utils.Response[string] "错误"
// @Security ApiKeyAuth
// @Router /roles [get]
func (c *RoleController) GetRoles(ctx *gin.Context) {
	page := utils.ParsePagination(ctx)
	roles, count, err := SystemService.Role.GetRoles(page)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("获取角色失败: "+err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(utils.NewPagination(roles, page, count)))
}

// @Summary 创建角色
// @Description 创建新角色
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param role body system.Role true "角色信息"
// @Success 200 {object} utils.Response[system.Role] "成功"
// @Failure 400 {object} utils.Response[string] "错误"
// @Security ApiKeyAuth
// @Router /roles [post]
func (c *RoleController) CreateRole(ctx *gin.Context) {
	var role system.Role
	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("请求参数无效: "+err.Error()))
		return
	}

	if err := SystemService.Role.CreateRole(&role); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("创建角色失败: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(role))
}

// @Summary 更新角色
// @Description 更新角色信息
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param role body system.Role true "角色信息"
// @Success 200 {object} utils.Response[string] "成功"
// @Failure 400 {object} utils.Response[string] "错误"
// @Security ApiKeyAuth
// @Router /roles/update [post]
func (c *RoleController) UpdateRole(ctx *gin.Context) {
	var role system.Role
	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("请求参数无效: "+err.Error()))
		return
	}

	if role.ID == 0 {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("角色ID不能为空"))
		return
	}

	if err := SystemService.Role.UpdateRole(&role); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("更新角色失败: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSuccessResponse("更新成功"))
}

// @Summary 删除角色
// @Description 删除指定角色
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param id body object{id=uint} true "角色ID"
// @Success 200 {object} utils.Response[string] "成功"
// @Failure 400 {object} utils.Response[string] "错误"
// @Security ApiKeyAuth
// @Router /roles/delete [post]
func (c *RoleController) DeleteRole(ctx *gin.Context) {
	var params struct {
		ID uint `json:"id" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("请求参数无效: "+err.Error()))
		return
	}

	if err := SystemService.Role.DeleteRole(params.ID); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("删除角色失败: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSuccessResponse("删除成功"))
}

// @Summary 批量删除角色
// @Description 批量删除多个角色
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param ids body object{ids=[]uint} true "角色ID列表"
// @Success 200 {object} utils.Response[string] "成功"
// @Failure 400 {object} utils.Response[string] "错误"
// @Security ApiKeyAuth
// @Router /roles/batchDelete [post]
func (c *RoleController) BatchDeleteRole(ctx *gin.Context) {
	var params struct {
		IDs []uint `json:"ids" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("请求参数无效: "+err.Error()))
		return
	}

	if len(params.IDs) == 0 {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("至少选择一个角色进行删除"))
		return
	}

	if err := SystemService.Role.BatchDeleteRole(params.IDs); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("批量删除角色失败: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSuccessResponse("批量删除成功"))
}

// @Summary 更新角色菜单
// @Description 更新角色的菜单权限
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param menuInfo body object{roleId=uint,menuIds=[]uint} true "角色菜单信息"
// @Success 200 {object} utils.Response[string] "成功"
// @Failure 400 {object} utils.Response[string] "错误"
// @Security ApiKeyAuth
// @Router /roles/menus [post]
func (c *RoleController) UpdateRoleMenu(ctx *gin.Context) {
	var params struct {
		RoleID  uint   `json:"roleId" binding:"required"`
		MenuIDs []uint `json:"menuIds" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("请求参数无效: "+err.Error()))
		return
	}

	if err := SystemService.Role.UpdateRoleMenu(params.RoleID, params.MenuIDs); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("更新角色菜单失败: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSuccessResponse("更新成功"))
}

// @Summary 更新角色按钮权限
// @Description 更新角色的按钮权限
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param buttonInfo body object{roleId=uint,buttonIds=[]uint} true "角色按钮信息"
// @Success 200 {object} utils.Response[string] "成功"
// @Failure 400 {object} utils.Response[string] "错误"
// @Security ApiKeyAuth
// @Router /roles/buttons [post]
func (c *RoleController) UpdateRoleButton(ctx *gin.Context) {
	var params struct {
		RoleID    uint   `json:"roleId" binding:"required"`
		ButtonIDs []uint `json:"buttonIds" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("请求参数无效: "+err.Error()))
		return
	}

	if err := SystemService.Role.UpdateRoleButton(params.RoleID, params.ButtonIDs); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("更新角色按钮权限失败: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSuccessResponse("更新成功"))
}
