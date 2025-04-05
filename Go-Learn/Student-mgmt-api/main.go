package main

import (
	"students-mgmt-api/src/config"
	"students-mgmt-api/src/routes"
)

func main() {

	config.ConnectDB()	
	r := routes.Routes()
	r.Run(":3000")
}
