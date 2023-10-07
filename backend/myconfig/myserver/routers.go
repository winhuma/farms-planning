package myserver

import (
	"farms-planning/app/handlers"
	"farms-planning/app/repo"
	serETC "farms-planning/app/services/other"
	serWRC "farms-planning/app/services/water_rain_collect"
	serWR "farms-planning/app/services/water_require"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type appRoute struct {
	RoutePath   string
	RouteMethod string
	RouteFunc   func(*fiber.Ctx) error
}

func Route(app *fiber.App, db *gorm.DB) {

	r := repo.NewRepositories(db)
	rwr := repo.NewRepoWaterRequire(db)

	s := serETC.NewServices(r)
	swr := serWR.NewServiceWaterRequire(r, rwr)
	swrc := serWRC.NewWaterRainCol()

	h := handlers.NewHandlers(s)
	hwr := handlers.NewHandlersWaterRequire(swr)
	_ = swrc

	myroute := []appRoute{}
	myroute = append(myroute, GetRouteWater(hwr)...)
	myroute = append(myroute, GetRouteOther(h)...)
	for _, handle := range myroute {
		log.Info().Msgf("[%s] %s", handle.RouteMethod, handle.RoutePath)
		app.Add(handle.RouteMethod, handle.RoutePath, handle.RouteFunc)
	}
}

func GetRouteWater(h handlers.HandlersWaterRequire) []appRoute {
	return []appRoute{
		{
			RoutePath:   "/waters",
			RouteMethod: fiber.MethodGet,
			RouteFunc:   h.WaterRequireDataGet,
		},
		{
			RoutePath:   "/waters/calculate/day",
			RouteMethod: fiber.MethodPost,
			RouteFunc:   h.WaterDayCal,
		},
		{
			RoutePath:   "/waters/calculate/industry",
			RouteMethod: fiber.MethodPost,
			RouteFunc:   h.WaterIndustryCal,
		},
		{
			RoutePath:   "/waters/calculate/plant",
			RouteMethod: fiber.MethodPost,
			RouteFunc:   h.WaterPlantCal,
		},
		{
			RoutePath:   "/waters/calculate/animal",
			RouteMethod: fiber.MethodPost,
			RouteFunc:   h.WaterAnimalCal,
		},
	}
}

func GetRouteOther(h handlers.Handlers) []appRoute {
	return []appRoute{
		{
			RoutePath:   "/",
			RouteMethod: fiber.MethodGet,
			RouteFunc:   h.Hello,
		},
		{
			RoutePath:   "/province",
			RouteMethod: fiber.MethodGet,
			RouteFunc:   h.ProvinceGet,
		},
	}
}
