package system

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gjssss/soybean-admin-go/models/system"
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

func (c *MenuController) CreateMenu(ctx *gin.Context) {
	var menu system.Menu
	if err := ctx.ShouldBindJSON(&menu); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("无效的菜单信息: "+err.Error()))
		return
	}

	err := SystemService.Menu.CreateMenu(&menu)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(menu))
}

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

	err := SystemService.Menu.UpdateMenu(&menu)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(menu))
}

func (c *MenuController) DeleteMenu(ctx *gin.Context) {
	idStr := ctx.Query("id")
	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("菜单ID不能为空"))
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("菜单ID必须为数字"))
		return
	}

	err = SystemService.Menu.DeleteMenu(uint(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(""))
}

func (c *MenuController) BatchDeleteMenu(ctx *gin.Context) {
	var ids []uint
	if err := ctx.ShouldBindJSON(&ids); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("无效的菜单ID列表: "+err.Error()))
		return
	}

	if len(ids) == 0 {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("菜单ID列表不能为空"))
		return
	}

	err := SystemService.Menu.BatchDeleteMenu(ids)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(""))
}
