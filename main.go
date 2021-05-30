package main

import (
	config "homefill/backend/config"
	"homefill/backend/db"
	"homefill/backend/routes"
)

func main() {
	config.LoadConfig()
	db.ConnectTODB()
	db.RunDBScripts()
	routes.StartServer()
}
