package auth

import (
	config "homefill/backend/configs"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserId string
	jwt.StandardClaims
}

// TODO error handling
func GenerateJwtToken(userId string) (string, error) {
	expirationTime := time.Now().Add(time.Minute * 5)

	claims := &Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := []byte(config.JWT_KEY)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
