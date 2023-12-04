package service

import "context"

type Services interface {
	ProvinceGetAll(ctx context.Context) (result interface{}, err error)
	GeneratePDF(ctx context.Context, userData map[string]interface{}) ([]byte, error)
}
