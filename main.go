package main

import (
	"jk/go-sportsapp/database"
	"jk/go-sportsapp/server"
)

func main() {
	database.ConnentDB()

	r := server.ConnectServerRouter()
	r.Run(":8080")
}
