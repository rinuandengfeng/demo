package jwt

import (
	"fmt"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	secretKey := "rndf key"
	token, err := GenerateToken(secretKey, "rndf")
	if err != nil {
		fmt.Println("Failed to Generate token！", err)
		return
	}
	fmt.Println("Jwt Token: ", token)
}

func TestParse(t *testing.T) {
	secretKey := "rndf key"
	token, err := GenerateToken(secretKey, "rndf")
	if err != nil {
		fmt.Println("Failed to Generate token！", err)
		return
	}
	fmt.Println("Jwt Token: ", token)

	claims, err := Parse(token, secretKey)
	if err != nil {
		fmt.Println("Failed to Decode! ", err)
		return
	}
	fmt.Println("userClaims: ", claims)

}
