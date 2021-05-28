package db

import (
	"database/sql"
	"fmt"
	"homefill/backend/model"
	"log"
)

func GetUserFromId(id string) (*model.User, error) {

	user := model.User{}

	query := fmt.Sprintf("select * from user_info where UserId = '%s'", id)
	err := DB.QueryRow(query).Scan(&user.ID, &user.Name, &user.Picture)

	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("user nf")
	case err != nil:
		log.Printf("Query error: %v\n", err)
		return nil, fmt.Errorf("qe")
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
		return err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}
