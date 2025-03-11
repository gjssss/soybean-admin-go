package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
