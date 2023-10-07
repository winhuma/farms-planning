package services

import "farms-planning/myconfig/models"

type ServiceWaterRequire interface {
	WaterRequireDayGetAll() (result interface{}, err error)
	WaterRequireDayCal(userData models.BodyWaterAreaCal) (result interface{}, err error)
	WaterRequireIndustryGetAll() (result interface{}, err error)
	WaterRequireIndustryCal(userData models.BodyWaterIndustryCal) (result interface{}, err error)
	WaterRequirePlantGetAll() (interface{}, error)
	WaterRequirePlantCal(userData models.BodyWaterPlantCal) (interface{}, error)
	WaterRequireAnimalGetAll() (result interface{}, err error)
	WaterRequireAnimalCal(userData models.BodyWaterAnimalCal) (result float64, failMSG string, err error)
}
