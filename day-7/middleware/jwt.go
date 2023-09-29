package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	UserId uint   `json:"userId"`
	Name   string `json:"name"`
	jwt.RegisteredClaims
}

func GenerateTokenJWT(userId uint, name string) string {
	claims := JwtCustomClaims{
		userId,
		name,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	resultJWT, _ := token.SignedString([]byte(os.Getenv("SECRET_JWT")))
	return resultJWT
}
