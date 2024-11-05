package helper

import (
	"miniproject/configs"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtCustomClaims struct {
	Id   string `json:"name"`
	Role string `json:"admin"`
	jwt.RegisteredClaims
}

func CreateToken(userId string, role string) (string, error) {

	claims := &jwtCustomClaims{
		userId,
		role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(configs.Cfg.JWTSecret))
}
