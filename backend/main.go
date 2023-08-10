package main

import (
	"os"
	"plan-farm/myconfig/myserver"
	"plan-farm/pkg/myconnect"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	myconnect.NewPostgres(os.Getenv("DB_CONNECT"))
	app := myserver.New()
	myserver.Route(app)
	myserver.Run(app, os.Getenv("PORT"))
}
