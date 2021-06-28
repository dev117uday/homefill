package main

import (
	config "go-auth-google/backend/config"
	"go-auth-google/backend/db"
	"go-auth-google/backend/routes"
)

func main() {
	config.LoadConfig()

	pgRepository := &db.PostgreSQL{}

	db.DB = db.DataBaseRepo{
		Repo: pgRepository,
	}
	db.DB.Repo.ConnectTODB()
	db.DB.Repo.RunDBScripts()

	routes.StartServer()
}
