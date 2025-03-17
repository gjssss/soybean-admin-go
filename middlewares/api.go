package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gjssss/soybean-admin-go/repositories"
	"github.com/gjssss/soybean-admin-go/utils"
	"github.com/gjssss/soybean-admin-go/utils/cache"
)

func ApiMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method
		uid, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusBadRequest, utils.NewErrorResponse("无法获取用户ID"))
			c.Abort()
		}
		roleIds, err := repositories.System.User.GetUserRoleIds(uid.(uint))
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.NewErrorResponse("获取用户角色失败: "+err.Error()))
			c.Abort()
		}
		c.Set("roleIds", roleIds)
		flag := false
		for _, roleId := range roleIds {
			if cache.ApiCache.Has(path, method, roleId) {
				flag = true
				c.Abort()
			}
		}
		if !flag {
			c.JSON(http.StatusForbidden, utils.NewErrorResponse("无权限访问"))
			c.Abort()
		} else {
			c.Next()
		}
	}
}
