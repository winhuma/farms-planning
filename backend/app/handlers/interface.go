package handlers

import "github.com/gofiber/fiber/v2"

type Handlers interface {
	Hello(c *fiber.Ctx) error
	ProvinceGet(c *fiber.Ctx) error
}

type HandlersWaterRequire interface {
	WaterRequireDataGet(c *fiber.Ctx) error
	WaterDayCal(c *fiber.Ctx) error
	WaterIndustryCal(c *fiber.Ctx) error
	WaterPlantCal(c *fiber.Ctx) error
	WaterAnimalCal(c *fiber.Ctx) error
}
