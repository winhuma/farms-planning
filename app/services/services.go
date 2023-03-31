package services

import (
	"context"
	"fmt"
	"plan-farm/libs/connection"
	"plan-farm/libs/function"
	"plan-farm/myconfig/mymodels"
	"plan-farm/myconfig/myvar"
	"strconv"

	"google.golang.org/api/iterator"
)

func FirebaseGetByCollection(CollectionName string) ([]map[string]interface{}, error) {
	var mydata = []map[string]interface{}{}
	fs := connection.FirebaseGetDB()
	iter := fs.Collection(CollectionName).Documents(context.TODO())
	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, function.MyErrFormat(err)
		}
		mydata = append(mydata, doc.Data())
	}
	return mydata, nil
}

// ################### CAL ####################

func WaterAreaCal(userData mymodels.BodyWaterAreaCal) (interface{}, error) {
	areaData, err := FirebaseGetByCollection(myvar.CollectionWaterArea)
	if err != nil {
		return nil, err
	}
	var areaValue float64
	for _, v := range areaData {
		var checkArea bool
		for _, va := range v {
			if va == userData.AreaName {
				checkArea = true
			}
		}
		if checkArea {
			areaValue, err = strconv.ParseFloat(fmt.Sprint(v["ปริมาณ"]), 64)
			if err != nil {
				return nil, function.MyErrFormat(err)
			}
			break
		}
	}
	nDay, err := strconv.ParseFloat(userData.NumberDay, 64)
	if err != nil {
		return nil, function.MyErrFormat(err)
	}
	nPer, err := strconv.ParseFloat(userData.NumberPerson, 64)
	if err != nil {
		return nil, function.MyErrFormat(err)
	}
	result := areaValue * nDay * nPer
	return result, nil
}

func WaterIndustryCal(userData mymodels.BodyWaterIndustryCal) (interface{}, error) {
	industryData, err := FirebaseGetByCollection(myvar.CollectionWaterIndustry)
	if err != nil {
		return nil, err
	}

	var industryValue float64
	for _, v := range industryData {
		var checkFound bool
		for _, va := range v {
			if va == userData.IndustryType {
				checkFound = true
			}
		}
		if checkFound {
			industryValue, err = strconv.ParseFloat(fmt.Sprint(v["ปริมาณ"]), 64)
			if err != nil {
				return nil, function.MyErrFormat(err)
			}
			break
		}
	}
	result := industryValue * userData.IndustryAreaSize

	return result, nil
}

func WaterPlantCal(userData mymodels.BodyWaterPlantCal) (interface{}, error) {
	plantData, err := FirebaseGetByCollection(myvar.CollectionWaterIndustry)
	if err != nil {
		return nil, err
	}
	_ = plantData
	// for _, v := range plantData {
	// 	for _, va := range v {

	// 	}
	// }
	result := 0
	return result, nil
}
