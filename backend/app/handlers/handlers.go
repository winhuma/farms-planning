package handlers

import (
	services "farms-planning/app/services/other"
	"farms-planning/pkg/mymodels"

	"github.com/gofiber/fiber/v2"
)

type handlers struct {
	s services.Services
}

func NewHandlers(s services.Services) Handlers {
	return &handlers{
		s: s,
	}
}

func (h *handlers) Hello(c *fiber.Ctx) error {
	return c.Status(200).JSON(mymodels.SetResponse("success", "hello plan farm"))
}

func (h *handlers) ProvinceGet(c *fiber.Ctx) error {
	result, err := h.s.ProvinceGetAll()
	if err != nil {
		return err
	}
	return c.Status(200).JSON(mymodels.SetResponse("success", result))
}
