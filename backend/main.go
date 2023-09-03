package main

import (
	"os"
	"plan-farm/myconfig/myserver"
	"plan-farm/myconfig/myvar"
	"plan-farm/pkg/myconnect"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	myvar.AppObj.DB = myconnect.NewPostgres(os.Getenv("DB_CONNECT"))
	myserver.MiGrateAllDB(myvar.AppObj.DB)
	app := myserver.New()
	myserver.Route(app)
	myserver.Run(app, os.Getenv("PORT"))
}
