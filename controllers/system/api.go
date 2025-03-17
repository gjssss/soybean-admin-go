package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gjssss/soybean-admin-go/models/system"
	"github.com/gjssss/soybean-admin-go/utils"
)

type ApiController struct{}

// @Summary 创建API接口
// @Description 创建新的API接口
// @Tags API管理
// @Accept json
// @Produce json
// @Param api body system.Api true "API信息"
// @Success 200 {object} utils.Response[system.Api] "成功"
// @Failure 400 {object} utils.Response[string] "错误"
// @Security ApiKeyAuth
// @Router /apis [post]
func (c *ApiController) CreateApi(ctx *gin.Context) {
	var api system.Api
	if err := ctx.ShouldBindJSON(&api); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("参数错误："+err.Error()))
		return
	}

	if err := SystemService.Api.CreateApi(&api); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("创建接口失败："+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(api))
}

// @Summary 更新API接口
// @Description 更新API接口信息
// @Tags API管理
// @Accept json
// @Produce json
// @Param api body system.Api true "API信息"
// @Success 200 {object} utils.Response[system.Api] "成功"
// @Failure 400 {object} utils.Response[string] "错误"
// @Security ApiKeyAuth
// @Router /apis/update [post]
func (c *ApiController) UpdateApi(ctx *gin.Context) {
	var api system.Api
	if err := ctx.ShouldBindJSON(&api); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("参数错误："+err.Error()))
		return
	}

	if err := SystemService.Api.UpdateApi(&api); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("更新接口失败："+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(api))
}

// @Summary 删除API接口
// @Description 删除指定API接口
// @Tags API管理
// @Accept json
// @Produce json
// @Param id body object{id=uint} true "API ID"
// @Success 200 {object} utils.Response[string] "成功"
// @Failure 400 {object} utils.Response[string] "错误"
// @Security ApiKeyAuth
// @Router /apis/delete [post]
func (c *ApiController) DeleteApi(ctx *gin.Context) {
	var req struct {
		ID uint `json:"id"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("参数错误："+err.Error()))
		return
	}

	if err := SystemService.Api.DeleteApi(req.ID); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("删除接口失败："+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSuccessResponse("删除成功"))
}

// @Summary 获取角色API接口
// @Description 根据角色ID获取API接口列表
// @Tags API管理
// @Accept json
// @Produce json
// @Param roleId query uint true "角色ID"
// @Success 200 {object} utils.Response[[]system.Api] "成功"
// @Failure 400 {object} utils.Response[string] "错误"
// @Security ApiKeyAuth
// @Router /apis/role [get]
func (c *ApiController) GetApisByRoleID(ctx *gin.Context) {
	var req struct {
		RoleID uint `form:"roleId"`
	}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("参数错误："+err.Error()))
		return
	}

	apis, err := SystemService.Api.GetApisByRoleID(req.RoleID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("获取接口失败："+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(apis))
}

// @Summary 获取所有API接口
// @Description 获取所有API接口列表
// @Tags API管理
// @Accept json
// @Produce json
// @Success 200 {object} utils.Response[[]system.Api] "成功"
// @Failure 400 {object} utils.Response[string] "错误"
// @Security ApiKeyAuth
// @Router /apis [get]
func (c *ApiController) GetAllApis(ctx *gin.Context) {
	apis, err := SystemService.Api.GetAllApis()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("获取所有接口失败："+err.Error()))
		return
	}
	println(12312321)
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(apis))
}

// @Summary 更新角色API关联
// @Description 更新角色与API的关联关系
// @Tags API管理
// @Accept json
// @Produce json
// @Param apiInfo body object{roleId=uint,ids=[]uint} true "角色API关联信息"
// @Success 200 {object} utils.Response[string] "成功"
// @Failure 400 {object} utils.Response[string] "错误"
// @Security ApiKeyAuth
// @Router /apis/role [post]
func (c *ApiController) UpdateRoleApi(ctx *gin.Context) {
	var req struct {
		RoleID uint   `json:"roleId"`
		IDs    []uint `json:"ids"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("参数错误："+err.Error()))
		return
	}

	if err := SystemService.Api.UpdateRoleApi(req.RoleID, req.IDs); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("更新角色接口关联失败："+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSuccessResponse("更新成功"))
}
