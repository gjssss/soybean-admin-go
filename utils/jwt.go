package utils

import (
	"time"

	"github.com/gjssss/soybean-admin-go/global"
	"github.com/golang-jwt/jwt"
	"github.com/patrickmn/go-cache"
)

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

var blacklistCache = cache.New(15*time.Minute, 5*time.Minute)

func GenerateTokens(userID uint) (string, string, error) {
	// JWT Key 等于 JWT Key + 启动时间，这样可以保证每次启动服务 JWT Key 都不一样，自动失效之前的 Token
	// var jwtKey = []byte(global.Config.Secret.JwtKey + strconv.FormatInt(global.Config.Secret.StartTime, 10))
	// var jwtKey = []byte(global.Config.Secret.JwtKey + strconv.FormatInt(global.Config.Secret.StartTime, 10))
	var jwtKey = []byte(global.Config.Secret.JwtKey)
	// Access Token (15分钟过期)
	accessClaims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
			Issuer:    "soybean-admin-go",
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessString, err := accessToken.SignedString(jwtKey)
	if err != nil {
		return "", "", err
	}

	// Refresh Token (7天过期)
	refreshClaims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(168 * time.Hour).Unix(),
			Issuer:    "soybean-admin-go",
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshString, err := refreshToken.SignedString(jwtKey)

	return accessString, refreshString, err
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// return []byte(global.Config.Secret.JwtKey + strconv.FormatInt(global.Config.Secret.StartTime, 10)), nil
		return []byte(global.Config.Secret.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}

func RefreshToken(tokenString string, refreshToken string) (string, string, error) {
	claims, err := ParseToken(refreshToken)
	if err != nil {
		return "", "", err
	}
	blacklistCache.Set(tokenString, true, 15*time.Minute)
	accessString, refreshString, err := GenerateTokens(claims.UserID)
	if err != nil {
		return "", "", err
	}
	return accessString, refreshString, nil
}

func CheckBlacklist(tokenString string) bool {
	_, found := blacklistCache.Get(tokenString)
	return found
}
