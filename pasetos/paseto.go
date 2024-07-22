package pasetos

import (
	"github.com/o1egl/paseto"
	"time"
)

// GenerateToken 生成token
func GenerateToken(secretKey []byte) (string, error) {

	// 创建一个 PASETO V1 实例
	v1 := paseto.NewV1()

	// 定义令牌内容
	jsonToken := paseto.JSONToken{
		Audience:   "rndf.me",                      // 受众
		Issuer:     "rndf.me",                      // 签发者
		Subject:    "test",                         //主题
		IssuedAt:   time.Now(),                     // 签发时间
		Expiration: time.Now().Add(24 * time.Hour), // 过期是时间
	}

	footer := "some footer"

	// 对令牌进行签名
	token, err := v1.Encrypt(secretKey, jsonToken, footer)

	return token, err

}

func Decode(token string, secretKey []byte) (paseto.JSONToken, string, error) {
	v1 := paseto.NewV1()
	// 解密令牌
	var newJsonToken paseto.JSONToken
	var newFooter string
	err := v1.Decrypt(token, secretKey, &newJsonToken, &newFooter)

	return newJsonToken, newFooter, err
}
