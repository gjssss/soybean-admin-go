package system

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gjssss/soybean-admin-go/utils"
)

type ButtonController struct{}

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
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(buttons))
}

func (c *ButtonController) GetButtons(ctx *gin.Context) {
	buttons, err := SystemService.Button.GetButtons()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(buttons))
}

func (c *ButtonController) GetUserButtons(ctx *gin.Context) {
	uid, _ := ctx.Get("userID")
	buttons, err := SystemService.Button.GetButtonsByUserId(uid.(uint))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(buttons))
}
