package handlers

import (
	"encoding/json"
	services "farms-planning/app/services/other"
	"farms-planning/pkg/mymodels"
	"farms-planning/pkg/runapp/fiber/logger"

	"github.com/gofiber/fiber/v2"
)

type handlers struct {
	log logger.LoggerFiber
	s   services.Services
}

func NewHandlers(log logger.LoggerFiber, s services.Services) Handlers {
	return &handlers{
		log: log,
		s:   s,
	}
}

func (h *handlers) Hello(c *fiber.Ctx) error {
	return c.Status(200).JSON(mymodels.SetResponse("success", "hello plan farm"))
}

func (h *handlers) ProvinceGet(c *fiber.Ctx) error {
	result, err := h.s.ProvinceGetAll(c.Context())
	if err != nil {
		return err
	}
	return c.Status(200).JSON(mymodels.SetResponse("success", result))
}

func (h *handlers) WeatherGetAll(c *fiber.Ctx) error {
	result, err := h.s.ProvinceGetAll(c.Context())
	if err != nil {
		return err
	}
	return c.Status(200).JSON(mymodels.SetResponse("success", result))
}

func (h *handlers) GeneratePDF(c *fiber.Ctx) error {
	var mybody map[string]interface{}
	getbody := c.Body()
	if err := json.Unmarshal(getbody, &mybody); err != nil {
		return err
	}
	pdfBytes, err := h.s.GeneratePDF(c.Context(), mybody)
	if err != nil {
		return err
	}

	// Set response headers
	c.Set("Content-Disposition", "inline; filename=output.pdf")
	c.Set("Content-Type", "application/pdf")

	// Send PDF as response
	return c.Send(pdfBytes)
}
