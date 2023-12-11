package appinit

import (
	"farms-planning/app/handlers"
	"farms-planning/app/repo"
	serETC "farms-planning/app/services/other"
	serWRC "farms-planning/app/services/water_rain_collect"
	serWR "farms-planning/app/services/water_require"
	"farms-planning/pkg/runapp/fiber/logger"

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

	logz := logger.NewLogger()

	r := repo.NewRepositories(logz, db)
	rwr := repo.NewRepoWaterRequire(logz, db)
	rdv := repo.NewRepoWaterValue(logz, db)

	s := serETC.NewServices(logz, r)
	swr := serWR.NewServiceWaterRequire(logz, r, rwr)
	swrc := serWRC.NewWaterRainCol(logz, r, rdv)

	h := handlers.NewHandlers(logz, s)
	hwr := handlers.NewHandlersWaterRequire(logz, swr)
	hwrc := handlers.NewHandlersWaterCapacity(logz, swrc)

	myroute := []appRoute{}
	myroute = append(myroute, GetRouteWaterCapacity(hwrc)...)
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

func GetRouteWaterCapacity(h handlers.HandlersWaterCapacity) []appRoute {
	return []appRoute{
		{
			RoutePath:   "/capacity/avgrunoffyear",
			RouteMethod: fiber.MethodGet,
			RouteFunc:   h.GetAverageRunoffPerYearAll,
		},
		{
			RoutePath:   "/capacity/avgrunoffyear",
			RouteMethod: fiber.MethodPost,
			RouteFunc:   h.GetAverageRunoffPerYearByProvinceID,
		},
		{
			RoutePath:   "/capacity/surfacearea",
			RouteMethod: fiber.MethodPost,
			RouteFunc:   h.CalSurfaceAreaAtReservoirLevel,
		},
		{
			RoutePath:   "/capacity/areareceiverainwater",
			RouteMethod: fiber.MethodPost,
			RouteFunc:   h.CalareaReceivesRainwater,
		},
		{
			RoutePath:   "/capacity",
			RouteMethod: fiber.MethodPost,
			RouteFunc:   h.CalWaterCopacity,
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
		{
			RoutePath:   "/weather",
			RouteMethod: fiber.MethodGet,
			RouteFunc:   h.WeatherGetAll,
		},
		{
			RoutePath:   "/pdf",
			RouteMethod: fiber.MethodPost,
			RouteFunc:   h.GeneratePDF,
		},
	}
}
