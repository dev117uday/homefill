package auth

import (
	"fmt"
	config "homefill/backend/configs"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserId string
	jwt.StandardClaims
}

// GenerateJwtToken : generates jwt token from given user id with expire time 3 hours
func GenerateJwtToken(userId string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["userid"] = userId
	claims["time"] = time.Now().Add(time.Hour * 3)

	tokenString, err := token.SignedString(config.JWT_KEY)

	if err != nil {
		return "", fmt.Errorf("something went wrong %s", err)
	}
	return tokenString, nil
}
