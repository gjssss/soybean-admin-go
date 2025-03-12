package system

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gjssss/soybean-admin-go/utils"
)

type ButtonController struct{}

// 获取所有按钮（GET）
func (c *ButtonController) GetButtons(ctx *gin.Context) {
	buttons, err := SystemService.Button.GetButtons()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("获取按钮列表失败: "+err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(buttons))
}

// 根据角色ID获取按钮（GET）
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

// 获取当前用户的按钮（GET）
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
