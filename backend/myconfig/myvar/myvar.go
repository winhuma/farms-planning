package myvar

import "plan-farm/myconfig/mymodels"

const (
	CollectionWaterArea     = "water_require_day"
	CollectionWaterAnimal   = "water_require_animal"
	CollectionWaterPlant    = "water_require_plant"
	CollectionWaterIndustry = "water_require_industry"
	CollectionWaterPerson   = "water_require_person"
)

const (
	AREA     = "area"
	PLANT    = "plant"
	INDUSTRY = "industry"
	PERSON   = "person"
	ANIMAL   = "animal"
)

var ParamWater = []string{AREA, PLANT, INDUSTRY, PERSON, ANIMAL}

var AppObj = mymodels.AppObject{}
