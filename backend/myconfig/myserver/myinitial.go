package myserver

import (
	"farms-planning/myconfig/models"
	"farms-planning/pkg/myfunc"

	"gorm.io/gorm"
)

func MiGrateAllDB(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.DBProvince{},
		&models.DBWaterRequiryArea{},
		&models.DBWaterRequirePlant{},
		&models.DBWaterRequiryIndustry{},
		&models.DBWaterRequireAnimal{},
	)
	if err != nil {
		panic(myfunc.MyErrFormat(err).Error())
	}
}
