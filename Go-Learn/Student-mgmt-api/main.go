package main

import (
	"students-mgmt-api/config"
	"students-mgmt-api/routes"
)

func main() {

	config.ConnectDB()
	//go channels
	//Graceful shutdown
	r := routes.Routes()
	r.Run()
}
