package main

import (
	"farms-planning/myconfig/appinit"
	"farms-planning/myconfig/myvar"
	"farms-planning/pkg/myconnect"
	"farms-planning/pkg/runapp/fiber/server"
)

func main() {
	myvar.SetEnv()
	db := myconnect.NewPostgres(myvar.ENV_DB_CONNECT)
	appinit.MiGrateAllDB(db)
	app := server.New()
	appinit.Route(app, db)
	server.Run(app, myvar.ENV_PORT)
}
