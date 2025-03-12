package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gjssss/soybean-admin-go/models/system"
	"github.com/gjssss/soybean-admin-go/utils"
)

type RoleController struct{}

func (c *RoleController) GetRoles(ctx *gin.Context) {
	page := utils.ParsePagination(ctx)
	roles, count, err := SystemService.Role.GetRoles(page)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(utils.NewPagination(roles, page, count)))
}

func (c *RoleController) GetAllRoles(ctx *gin.Context) {
	roles, err := SystemService.Role.GetAllRoles()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(roles))
}

func (c *RoleController) CreateRole(ctx *gin.Context) {
	var role system.Role
	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("请求参数无效"))
		return
	}

	if err := SystemService.Role.CreateRole(&role); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(role))
}

func (c *RoleController) UpdateRole(ctx *gin.Context) {
	var role system.Role
	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("请求参数无效"))
		return
	}

	if err := SystemService.Role.UpdateRole(&role); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(""))
}

func (c *RoleController) DeleteRole(ctx *gin.Context) {
	var params struct {
		ID uint `json:"id" form:"id"`
	}

	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("请求参数无效"))
		return
	}

	if err := SystemService.Role.DeleteRole(params.ID); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(""))
}

func (c *RoleController) BatchDeleteRole(ctx *gin.Context) {
	var ids []uint
	if err := ctx.ShouldBind(&ids); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("请求参数无效"))
		return
	}

	if err := SystemService.Role.BatchDeleteRole(ids); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(""))
}

func (c *RoleController) UpdateRoleMenu(ctx *gin.Context) {
	var params struct {
		RoleID  uint   `json:"roleId"`
		MenuIDs []uint `json:"menuIds"`
	}

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("请求参数无效"))
		return
	}

	if err := SystemService.Role.UpdateRoleMenu(params.RoleID, params.MenuIDs); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(""))
}

func (c *RoleController) UpdateRoleButton(ctx *gin.Context) {
	var params struct {
		RoleID    uint   `json:"roleId"`
		ButtonIDs []uint `json:"buttonIds"`
	}

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("请求参数无效"))
		return
	}

	if err := SystemService.Role.UpdateRoleButton(params.RoleID, params.ButtonIDs); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(""))
}
