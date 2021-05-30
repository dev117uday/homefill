package auth

import (
	"fmt"
	config "homefill/backend/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func VerifyJwt(tokenString string) (string, error) {

	tkn, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, err := token.Method.(*jwt.SigningMethodHMAC); !err {
			return nil, fmt.Errorf("there was an error")
		}
		return config.JWT_KEY, nil
	})
	if err != nil {
		return "", fmt.Errorf("error : %s", err)
	}

	claims := tkn.Claims.(jwt.MapClaims)
	s := claims["time"].(string)
	t, err := time.Parse(time.RFC3339, s)

	if err != nil {
		return "", fmt.Errorf("error : %s", err)
	}

	if time.Now().Before(t) {
		return claims["userid"].(string), nil
	} else {
		return "", fmt.Errorf("time is after")
	}
}
