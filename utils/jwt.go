package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func GenerateToken(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = email
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return t, nil
}

func GetJWTKey() []byte {
	var key string
	key = os.Getenv("JWT_KEY")

	if key == "" {
		fmt.Println("JWT token is empty")
	}
	return []byte(key)
}
