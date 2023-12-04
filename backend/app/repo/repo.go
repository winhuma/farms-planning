package repo

import (
	"farms-planning/myconfig/models"
	"farms-planning/pkg/myfunc"

	"farms-planning/pkg/runapp/fiber/logger"

	"gorm.io/gorm"
)

type repo struct {
	log logger.LoggerFiber
	DB  *gorm.DB
}

func NewRepositories(log logger.LoggerFiber, db *gorm.DB) Repo {
	return &repo{
		log: log,
		DB:  db,
	}
}

// #################### ETC ####################
func (r *repo) ProvinceGetAll() (result []models.DBProvince, err error) {
	mydb := r.DB
	err = mydb.Table(
		models.DBProvince.TableName(models.DBProvince{})).
		Scan(&result).Error
	if err != nil {
		return result, myfunc.MyErrFormat(err)
	}
	return result, nil
}
