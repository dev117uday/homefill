package db

import (
	"database/sql"
	"fmt"
	"homefill/backend/model"
	"log"
)

func GetUserFromId(id string) (*model.User, bool) {
	var user model.User
	query := fmt.Sprintf("select * from user_info where UserId = '%s'", id)
	err := DB.QueryRow(query).Scan(&user.ID, &user.Name, &user.Picture)

	switch {
	case err == sql.ErrNoRows:
		return nil, false
	case err != nil:
		log.Fatalf("query error: %v\n", err)
	default:
	}

	return &user, true
}

func InsertUser(user *model.User) bool {
	query := fmt.Sprintf("insert into user_info (UserId, Name, Picture) VALUES ('%s','%s','%s');", user.ID, user.Name, user.Picture)
	result, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if rows != 1 {
		return false
	}
	return true
}
