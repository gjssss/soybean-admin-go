package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gjssss/soybean-admin-go/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusOK, utils.NewLogoutModelResponse("7777", "请先登录"))
			return
		}

		token := strings.Split(tokenString, " ")[1]
		c.Set("accessToken", token)

		if utils.CheckBlacklist(token) {
			c.AbortWithStatusJSON(http.StatusOK, utils.NewLogoutModelResponse("8888", "登录已过期"))
			return
		}

		claims, err := utils.ParseToken(token)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, utils.NewLogoutModelResponse("8888", "登录已过期"))
			return
		}

		// 将用户信息存入上下文
		c.Set("userID", claims.UserID)

		c.Next()
	}
}
