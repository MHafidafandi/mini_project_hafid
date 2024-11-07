package helper

import (
	"fmt"
	"miniproject/configs"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
	Id   string `json:"name"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func CreateToken(userId string, role string) (string, error) {

	claims := &JwtCustomClaims{
		userId,
		role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(configs.Cfg.JWTSecret))
}

func ExtractToken(c echo.Context) (interface{}, error) {
	claims := &JwtCustomClaims{}
	tokenString := c.Request().Header.Get("Authorization")

	formattedTokenString := strings.Replace(tokenString, "Bearer ", "", 1)

	token, err := jwt.ParseWithClaims(formattedTokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(configs.Cfg.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if token.Valid {
		data := map[string]string{
			"user_id": claims.Id,
			"role":    claims.Role,
		}

		return data, nil
	}

	return nil, err
}
