package service

import (
	"fmt"
	"homefill/backend/auth"
	"homefill/backend/db"
	"homefill/backend/model"
)

func GenerateJwtTokenService(user *model.User) (string, error) {

	err := db.InsertUser(user)
	if err != nil {
		return "", fmt.Errorf("%s", err)
	}

	token, err := auth.GenerateJwtToken(user.ID)
	if err != nil {
		return "", fmt.Errorf("%s", err)
	}

	return token, nil
}
