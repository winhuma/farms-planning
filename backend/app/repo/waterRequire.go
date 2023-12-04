package repo

import (
	"farms-planning/myconfig/models"
	"farms-planning/pkg/myfunc"
	"farms-planning/pkg/runapp/fiber/logger"

	"gorm.io/gorm"
)

type waterRequire struct {
	log logger.LoggerFiber
	DB  *gorm.DB
}

func NewRepoWaterRequire(log logger.LoggerFiber, db *gorm.DB) RepoWaterRequire {
	return &waterRequire{
		log: log,
		DB:  db,
	}
}

// ######################## Area ########################

func (r *waterRequire) WaterRequireAreaGetAll() (result []models.DBWaterRequiryArea, err error) {
	mydb := r.DB
	err = mydb.Table(
		models.DBWaterRequiryArea.TableName(models.DBWaterRequiryArea{})).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}

func (r *waterRequire) WaterRequireDayGetByID(areaID int) (result models.DBWaterRequiryArea, err error) {
	mydb := r.DB
	err = mydb.Table(
		models.DBWaterRequiryArea.TableName(models.DBWaterRequiryArea{})).
		Where("id=?", areaID).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}

// ######################## Industry ########################

func (r *waterRequire) WaterRequireIndustryGetAll() (result []models.DBWaterRequiryIndustry, err error) {
	mydb := r.DB
	err = mydb.Table(
		models.DBWaterRequiryIndustry.TableName(models.DBWaterRequiryIndustry{})).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}

func (r *waterRequire) WaterRequireIndustryGetByID(industryID int) (result models.DBWaterRequiryIndustry, err error) {
	mydb := r.DB
	err = mydb.Table(
		models.DBWaterRequiryIndustry.TableName(models.DBWaterRequiryIndustry{})).
		Where("id=?", industryID).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}

// ######################## Plant ########################

func (r *waterRequire) WaterRequirePlantGetAll() (result []models.DBWaterRequirePlant, err error) {
	mydb := r.DB
	err = mydb.Table(
		models.DBWaterRequirePlant.TableName(models.DBWaterRequirePlant{})).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}

func (r *waterRequire) WaterRequirePlantGetByID(plantID int) (result models.DBWaterRequirePlant, err error) {
	mydb := r.DB
	err = mydb.Table(
		models.DBWaterRequirePlant.TableName(models.DBWaterRequirePlant{})).
		Where("id=?", plantID).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}

func (r *waterRequire) WaterRequirePlantGetByProvinceID(provinceID int) (result models.DBWaterRequirePlant, err error) {
	mydb := r.DB
	err = mydb.Table(
		models.DBWaterRequirePlant.TableName(models.DBWaterRequirePlant{})).
		Where("province_id=?", provinceID).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}

// #################### ANIMAL ####################
func (r *waterRequire) WaterRequireAnimalGetAll() (result []models.DBWaterRequireAnimal, err error) {
	mydb := r.DB
	err = mydb.Table(
		models.DBWaterRequireAnimal.TableName(models.DBWaterRequireAnimal{})).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}

func (r *waterRequire) WaterRequireAnimalGetByID(animalID int) (result models.DBWaterRequireAnimal, err error) {
	mydb := r.DB
	err = mydb.Table(
		models.DBWaterRequireAnimal.TableName(models.DBWaterRequireAnimal{})).
		Where("id=?", animalID).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}
