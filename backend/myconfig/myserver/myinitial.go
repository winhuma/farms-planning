package myserver

import (
	"plan-farm/myconfig/mymodels"

	"gorm.io/gorm"
)

func MiGrateAllDB(db *gorm.DB) {
	err := db.AutoMigrate(
		&mymodels.DBProvince{},
		&mymodels.DBWaterRequiryArea{},
		&mymodels.DBWaterRequirePerson{},
		&mymodels.DBWaterRequirePlant{},
		&mymodels.DBWaterRequiryIndustry{},
		&mymodels.DBWaterRequireAnimal{},
	)
	if err != nil {
		panic("MiGrateAllDB : " + err.Error())
	}
}
