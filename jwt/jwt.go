package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type UserClaims struct {
	UserName string `json:"userName"`
	jwt.RegisteredClaims
}

// GenerateToken 生成token
func GenerateToken(secretKey string, username string) (string, error) {
	claims := UserClaims{
		UserName: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "rndf",                                             // 签发人
			Subject:   "test",                                             // 主题
			Audience:  nil,                                                // 受众
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 签发时间
			ID:        "1",                                                // Id
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 进行签名，并获取签名后的字符串
	token, err := tokenClaims.SignedString([]byte(secretKey))
	return token, err
}

// Parse 验证token
func Parse(token string, secretKey string) (*UserClaims, error) {
	// 获取token的claims
	tokenClaims, err := jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	// 判断token是否有效
	if tokenClaims != nil {
		//tokenClaims.Valid 验证token是否有效
		if claims, ok := tokenClaims.Claims.(*UserClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
