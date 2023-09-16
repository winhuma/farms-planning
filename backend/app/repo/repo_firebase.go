package repo

import (
	"context"
	"plan-farm/myconfig/myvar"
	"plan-farm/pkg/myconnect"
	"plan-farm/pkg/myfunc"

	"google.golang.org/api/iterator"
)

func FirebaseGetByCollection(CollectionName string) ([]map[string]interface{}, error) {
	var mydata = []map[string]interface{}{}
	fs := myconnect.FirebaseGetDB()
	iter := fs.Collection(CollectionName).Documents(context.TODO())
	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, myfunc.MyErrFormat(err)
		}
		mydata = append(mydata, doc.Data())
	}
	return mydata, nil
}

func FBGetWaterRequireDay() ([]map[string]interface{}, error) {
	areaData, err := FirebaseGetByCollection(myvar.CollectionWaterDay)
	if err != nil {
		return nil, err
	}
	return areaData, nil
}

func FBGetWaterRequireIndustry() ([]map[string]interface{}, error) {
	industryData, err := FirebaseGetByCollection(myvar.CollectionWaterIndustry)
	if err != nil {
		return nil, err
	}
	return industryData, nil
}
