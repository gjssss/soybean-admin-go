package system

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gjssss/soybean-admin-go/models/system"
	"github.com/gjssss/soybean-admin-go/utils"
)

type ButtonController struct{}

// @Summary 获取所有按钮
// @Description 获取系统所有按钮列表
// @Tags 按钮管理
// @Accept json
// @Produce json
// @Success 200 {object} utils.Response[[]system.Button] "成功"
// @Failure 400 {object} utils.Response[string] "错误"
// @Security ApiKeyAuth
// @Router /buttons [get]
func (c *ButtonController) GetButtons(ctx *gin.Context) {
	buttons, err := SystemService.Button.GetButtons()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("获取按钮列表失败: "+err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(buttons))
}

// @Summary 获取角色按钮
// @Description 获取指定角色的按钮列表
// @Tags 按钮管理
// @Accept json
// @Produce json
// @Param roleId query uint true "角色ID"
// @Success 200 {object} utils.Response[[]system.Button] "成功"
// @Failure 400 {object} utils.Response[string] "错误"
// @Security ApiKeyAuth
// @Router /buttons/role [get]
func (c *ButtonController) GetButtonsByRoleId(ctx *gin.Context) {
	rid := ctx.Query("roleId")
	if rid == "" {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("roleId不能为空"))
		return
	}

	roleId, err := strconv.ParseUint(rid, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("roleId必须为数字"))
		return
	}

	buttons, err := SystemService.Button.GetButtonsByRoleId(uint(roleId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("获取角色按钮失败: "+err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(buttons))
}

// @Summary 获取用户按钮
// @Description 获取当前用户的按钮权限
// @Tags 按钮管理
// @Accept json
// @Produce json
// @Success 200 {object} utils.Response[[]system.Button] "成功"
// @Failure 400 {object} utils.Response[string] "错误"
// @Security ApiKeyAuth
// @Router /buttons/user [get]
func (c *ButtonController) GetUserButtons(ctx *gin.Context) {
	uid, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("无法获取用户ID"))
		return
	}

	buttons, err := SystemService.Button.GetButtonsByUserId(uid.(uint))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("获取用户按钮失败: "+err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(buttons))
}

// @Summary 创建按钮
// @Description 创建新的按钮
// @Tags 按钮管理
// @Accept json
// @Produce json
// @Param button body system.Button true "按钮信息"
// @Success 200 {object} utils.Response[system.Button] "成功"
// @Failure 400 {object} utils.Response[string] "错误"
// @Security ApiKeyAuth
// @Router /buttons [post]
func (c *ButtonController) CreateButton(ctx *gin.Context) {
	var button system.Button
	if err := ctx.ShouldBindJSON(&button); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("请求参数无效: "+err.Error()))
		return
	}

	if button.Code == "" {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("按钮代码不能为空"))
		return
	}

	if err := SystemService.Button.CreateButton(&button); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("创建按钮失败: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(button))
}

// @Summary 更新按钮
// @Description 更新按钮信息
// @Tags 按钮管理
// @Accept json
// @Produce json
// @Param button body system.Button true "按钮信息"
// @Success 200 {object} utils.Response[string] "成功"
// @Failure 400 {object} utils.Response[string] "错误"
// @Security ApiKeyAuth
// @Router /buttons/update [post]
func (c *ButtonController) UpdateButton(ctx *gin.Context) {
	var button system.Button
	if err := ctx.ShouldBindJSON(&button); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("请求参数无效: "+err.Error()))
		return
	}

	if button.ID == 0 {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("按钮ID不能为空"))
		return
	}

	if button.Code == "" {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("按钮代码不能为空"))
		return
	}

	if err := SystemService.Button.UpdateButton(&button); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("更新按钮失败: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSuccessResponse("更新成功"))
}

// @Summary 删除按钮
// @Description 删除指定按钮
// @Tags 按钮管理
// @Accept json
// @Produce json
// @Param id body object{id=uint} true "按钮ID"
// @Success 200 {object} utils.Response[string] "成功"
// @Failure 400 {object} utils.Response[string] "错误"
// @Security ApiKeyAuth
// @Router /buttons/delete [post]
func (c *ButtonController) DeleteButton(ctx *gin.Context) {
	var params struct {
		ID uint `json:"id" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("请求参数无效: "+err.Error()))
		return
	}

	if err := SystemService.Button.DeleteButton(params.ID); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("删除按钮失败: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSuccessResponse("删除成功"))
}

// @Summary 批量删除按钮
// @Description 批量删除多个按钮
// @Tags 按钮管理
// @Accept json
// @Produce json
// @Param ids body object{ids=[]uint} true "按钮ID列表"
// @Success 200 {object} utils.Response[string] "成功"
// @Failure 400 {object} utils.Response[string] "错误"
// @Security ApiKeyAuth
// @Router /buttons/batchDelete [post]
func (c *ButtonController) BatchDeleteButton(ctx *gin.Context) {
	var params struct {
		IDs []uint `json:"ids" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("请求参数无效: "+err.Error()))
		return
	}

	if len(params.IDs) == 0 {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("至少选择一个按钮进行删除"))
		return
	}

	if err := SystemService.Button.BatchDeleteButton(params.IDs); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("批量删除按钮失败: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSuccessResponse("批量删除成功"))
}
