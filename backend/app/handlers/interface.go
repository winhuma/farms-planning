package handlers

import "github.com/gofiber/fiber/v2"

type Handlers interface {
	Hello(c *fiber.Ctx) error
	ProvinceGet(c *fiber.Ctx) error
	WeatherGetAll(c *fiber.Ctx) error
	GeneratePDF(c *fiber.Ctx) error
}

type HandlersWaterRequire interface {
	WaterRequireDataGet(c *fiber.Ctx) error
	WaterDayCal(c *fiber.Ctx) error
	WaterIndustryCal(c *fiber.Ctx) error
	WaterPlantCal(c *fiber.Ctx) error
	WaterAnimalCal(c *fiber.Ctx) error
}

type HandlersWaterCapacity interface {
	GetAverageRunoffPerYearAll(c *fiber.Ctx) error
	GetAverageRunoffPerYearByProvinceID(c *fiber.Ctx) error
	CalSurfaceAreaAtReservoirLevel(c *fiber.Ctx) error
	CalareaReceivesRainwater(c *fiber.Ctx) error
	CalAverageRunoffPerYear(c *fiber.Ctx) error
	CalWaterLostFromEvaporation(c *fiber.Ctx) error
	CalWaterCopacity(c *fiber.Ctx) error
}
