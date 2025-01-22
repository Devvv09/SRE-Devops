package main

import (
	"students-mgmt-api/config"
	"students-mgmt-api/routes"
)

func main() {

	config.ConnectDB()
	r := routes.Routes()
	r.Run()

}
