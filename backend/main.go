package main

import (
	config "homefill/backend/configs"
	"homefill/backend/db"
	"homefill/backend/routes"
)

func main() {
	config.LoadConfigs()
	db.ConnectTODB()
	db.RunDbScripts()
	routes.StartServer()
}
