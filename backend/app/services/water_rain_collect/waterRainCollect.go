package services

import (
	"context"
	"farms-planning/app/repo"
	"farms-planning/pkg/runapp/fiber/logger"
)

type waterRainCol struct {
	log           logger.LoggerFiber
	repoDataValue repo.RepoWaterDataValue
}

func NewWaterRainCol(log logger.LoggerFiber, repoDataValue repo.RepoWaterDataValue) ServiceWaterRainCollect {
	return &waterRainCol{
		log:           log,
		repoDataValue: repoDataValue,
	}
}

func (s *waterRainCol) CalSurfaceAreaAtReservoirLevel(ctx context.Context) (float64, error) {
	var result float64
	return result, nil
}

func (s *waterRainCol) CalAreaTerraceReceivesRainWater(ctx context.Context) (float64, error) {
	var result float64
	return result, nil
}

func (s *waterRainCol) CalRunOffWater(ctx context.Context) (float64, error) {
	var result float64
	return result, nil
}

func (s *waterRainCol) CalWaterLostFromEvaporation(ctx context.Context) (float64, error) {
	var result float64
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
