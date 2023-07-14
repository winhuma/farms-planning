package handlers

import (
	"encoding/json"
	"plan-farm/app/services"
	"plan-farm/myconfig/mymodels"
	"plan-farm/myconfig/myvar"

	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	return c.Status(200).SendString("hello plan farm")
}

func WaterRequireDataGet(c *fiber.Ctx) error {
	var collectionName string
	Action := c.Params("action")
	switch Action {
	case "area":
		collectionName = myvar.CollectionWaterArea
	case "plant":
		collectionName = myvar.CollectionWaterPlant
	case "industry":
		collectionName = myvar.CollectionWaterIndustry
	}
	mydata, err := services.FirebaseGetByCollection(collectionName)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(map[string]interface{}{"message": "success", "data": mydata})
}

func WaterAreaCal(c *fiber.Ctx) error {
	var mybody = mymodels.BodyWaterAreaCal{}
	getbody := c.Body()
	err := json.Unmarshal(getbody, &mybody)
	if err != nil {
		return err
	}
	result, err := services.WaterAreaCal(mybody)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(map[string]interface{}{"message": "success", "data": result})
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
	return c.Status(200).JSON(map[string]interface{}{"message": "success", "data": result})
}

func WaterPlantCal(c *fiber.Ctx) error {
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
	return c.Status(200).JSON(map[string]interface{}{"message": "success", "data": result})
}
