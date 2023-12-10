package repo

import "farms-planning/myconfig/models"

type Repo interface {
	ProvinceGetAll() (result []models.DBProvince, err error)
	ProvinceWeatherGetAll() (result []models.DBProvinceWeather, err error)
	ProvinceWeatherGetByProvince(provinceID int) (result models.DBProvinceWeather, err error)
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

type RepoWaterDataValue interface {
	WaterEvaporationRateGetAll() ([]models.DBWaterEvaporationRate, error)
	WaterEvaporationRateByID(eid int) (models.DBWaterEvaporationRate, error)
	WaterEvaporationRateByProvinceID(provinceID int) (models.DBWaterEvaporationRate, error)

	WaterRunOffRateGetAll() ([]models.DBWaterRunOffRate, error)
	WaterRunOffRateByID(roID int) (models.DBWaterRunOffRate, error)
	WaterRunOffRateByProvinceID(provinceID int) (models.DBWaterRunOffRate, error)

	WaterAverageRainPerYearGetAll() ([]models.DBWaterAverageRainPerYear, error)
	WaterAverageRainPerYearByID(arID int) (models.DBWaterAverageRainPerYear, error)
	WaterAverageRainPerYearByProvince(provinceID int) (models.DBWaterAverageRainPerYear, error)

	WaterFloodPeakRateGetAll() ([]models.DBWaterFloodPeakRate, error)
	WaterFloodPeakRateByID(fprID int) (models.DBWaterFloodPeakRate, error)
	WaterFloodPeakRateByProvinceID(provinceID int) (models.DBWaterFloodPeakRate, error)

	WaterLeakageRateGetAll() ([]models.DBWaterLeakageRate, error)
	WaterLeakageRateByID(lrID int) (models.DBWaterLeakageRate, error)
	WaterLeakageRateByProvinceID(provinceID int) (models.DBWaterLeakageRate, error)
}
