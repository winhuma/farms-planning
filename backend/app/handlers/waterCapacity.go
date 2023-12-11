package handlers

import (
	"encoding/json"
	services "farms-planning/app/services/water_rain_collect"
	"farms-planning/myconfig/models"
	"farms-planning/pkg/mymodels"
	"farms-planning/pkg/runapp/fiber/logger"

	"github.com/gofiber/fiber/v2"
)

type hWaterCapacity struct {
	log logger.LoggerFiber
	swc services.ServiceWaterRainCollect
}

func NewHandlersWaterCapacity(log logger.LoggerFiber, swc services.ServiceWaterRainCollect) HandlersWaterCapacity {
	return &hWaterCapacity{
		log: log,
		swc: swc,
	}
}

func (h *hWaterCapacity) CalSurfaceAreaAtReservoirLevel(c *fiber.Ctx) error {
	var mybody models.BodyCalSurfaceArea
	if err := json.Unmarshal(c.Body(), &mybody); err != nil {
		return err
	}
	result, err := h.swc.CalSurfaceAreaAtReservoirLevel(c.Context(), mybody)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(mymodels.SetResponse("success", result))
}

func (h *hWaterCapacity) GetAverageRunoffPerYearAll(c *fiber.Ctx) error {
	result, err := h.swc.GetAverageRunoffPerYearAll(c.Context())
	if err != nil {
		return err
	}
	return c.Status(200).JSON(mymodels.SetResponse("success", result))
}

func (h *hWaterCapacity) GetAverageRunoffPerYearByProvinceID(c *fiber.Ctx) error {
	var mybody models.BodyGetWeatherByProvinceID
	if err := json.Unmarshal(c.Body(), &mybody); err != nil {
		return err
	}
	result, err := h.swc.GetAverageRunoffPerYearByProvinceID(c.Context(), mybody.ProvinceID)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(mymodels.SetResponse("success", result))
}

func (h *hWaterCapacity) CalareaReceivesRainwater(c *fiber.Ctx) error {
	var mybody models.BodyCalAreaReceivesRainwater
	if err := json.Unmarshal(c.Body(), &mybody); err != nil {
		return err
	}
	result, err := h.swc.CalAreaTerraceReceivesRainWater(c.Context(), mybody)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(mymodels.SetResponse("success", result))
}

func (h *hWaterCapacity) CalAverageRunoffPerYear(c *fiber.Ctx) error {
	var mybody models.BodyCalAverageRunoffPerYear
	if err := json.Unmarshal(c.Body(), &mybody); err != nil {
		return err
	}
	result, err := h.swc.CalAverageRunoffPerYear(c.Context(), mybody)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(mymodels.SetResponse("success", result))
}

func (h *hWaterCapacity) CalWaterLostFromEvaporation(c *fiber.Ctx) error {
	var mybody models.BodyCalWaterLostFromEvaporation
	if err := json.Unmarshal(c.Body(), &mybody); err != nil {
		return err
	}
	result, err := h.swc.CalWaterLostFromEvaporation(c.Context(), mybody)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(mymodels.SetResponse("success", result))
}

func (h *hWaterCapacity) CalWaterLostFromLeakage(c *fiber.Ctx) error {
	var mybody models.BodyCalWaterLostFromLeakage
	if err := json.Unmarshal(c.Body(), &mybody); err != nil {
		return err
	}
	result, err := h.swc.CalWaterLostFromLeakage(c.Context(), mybody)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(mymodels.SetResponse("success", result))
}

func (h *hWaterCapacity) CalWaterCopacity(c *fiber.Ctx) error {
	var mybody models.BodyCalWaterCapacity
	if err := json.Unmarshal(c.Body(), &mybody); err != nil {
		return err
	}
	result, err := h.swc.CalWaterCapacity(c.Context(), mybody)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(mymodels.SetResponse("success", result))
}
