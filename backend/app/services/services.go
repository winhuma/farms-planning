package services

import (
	"encoding/json"
	"fmt"
	"plan-farm/app/repo"
	"plan-farm/myconfig/mymodels"
	"plan-farm/pkg/myfunc"
	"strconv"
)

// ################### Day ####################

func WaterRequireDayGetAll() (result interface{}, err error) {
	return repo.WaterRequireAreaGetAll()
}

func WaterRequireDayCal(userData mymodels.BodyWaterAreaCal) (result interface{}, err error) {
	areaData, err := repo.WaterRequireDayGetByID(userData.AreaID)
	if err != nil {
		return nil, err
	}

	result = FormulaWaterDay(areaData.Value, float64(userData.NumberDay), float64(userData.NumberPerson))
	return result, nil
}

// ################### Industry ####################

func WaterRequireIndustryGetAll() (result interface{}, err error) {
	return repo.WaterRequireIndustryGetAll()
}

func WaterIndustryCal(userData mymodels.BodyWaterIndustryCal) (result interface{}, err error) {
	industryData, err := repo.WaterRequireIndustryGetByID(userData.IndustryID)
	if err != nil {
		return nil, err
	}

	result = FormulaWaterIndustry(industryData.Value, userData.IndustryAreaSize)

	return result, nil
}

// ################### Plant ####################

func WaterRequirePlantGetAll() (interface{}, error) {
	rawData, err := repo.WaterRequirePlantGetAll()
	if err != nil {
		return nil, err
	}
	allProvince, err := repo.ProvinceGetAll()
	if err != nil {
		return nil, err
	}

	var mapProvince = map[int]string{}
	for _, province := range allProvince {
		mapProvince[province.ID] = province.ProvinceName
	}
	var result = []map[string]interface{}{}
	for _, plant := range rawData {
		pData := map[string]interface{}{}
		err := json.Unmarshal([]byte(plant.PlantData), &pData)
		if err != nil {
			return nil, myfunc.MyErrFormat(err)
		}

		var re = map[string]interface{}{
			"id":            plant.ID,
			"province_id":   plant.ProvinceID,
			"province_name": mapProvince[plant.ProvinceID],
			"data":          pData,
		}
		result = append(result, re)
	}
	return result, err
}

func WaterRequirePlantCal(userData mymodels.BodyWaterPlantCal) (interface{}, error) {
	dPlantByProvince, err := repo.WaterRequirePlantGetByProvinceID(userData.ProvinceID)
	if err != nil {
		return nil, myfunc.MyErrFormat(err)
	}

	pData := map[string]interface{}{}
	err = json.Unmarshal([]byte(dPlantByProvince.PlantData), &pData)
	if err != nil {
		return nil, err
	}

	plantValue, err := strconv.ParseFloat(fmt.Sprint(pData[userData.PlantName]), 64)
	if err != nil {
		return nil, err
	}

	result := FormulaWaterPlant(plantValue, userData.FarmArea)
	return result, nil
}

// ################### Person ####################
func WaterRequirePersonGetAll() (result interface{}, err error) {
	return repo.WaterRequirePersonGetAll()
}

// ################### ANIMAL ####################

func WaterRequireAnimalGetAll() (result interface{}, err error) {
	return repo.WaterRequireAnimalGetAll()
}

func WaterRequireAnimalCal(userData mymodels.BodyWaterAnimalCal) (result float64, failMSG string, err error) {
	dAnimal, err := repo.WaterRequireAnimalGetByID(userData.AnimalID)
	if err != nil {
		return result, failMSG, myfunc.MyErrFormat(err)
	}

	if dAnimal.ID == 0 {
		failMSG = "not found this animal id."
		return result, failMSG, nil
	}

	result = FormulaWaterAnimal(dAnimal.Value, float64(userData.NumAnimal), float64(userData.NumDay))
	return result, failMSG, nil
}

// ################### OTHER ###################
func ProvinceGetAll() (result interface{}, err error) {
	return repo.ProvinceGetAll()
}
