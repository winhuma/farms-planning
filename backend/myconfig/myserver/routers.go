package myserver

import (
	"plan-farm/app/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type appRoute struct {
	RoutePath   string
	RouteMethod string
	RouteFunc   func(*fiber.Ctx) error
}

func Route(app *fiber.App) {

	myroute := []appRoute{}
	myroute = append(myroute, GetRouteWater()...)
	myroute = append(myroute, GetRouteOther()...)
	for _, handle := range myroute {
		log.Info().Msgf("[%s] %s", handle.RouteMethod, handle.RoutePath)
		app.Add(handle.RouteMethod, handle.RoutePath, handle.RouteFunc)
	}
}

func GetRouteWater() []appRoute {
	return []appRoute{
		{
			RoutePath:   "/",
			RouteMethod: fiber.MethodGet,
			RouteFunc:   handlers.Hello,
		},
		{
			RoutePath:   "/waters",
			RouteMethod: fiber.MethodGet,
			RouteFunc:   handlers.WaterRequireDataGet,
		},
		{
			RoutePath:   "/waters/calculate/day",
			RouteMethod: fiber.MethodPost,
			RouteFunc:   handlers.WaterDayCal,
		},
		{
			RoutePath:   "/waters/calculate/industry",
			RouteMethod: fiber.MethodPost,
			RouteFunc:   handlers.WaterIndustryCal,
		},
		{
			RoutePath:   "/waters/calculate/plant",
			RouteMethod: fiber.MethodPost,
			RouteFunc:   handlers.WaterPlantCal,
		},
		{
			RoutePath:   "/waters/calculate/animal",
			RouteMethod: fiber.MethodPost,
			RouteFunc:   handlers.WaterAnimalCal,
		},
	}
}

func GetRouteOther() []appRoute {
	return []appRoute{
		{
			RoutePath:   "/province",
			RouteMethod: fiber.MethodGet,
			RouteFunc:   handlers.ProvinceGet,
		},
	}
}
