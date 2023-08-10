package handlers

import (
	"encoding/json"
	"plan-farm/app/services"
	"plan-farm/myconfig/mymodels"
	"plan-farm/pkg/models"
	"plan-farm/pkg/myfunc"

	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	return c.Status(200).SendString("hello plan farm")
}

func WaterRequireDataGet(c *fiber.Ctx) error {
	var result interface{}
	var err error
	Action := c.Query("type")

	if !myfunc.CheckDupStringInList(Action, []string{"area", "plant", "industry", "person"}) {
		return c.Status(400).JSON(models.ResponseFail("type not match."))
	}

	switch Action {
	case "area":
		result, err = services.WaterRequireAreaGetAll()
	case "plant":
		result, err = services.WaterRequirePlantGetAll()
	case "industry":
		result, err = services.WaterRequireIndustryGetAll()
	case "person":
		result, err = services.WaterRequirePersonGetAll()
	}
	if err != nil {
		return err
	}
	return c.Status(200).JSON(models.ResponseSuccess("success", result))
}

func WaterAreaCal(c *fiber.Ctx) error {
	var mybody = mymodels.BodyWaterAreaCal{}
	getbody := c.Body()
	err := json.Unmarshal(getbody, &mybody)
	if err != nil {
		return err
	}
	result, err := services.WaterRequireAreaCal(mybody)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(models.ResponseSuccess("success", result))
}

func WaterIndustryCal(c *fiber.Ctx) error {
	var mybody = mymodels.BodyWaterIndustryCal{}
	getbody := c.Body()
	err := json.Unmarshal(getbody, &mybody)
	if err != nil {
		return err
	}
	result, err := services.WaterIndustryCal(mybody)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(models.ResponseSuccess("success", result))
}

func WaterPlantCal(c *fiber.Ctx) error {
	var mybody = mymodels.BodyWaterPlantCal{}
	getbody := c.Body()
	err := json.Unmarshal(getbody, &mybody)
	if err != nil {
		return err
	}
	result, err := services.WaterRequirePlantCal(mybody)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(models.ResponseSuccess("success", result))
}
