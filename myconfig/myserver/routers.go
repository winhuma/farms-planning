package myserver

import (
	"plan-farm/app/handlers"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	app.Get("/", handlers.Hello)
	app.Get("/waters/:action", handlers.WaterRequireDataGet)
	app.Post("/water/calculate/area", handlers.WaterAreaCal)
	app.Post("/water/calculate/industry", handlers.WaterIndustryCal)
	app.Post("/water/calculate/plant", handlers.WaterPlantCal)
}
