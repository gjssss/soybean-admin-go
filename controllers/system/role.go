package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gjssss/soybean-admin-go/models/system"
	"github.com/gjssss/soybean-admin-go/utils"
)

type RoleController struct{}

// 获取所有角色（GET）
func (c *RoleController) GetAllRoles(ctx *gin.Context) {
	roles, err := SystemService.Role.GetAllRoles()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("获取所有角色失败: "+err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(roles))
}

// 获取分页角色（GET）
func (c *RoleController) GetRoles(ctx *gin.Context) {
	page := utils.ParsePagination(ctx)
	roles, count, err := SystemService.Role.GetRoles(page)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("获取角色失败: "+err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(utils.NewPagination(roles, page, count)))
}

// 创建角色（POST）
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

// 更新角色（POST）
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

// 删除角色（POST）
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

// 批量删除角色（POST）
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

// 更新角色菜单（POST）
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

// 更新角色按钮权限（POST）
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
