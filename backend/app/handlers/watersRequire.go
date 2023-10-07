package handlers

import (
	"encoding/json"
	services "farms-planning/app/services/water_require"
	"farms-planning/myconfig/models"
	"farms-planning/myconfig/myvar"
	"farms-planning/pkg/myfunc"
	"farms-planning/pkg/mymodels"

	"github.com/gofiber/fiber/v2"
)

type hWaterRequire struct {
	swr services.ServiceWaterRequire
}

func NewHandlersWaterRequire(swr services.ServiceWaterRequire) HandlersWaterRequire {
	return &hWaterRequire{
		swr: swr,
	}
}

func (h *hWaterRequire) WaterRequireDataGet(c *fiber.Ctx) error {
	var result interface{}
	var err error
	Action := c.Query("type")

	if !myfunc.CheckDupStringInList(Action, myvar.ParamWater) {
		return c.Status(400).JSON(mymodels.SetResponse("type not match."))
	}

	switch Action {
	case myvar.DAY:
		result, err = h.swr.WaterRequireDayGetAll()
	case myvar.PLANT:
		result, err = h.swr.WaterRequirePlantGetAll()
	case myvar.INDUSTRY:
		result, err = h.swr.WaterRequireIndustryGetAll()
	case myvar.ANIMAL:
		result, err = h.swr.WaterRequireAnimalGetAll()
	}
	if err != nil {
		return err
	}
	return c.Status(200).JSON(mymodels.SetResponse("success", result))
}

func (h *hWaterRequire) WaterDayCal(c *fiber.Ctx) error {
	var mybody = models.BodyWaterAreaCal{}
	getbody := c.Body()
	if err := json.Unmarshal(getbody, &mybody); err != nil {
		return err
	}
	result, err := h.swr.WaterRequireDayCal(mybody)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(mymodels.SetResponse("success", result))
}

func (h *hWaterRequire) WaterIndustryCal(c *fiber.Ctx) error {
	var mybody = models.BodyWaterIndustryCal{}
	getbody := c.Body()
	if err := json.Unmarshal(getbody, &mybody); err != nil {
		return err
	}
	result, err := h.swr.WaterRequireIndustryCal(mybody)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(mymodels.SetResponse("success", result))
}

func (h *hWaterRequire) WaterPlantCal(c *fiber.Ctx) error {
	var mybody = models.BodyWaterPlantCal{}
	getbody := c.Body()
	if err := json.Unmarshal(getbody, &mybody); err != nil {
		return err
	}
	result, err := h.swr.WaterRequirePlantCal(mybody)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(mymodels.SetResponse("success", result))
}

func (h *hWaterRequire) WaterAnimalCal(c *fiber.Ctx) error {
	var mybody = models.BodyWaterAnimalCal{}
	if err := json.Unmarshal(c.Body(), &mybody); err != nil {
		return err
	}
	result, failMSG, err := h.swr.WaterRequireAnimalCal(mybody)
	if err != nil {
		return err
	}
	if failMSG != "" {
		return c.Status(400).JSON(mymodels.SetResponse(failMSG))
	}
	return c.Status(200).JSON(mymodels.SetResponse("success", result))
}
