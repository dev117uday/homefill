package service

import (
	"errors"
	"homefill/backend/auth"
	"homefill/backend/db"
	"homefill/backend/model"
)

func GenerateJwtTokenService(user *model.User) (string, error) {

	err := db.InsertUser(user)
	if err != nil {
		return "", errors.New("error inserting user")
	}

	token, err := auth.GenerateJwtToken(user.ID)

	if err != nil {
		return "", errors.New("error generating jwt")
	}

	return token, nil
}
