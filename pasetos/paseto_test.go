package pasetos

import (
	"fmt"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	// 创建一个密钥
	secretKey := []byte("Rndf KEY")
	// 生成token
	token, err := GenerateToken(secretKey)
	if err != nil {
		fmt.Println("Failed Generate Token! ", err)
		return
	}
	fmt.Println("Generated Token:", token)
}

func TestDecode(t *testing.T) {
	// 创建一个密钥
	secretKey := []byte("Rndf KEY")
	// 生成token
	token, err := GenerateToken(secretKey)
	if err != nil {
		fmt.Println("Failed Generate Token! ", err)
		return
	}

	// 解密token
	newJsonToken, newFooter, err := Decode(token, secretKey)
	if err != nil {
		fmt.Println("Error decrypting token:", err)
		return
	}

	fmt.Println("Token Content:", newJsonToken)
	fmt.Println("Token Footer:", newFooter)
}
