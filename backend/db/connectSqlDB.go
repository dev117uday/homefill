package db

import (
	"database/sql"
	config "homefill/backend/configs"
	"log"

	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

func ConnectTODB() {
	var err error
	DB, err = sql.Open("postgres", config.PGSQL_CS)
	if err != nil {
		log.Fatal(err)
	}
}

func RunDbScripts() {
	_, err := DB.Exec(`
		create table if not exists user_info (
			UserId varchar(30) primary key,
			Name       varchar(30) not null,
			Picture    varchar(90) not null
		);
	`)

	if err != nil {
		log.Fatal("unable to connect to db")
	}
}
