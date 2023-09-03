package services

import (
	"encoding/json"
	"fmt"
	"plan-farm/app/repo"
	"plan-farm/myconfig/mymodels"
	"plan-farm/pkg/myfunc"
	"strconv"
)

// ################### Area ####################

func WaterRequireAreaGetAll() (result interface{}, err error) {
	return repo.WaterRequireAreaGetAll()
}

func WaterRequireAreaCal(userData mymodels.BodyWaterAreaCal) (result interface{}, err error) {
	areaData, err := repo.WaterRequireAreaGetByID(userData.AreaID)
	if err != nil {
		return nil, err
	}

	result = FormulaWaterArea(areaData.Value, userData.NumberDay, userData.NumberPerson)
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
	dPlantInProvince, err := repo.WaterRequirePlantGetByID(userData.ProvinceID)
	if err != nil {
		return nil, myfunc.MyErrFormat(err)
	}

	pData := map[string]interface{}{}
	err = json.Unmarshal([]byte(dPlantInProvince.PlantData), &pData)
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

// ################### OTHER ###################
func ProvinceGetAll() (result interface{}, err error) {
	return repo.ProvinceGetAll()
}
