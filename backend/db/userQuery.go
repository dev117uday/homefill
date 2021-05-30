package db

import (
	"database/sql"
	"fmt"
	"homefill/backend/errset"
	"homefill/backend/logs"
	"homefill/backend/model"
)

func GetUserFromId(id string) (*model.User, error) {

	user := model.User{}

	query := fmt.Sprintf("select * from user_info where UserId = '%s'", id)
	err := DB.QueryRow(query).Scan(&user.ID, &user.Name, &user.Picture)

	switch {
	case err == sql.ErrNoRows:
		logs.LogIt(logs.LogWarn, "GetUserFromId", "user not found", err)
		return nil, errset.ErrNotFound
	case err != nil:
		logs.LogIt(logs.LogWarn, "GetUserFromId", "unable to run query", err)
		return nil, errset.ErrInternalServer
	default:
	}

	return &user, nil
}

func InsertUser(user *model.User) error {

	query := fmt.Sprintf(`
		INSERT INTO user_info (UserId, Name, Picture) VALUES ('%s','%s','%s') 
		ON CONFLICT DO NOTHING ;`,
		user.ID, user.Name, user.Picture)

	result, err := DB.Exec(query)
	if err != nil {
		logs.LogIt(logs.LogWarn, "InsertUser", "unable to run query", err)
		return errset.ErrInternalServer
	}
	_, err = result.RowsAffected()
	if err != nil {
		logs.LogIt(logs.LogWarn, "InsertUser", "unable to insert user", err)
		return errset.ErrInternalServer
	}

	return nil
}
