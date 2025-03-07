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
			c.AbortWithStatusJSON(401, gin.H{"error": "未提供认证令牌"})
			return
		}

		token := strings.Split(tokenString, " ")[1]

		// 解析令牌
		claims, err := utils.ParseToken(token)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "无效令牌"})
			return
		}

		// 将用户信息存入上下文
		c.Set("userID", claims.UserID)

		c.Next()
	}
}
