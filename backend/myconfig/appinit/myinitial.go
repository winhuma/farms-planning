package appinit

import (
	"farms-planning/myconfig/models"
	"farms-planning/pkg/myfunc"

	"gorm.io/gorm"
)

func MiGrateAllDB(db *gorm.DB) {
	err := db.AutoMigrate(

		&models.DBProvince{},
		&models.DBProvinceWeather{},

		// group data water require
		&models.DBWaterRequiryArea{},
		&models.DBWaterRequirePlant{},
		&models.DBWaterRequiryIndustry{},
		&models.DBWaterRequireAnimal{},

		// group data water rate
		&models.DBWaterAverageRainPerYear{},
		&models.DBWaterEvaporationRate{},
		&models.DBWaterFloodPeakRate{},
		&models.DBWaterLeakageRate{},
		&models.DBWaterRunOffRate{},
	)
	if err != nil {
		panic(myfunc.MyErrFormat(err).Error())
	}
}
