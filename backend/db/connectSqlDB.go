package db

import (
	"database/sql"
	config "homefill/backend/config"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var (
	DB *sql.DB
)

func ConnectTODB() {

	db, _ := sql.Open("postgres", config.PGSQL_CS)
	_, err := db.Exec("select VERSION();")
	if err != nil {
		config.Log.WithFields(logrus.Fields{
			"fn":  "ConnectTODB",
			"err": err.Error(),
		}).Fatal("unable to connect to db")
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
		config.Log.WithFields(logrus.Fields{
			"fn":  "RunDBScripts",
			"err": err.Error(),
		}).Fatal("unable to create db table")
	}
}
