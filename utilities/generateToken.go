package utilities

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/vinitparekh17/project-x/config"
	"github.com/vinitparekh17/project-x/handler"
)

type JWTConfig struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GetSecret() string {
	secret, er := config.GetEnv("JWT_SECRET")
	handler.ErrorHandler(er)
	return secret
}

func GenerateJWT(email string) string {
	customClaims := JWTConfig{
		email,
		jwt.RegisteredClaims{
			Issuer:    "Authx",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	secret := GetSecret()
	tokenString, err := token.SignedString([]byte(secret))
	handler.ErrorHandler(err)
	return tokenString
}

func VerifyToken(token string) (status bool) {
	secret := GetSecret()
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	handler.ErrorHandler(err)
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		fmt.Println(claims["email"], claims["exp"])
		return true
	}
	return false
}
