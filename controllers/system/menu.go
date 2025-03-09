package system

import (
	"net/http"

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
