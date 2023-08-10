package repo

import (
	"plan-farm/myconfig/mymodels"
	"plan-farm/pkg/myconnect"
	"plan-farm/pkg/myfunc"
)

// ######################## Area ########################

func WaterRequireAreaGetAll() (result []mymodels.DBWaterRequiryArea, err error) {
	mydb := myconnect.DBInstance()
	err = mydb.Table(
		mymodels.DBWaterRequiryArea.TableName(mymodels.DBWaterRequiryArea{})).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}

func WaterRequireAreaGetByID(areaID int) (result mymodels.DBWaterRequiryArea, err error) {
	mydb := myconnect.DBInstance()
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
	mydb := myconnect.DBInstance()
	err = mydb.Table(
		mymodels.DBWaterRequiryIndustry.TableName(mymodels.DBWaterRequiryIndustry{})).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}

func WaterRequireIndustryGetByID(industryID int) (result mymodels.DBWaterRequiryIndustry, err error) {
	mydb := myconnect.DBInstance()
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
	mydb := myconnect.DBInstance()
	err = mydb.Table(
		mymodels.DBWaterRequirePlant.TableName(mymodels.DBWaterRequirePlant{})).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}

func WaterRequirePlantGetByID(plantID int) (result mymodels.DBWaterRequirePlant, err error) {
	mydb := myconnect.DBInstance()
	err = mydb.Table(
		mymodels.DBWaterRequirePlant.TableName(mymodels.DBWaterRequirePlant{})).
		Where("id=?", plantID).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}

// ######################## Person ########################

func WaterRequirePersonGetAll() (result []mymodels.DBWaterRequirePerson, err error) {
	mydb := myconnect.DBInstance()
	err = mydb.Table(
		mymodels.DBWaterRequirePerson.TableName(mymodels.DBWaterRequirePerson{})).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}

func WaterRequirePersonGetByID(wPersonID int) (result mymodels.DBWaterRequirePerson, err error) {
	mydb := myconnect.DBInstance()
	err = mydb.Table(
		mymodels.DBWaterRequirePerson.TableName(mymodels.DBWaterRequirePerson{})).
		Where("id=?", wPersonID).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}

// #################### ETC ####################
func ProvinceGetAll() (result []mymodels.DBProvince, err error) {
	mydb := myconnect.DBInstance()
	err = mydb.Table(
		mymodels.DBProvince.TableName(mymodels.DBProvince{})).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}
