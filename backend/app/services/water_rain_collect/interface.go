package services

import (
	"context"
	"farms-planning/myconfig/models"
)

type ServiceWaterRainCollect interface {
	CalSurfaceAreaAtReservoirLevel(ctx context.Context, bodyData models.BodyCalSurfaceArea) (float64, error)
	GetAverageRunoffPerYearAll(ctx context.Context) (interface{}, error)
	GetAverageRunoffPerYearByProvinceID(ctx context.Context, ProvinceID int) (interface{}, error)
	CalAreaTerraceReceivesRainWater(ctx context.Context, bodyData models.BodyCalAreaReceivesRainwater) (float64, error)
	CalAverageRunoffPerYear(ctx context.Context, bodyData models.BodyCalAverageRunoffPerYear) (float64, error)
	CalWaterLostFromEvaporation(ctx context.Context, bodyData models.BodyCalWaterLostFromEvaporation) (float64, error)
}
