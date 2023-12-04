package repo

import (
	"farms-planning/myconfig/models"
	"farms-planning/pkg/myfunc"
	"farms-planning/pkg/runapp/fiber/logger"

	"gorm.io/gorm"
)

type waterDataValue struct {
	log logger.LoggerFiber
	DB  *gorm.DB
}

func NewRepoWaterValue(log logger.LoggerFiber, db *gorm.DB) RepoWaterDataValue {
	return &waterDataValue{
		log: log,
		DB:  db,
	}
}

func (r *waterDataValue) WaterEvaporationRateGetAll() (result []models.DBWaterEvaporationRate, err error) {
	err = r.DB.Table(models.DBWaterEvaporationRate{}.TableName()).Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}
func (r *waterDataValue) WaterEvaporationRateByID(eid int) (result models.DBWaterEvaporationRate, err error) {
	err = r.DB.Table(models.DBWaterEvaporationRate{}.TableName()).
		Where("id=?", eid).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}
func (r *waterDataValue) WaterEvaporationRateByProvinceID(provinceID int) (result models.DBWaterEvaporationRate, err error) {
	err = r.DB.Table(models.DBWaterEvaporationRate{}.TableName()).
		Where("province_id=?", provinceID).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}

func (r *waterDataValue) WaterRunOffRateGetAll() (result []models.DBWaterRunOffRate, err error) {
	err = r.DB.Table(models.DBWaterRunOffRate{}.TableName()).Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}
func (r *waterDataValue) WaterRunOffRateByID(roID int) (result models.DBWaterRunOffRate, err error) {
	err = r.DB.Table(models.DBWaterRunOffRate{}.TableName()).
		Where("id=?", roID).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}
func (r *waterDataValue) WaterRunOffRateByProvinceID(provinceID int) (result models.DBWaterRunOffRate, err error) {
	err = r.DB.Table(models.DBWaterRunOffRate{}.TableName()).
		Where("province_id=?", provinceID).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}

func (r *waterDataValue) WaterAverageRainPerYearGetAll() (result []models.DBWaterAverageRainPerYear, err error) {
	err = r.DB.Table(models.DBWaterAverageRainPerYear{}.TableName()).Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}
func (r *waterDataValue) WaterAverageRainPerYearByID(arID int) (result models.DBWaterAverageRainPerYear, err error) {
	err = r.DB.Table(models.DBWaterAverageRainPerYear{}.TableName()).
		Where("id=?", arID).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}
func (r *waterDataValue) WaterAverageRainPerYearByProvince(provinceID int) (result models.DBWaterAverageRainPerYear, err error) {
	err = r.DB.Table(models.DBWaterAverageRainPerYear{}.TableName()).
		Where("province_id=?", provinceID).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}

func (r *waterDataValue) WaterFloodPeakRateGetAll() (result []models.DBWaterFloodPeakRate, err error) {
	err = r.DB.Table(models.DBWaterFloodPeakRate{}.TableName()).Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}
func (r *waterDataValue) WaterFloodPeakRateByID(fprID int) (result models.DBWaterFloodPeakRate, err error) {
	err = r.DB.Table(models.DBWaterFloodPeakRate{}.TableName()).
		Where("id=?", fprID).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}
func (r *waterDataValue) WaterFloodPeakRateByProvinceID(provinceID int) (result models.DBWaterFloodPeakRate, err error) {
	err = r.DB.Table(models.DBWaterFloodPeakRate{}.TableName()).
		Where("province_id=?", provinceID).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}

func (r *waterDataValue) WaterLeakageRateGetAll() (result []models.DBWaterLeakageRate, err error) {
	err = r.DB.Table(models.DBWaterLeakageRate{}.TableName()).Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}
func (r *waterDataValue) WaterLeakageRateByID(lrID int) (result models.DBWaterLeakageRate, err error) {
	err = r.DB.Table(models.DBWaterLeakageRate{}.TableName()).
		Where("id=?", lrID).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}
func (r *waterDataValue) WaterLeakageRateByProvinceID(provinceID int) (result models.DBWaterLeakageRate, err error) {
	err = r.DB.Table(models.DBWaterLeakageRate{}.TableName()).
		Where("province_id=?", provinceID).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}
