package services

import (
	"context"
	"farms-planning/app/repo"
	"farms-planning/myconfig/models"
	"farms-planning/pkg/runapp/fiber/logger"
)

type waterRainCol struct {
	log           logger.LoggerFiber
	repo          repo.Repo
	repoDataValue repo.RepoWaterDataValue
}

func NewWaterRainCol(log logger.LoggerFiber, repo repo.Repo, repoDataValue repo.RepoWaterDataValue) ServiceWaterRainCollect {
	return &waterRainCol{
		log:           log,
		repo:          repo,
		repoDataValue: repoDataValue,
	}
}

func (s *waterRainCol) CalSurfaceAreaAtReservoirLevel(ctx context.Context, bodyData models.BodyCalSurfaceArea) (float64, error) {
	result := FormulaSurfaceAreaAtReservoirLevel(bodyData.AreaPool, bodyData.PercentArea)
	return result, nil
}

func (s *waterRainCol) CalAreaTerraceReceivesRainWater(ctx context.Context, bodyData models.BodyCalAreaReceivesRainwater) (float64, error) {
	result := FormulaAreaTerraceReceivesRainWater(bodyData.SurfaceArea)
	return result, nil
}

func (s *waterRainCol) GetAverageRunoffPerYearAll(ctx context.Context) (interface{}, error) {
	weatherValue, err := s.repo.ProvinceWeatherGetAll()
	if err != nil {
		return weatherValue, err
	}
	return weatherValue, nil
}

func (s *waterRainCol) GetAverageRunoffPerYearByProvinceID(ctx context.Context, ProvinceID int) (interface{}, error) {
	weatherValue, err := s.repo.ProvinceWeatherGetByProvince(ProvinceID)
	if err != nil {
		return weatherValue, err
	}
	return weatherValue, nil
}

func (s *waterRainCol) CalAverageRunoffPerYear(ctx context.Context, bodyData models.BodyCalAverageRunoffPerYear) (float64, error) {
	var result float64
	weatherValue, err := s.repo.ProvinceWeatherGetByProvince(bodyData.ProvinceID)
	if err != nil {
		return result, err
	}
	result = FormulaRunOffWater(weatherValue.RunoffRate, bodyData.AreaReceivesRainwater)
	return result, nil
}

func (s *waterRainCol) CalWaterLostFromEvaporation(ctx context.Context, bodyData models.BodyCalWaterLostFromEvaporation) (float64, error) {
	var result float64
	weatherValue, err := s.repo.ProvinceWeatherGetByProvince(bodyData.ProvinceID)
	if err != nil {
		return result, err
	}
	result = FormulaWaterLostFromEvaporation(weatherValue.EvaporationRate, bodyData.SurfaceArea)
	return result, nil
}

func (s *waterRainCol) CalWaterLostFromLeakage(ctx context.Context) (float64, error) {
	var result float64
	return result, nil
}

func (s *waterRainCol) CalWaterCopacity(ctx context.Context) (float64, error) {
	var result float64
	return result, nil
}
