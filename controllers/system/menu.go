package system

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gjssss/soybean-admin-go/utils"
)

type MenuController struct{}

func (c *MenuController) GetUserMenus(ctx *gin.Context) {
	uid, _ := ctx.Get("userID")
	menus, err := SystemService.Menu.GetMenusByUserId(uid.(uint))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(menus))
}

func (c *MenuController) GetMenus(ctx *gin.Context) {
	menus, err := SystemService.Menu.GetMenus()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(menus))
}

func (c *MenuController) GetMenusByRoleId(ctx *gin.Context) {
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
	menus, err := SystemService.Menu.GetMenusByRoleId(uint(roleId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(menus))
}
