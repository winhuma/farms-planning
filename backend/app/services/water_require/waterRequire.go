package services

import (
	"encoding/json"
	"farms-planning/app/repo"
	"farms-planning/myconfig/models"
	"farms-planning/pkg/myfunc"
	"farms-planning/pkg/runapp/fiber/logger"
	"strings"
)

type serviceWaterRequire struct {
	log logger.LoggerFiber
	r   repo.Repo
	rwr repo.RepoWaterRequire
}

func NewServiceWaterRequire(log logger.LoggerFiber, r repo.Repo, rwr repo.RepoWaterRequire) ServiceWaterRequire {
	return &serviceWaterRequire{
		log: log,
		r:   r,
		rwr: rwr,
	}
}

// ################### Day ####################

func (s *serviceWaterRequire) WaterRequireDayGetAll() (result interface{}, err error) {
	return s.rwr.WaterRequireAreaGetAll()
}

func (s *serviceWaterRequire) WaterRequireDayCal(userData models.BodyWaterAreaCal) (result interface{}, err error) {
	areaData, err := s.rwr.WaterRequireDayGetByID(userData.AreaID)
	if err != nil {
		return nil, err
	}

	result = FormulaWaterDay(areaData.Value, float64(userData.NumberDay), float64(userData.NumberPerson))
	return result, nil
}

// ################### Industry ####################

func (s *serviceWaterRequire) WaterRequireIndustryGetAll() (result interface{}, err error) {
	return s.rwr.WaterRequireIndustryGetAll()
}

func (s *serviceWaterRequire) WaterRequireIndustryCal(userData models.BodyWaterIndustryCal) (result interface{}, err error) {
	industryData, err := s.rwr.WaterRequireIndustryGetByID(userData.IndustryID)
	if err != nil {
		return nil, err
	}

	result = FormulaWaterIndustry(industryData.Value, userData.IndustryAreaSize, float64(userData.NumberDay))

	return result, nil
}

// ################### Plant ####################

func (s *serviceWaterRequire) WaterRequirePlantGetAll() (interface{}, error) {
	rawData, err := s.rwr.WaterRequirePlantGetAll()
	if err != nil {
		return nil, err
	}
	allProvince, err := s.r.ProvinceGetAll()
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

		for kName, _ := range pData {
			if strings.Contains(kName, "จังหวัด") {
				delete(pData, kName)
			} else if strings.Contains(kName, "ลำดับ") {
				delete(pData, kName)
			} else if strings.Contains(kName, "หน่วย") {
				delete(pData, kName)
			}
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

func (s *serviceWaterRequire) WaterRequirePlantCal(userData models.BodyWaterPlantCal) (interface{}, error) {
	result := FormulaWaterPlant(userData.PlantValue, userData.FarmArea)
	return result, nil
}

// ################### ANIMAL ####################

func (s *serviceWaterRequire) WaterRequireAnimalGetAll() (result interface{}, err error) {
	return s.rwr.WaterRequireAnimalGetAll()
}

func (s *serviceWaterRequire) WaterRequireAnimalCal(userData models.BodyWaterAnimalCal) (result float64, failMSG string, err error) {
	dAnimal, err := s.rwr.WaterRequireAnimalGetByID(userData.AnimalID)
	if err != nil {
		return result, failMSG, myfunc.MyErrFormat(err)
	}

	// if dAnimal.ID == 0 {
	// 	failMSG = "not found this animal id."
	// 	return result, failMSG, nil
	// }

	result = FormulaWaterAnimal(dAnimal.Value, float64(userData.NumAnimal), float64(userData.NumDay))
	return result, failMSG, nil
}
