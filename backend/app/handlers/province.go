package handlers

import (
	"plan-farm/app/services"
	"plan-farm/pkg/models"

	"github.com/gofiber/fiber/v2"
)

func ProvinceGet(c *fiber.Ctx) error {
	result, err := services.ProvinceGetAll()
	if err != nil {
		return err
	}
	return c.Status(200).JSON(models.ResponseSuccess("success", result))
}
