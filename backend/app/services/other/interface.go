package service

type Services interface {
	ProvinceGetAll() (result interface{}, err error)
}
