package myvar

import "plan-farm/myconfig/mymodels"

const (
	CollectionWaterDay      = "water_require_day"
	CollectionWaterAnimal   = "water_require_animal"
	CollectionWaterPlant    = "water_require_plant"
	CollectionWaterIndustry = "water_require_industry"
	CollectionWaterPerson   = "water_require_person"
)

const (
	DAY      = "day"
	PLANT    = "plant"
	INDUSTRY = "industry"
	PERSON   = "person"
	ANIMAL   = "animal"
)

var ParamWater = []string{DAY, PLANT, INDUSTRY, PERSON, ANIMAL}

var AppObj = mymodels.AppObject{}
