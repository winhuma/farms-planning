package main

import (
	"os"
	"plan-farm/libs/connection"
	"plan-farm/myconfig/myserver"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	connection.FirebaseConnect()
	app := myserver.New()
	myserver.Route(app)
	myserver.Run(app, os.Getenv("PORT"))
}
