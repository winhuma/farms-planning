package repo

import "farms-planning/myconfig/models"

type Repo interface {
	ProvinceGetAll() (result []models.DBProvince, err error)
}

type RepoWaterRequire interface {
	WaterRequireAreaGetAll() (result []models.DBWaterRequiryArea, err error)
	WaterRequireDayGetByID(areaID int) (result models.DBWaterRequiryArea, err error)
	WaterRequireIndustryGetAll() (result []models.DBWaterRequiryIndustry, err error)
	WaterRequireIndustryGetByID(industryID int) (result models.DBWaterRequiryIndustry, err error)
	WaterRequirePlantGetAll() (result []models.DBWaterRequirePlant, err error)
	WaterRequirePlantGetByID(plantID int) (result models.DBWaterRequirePlant, err error)
	WaterRequirePlantGetByProvinceID(provinceID int) (result models.DBWaterRequirePlant, err error)
	WaterRequireAnimalGetAll() (result []models.DBWaterRequireAnimal, err error)
	WaterRequireAnimalGetByID(animalID int) (result models.DBWaterRequireAnimal, err error)
}
