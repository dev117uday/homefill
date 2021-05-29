package db

import (
	"database/sql"
	"fmt"
	"homefill/backend/model"
)

func GetUserFromId(id string) (*model.User, error) {

	user := model.User{}

	query := fmt.Sprintf("select * from user_info where UserId = '%s'", id)
	err := DB.QueryRow(query).Scan(&user.ID, &user.Name, &user.Picture)

	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("error : %s", err)
	case err != nil:
		return nil, fmt.Errorf("error : %s", err)
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
		return fmt.Errorf("error : %s", err)
	}
	_, err = result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error : %s", err)
	}

	return nil
}
