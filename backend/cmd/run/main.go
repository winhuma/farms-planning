package main

import (
	"farms-planning/myconfig/myserver"
	"farms-planning/myconfig/myvar"
	"farms-planning/pkg/myconnect"
)

func main() {
	myvar.SetEnv()
	db := myconnect.NewPostgres(myvar.ENV_DB_CONNECT)
	myserver.MiGrateAllDB(db)
	app := myserver.New()
	myserver.Route(app, db)
	myserver.Run(app, myvar.ENV_PORT)
}
