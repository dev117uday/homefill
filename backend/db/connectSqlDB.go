package db

import (
	"database/sql"
	config "homefill/backend/config"
	"homefill/backend/logs"

	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

func ConnectTODB() {
	db, _ := sql.Open("postgres", config.PGSQL_CS)
	_, err := db.Exec("select VERSION();")
	if err != nil {
		logs.LogIt(logs.LogFatal, "ConnectTODB", "unable to connect to db", err)
	}
	DB = db
}

func RunDBScripts() {
	_, err := DB.Exec(`
		create table if not exists user_info (
			UserId varchar(30) primary key,
			Name       varchar(30) not null,
			Picture    varchar(90) not null
		);
	`)

	if err != nil {
		logs.LogIt(logs.LogFatal, "ConnectTODB", "unable to create db table", err)
	}
}
