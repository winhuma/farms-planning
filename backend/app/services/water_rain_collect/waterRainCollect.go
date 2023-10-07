package services

type waterRainCol struct{}

func NewWaterRainCol() ServiceWaterRainCollect {
	return &waterRainCol{}
}
