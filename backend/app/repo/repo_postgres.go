package repo

import (
	"plan-farm/myconfig/mymodels"
	"plan-farm/myconfig/myvar"
	"plan-farm/pkg/myfunc"
)

// ######################## Area ########################

func WaterRequireAreaGetAll() (result []mymodels.DBWaterRequiryArea, err error) {
	mydb := myvar.AppObj.DB
	err = mydb.Table(
		mymodels.DBWaterRequiryArea.TableName(mymodels.DBWaterRequiryArea{})).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}

func WaterRequireDayGetByID(areaID int) (result mymodels.DBWaterRequiryArea, err error) {
	mydb := myvar.AppObj.DB
	err = mydb.Table(
		mymodels.DBWaterRequiryArea.TableName(mymodels.DBWaterRequiryArea{})).
		Where("id=?", areaID).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}

// ######################## Industry ########################

func WaterRequireIndustryGetAll() (result []mymodels.DBWaterRequiryIndustry, err error) {
	mydb := myvar.AppObj.DB
	err = mydb.Table(
		mymodels.DBWaterRequiryIndustry.TableName(mymodels.DBWaterRequiryIndustry{})).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}

func WaterRequireIndustryGetByID(industryID int) (result mymodels.DBWaterRequiryIndustry, err error) {
	mydb := myvar.AppObj.DB
	err = mydb.Table(
		mymodels.DBWaterRequiryIndustry.TableName(mymodels.DBWaterRequiryIndustry{})).
		Where("id=?", industryID).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}

// ######################## Plant ########################

func WaterRequirePlantGetAll() (result []mymodels.DBWaterRequirePlant, err error) {
	mydb := myvar.AppObj.DB
	err = mydb.Table(
		mymodels.DBWaterRequirePlant.TableName(mymodels.DBWaterRequirePlant{})).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}

func WaterRequirePlantGetByID(plantID int) (result mymodels.DBWaterRequirePlant, err error) {
	mydb := myvar.AppObj.DB
	err = mydb.Table(
		mymodels.DBWaterRequirePlant.TableName(mymodels.DBWaterRequirePlant{})).
		Where("id=?", plantID).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}

func WaterRequirePlantGetByProvinceID(provinceID int) (result mymodels.DBWaterRequirePlant, err error) {
	mydb := myvar.AppObj.DB
	err = mydb.Table(
		mymodels.DBWaterRequirePlant.TableName(mymodels.DBWaterRequirePlant{})).
		Where("province_id=?", provinceID).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}

// ######################## Person ########################

func WaterRequirePersonGetAll() (result []mymodels.DBWaterRequirePerson, err error) {
	mydb := myvar.AppObj.DB
	err = mydb.Table(
		mymodels.DBWaterRequirePerson.TableName(mymodels.DBWaterRequirePerson{})).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}

func WaterRequirePersonGetByID(wPersonID int) (result mymodels.DBWaterRequirePerson, err error) {
	mydb := myvar.AppObj.DB
	err = mydb.Table(
		mymodels.DBWaterRequirePerson.TableName(mymodels.DBWaterRequirePerson{})).
		Where("id=?", wPersonID).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}

// #################### ANIMAL ####################
func WaterRequireAnimalGetAll() (result []mymodels.DBWaterRequireAnimal, err error) {
	mydb := myvar.AppObj.DB
	err = mydb.Table(
		mymodels.DBWaterRequireAnimal.TableName(mymodels.DBWaterRequireAnimal{})).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}

func WaterRequireAnimalGetByID(animalID int) (result mymodels.DBWaterRequireAnimal, err error) {
	mydb := myvar.AppObj.DB
	err = mydb.Table(
		mymodels.DBWaterRequireAnimal.TableName(mymodels.DBWaterRequireAnimal{})).
		Where("id=?", animalID).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}

// #################### ETC ####################
func ProvinceGetAll() (result []mymodels.DBProvince, err error) {
	mydb := myvar.AppObj.DB
	err = mydb.Table(
		mymodels.DBProvince.TableName(mymodels.DBProvince{})).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}
