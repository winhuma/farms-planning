package repo

import (
	"farms-planning/myconfig/models"
	"farms-planning/pkg/myfunc"

	"gorm.io/gorm"
)

type repo struct {
	DB *gorm.DB
}

func NewRepositories(db *gorm.DB) Repo {
	return &repo{
		DB: db,
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
