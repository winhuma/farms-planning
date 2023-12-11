package services

func FormulaSurfaceAreaAtReservoirLevel(areaPool float64, percentArea float64) float64 {
	return areaPool * percentArea
}

func FormulaAreaTerraceReceivesRainWater(sufface_area_colect_pool float64) float64 {
	return 10 * sufface_area_colect_pool
}

func FormulaRunOffWater(rainValue float64, AreaTerraceReceivesRainWater float64) float64 {
	return 400 * rainValue * AreaTerraceReceivesRainWater
}

func FormulaWaterLostFromEvaporation(evaporationValue float64, AreaTerraceReceivesRainWater float64) float64 {
	// EVP
	return (evaporationValue * AreaTerraceReceivesRainWater) / 2000
}

func FormulaWaterLostFromLeakage(LeakageValue float64, AreaTerraceReceivesRainWater float64) float64 {
	// SL
	return (LeakageValue * AreaTerraceReceivesRainWater) / 2000
}

func FormulaWaterCopacity(WaterLostFromEvaporation float64, WaterLostFromLeakage float64, requireWaterValue float64) float64 {
	// V
	return WaterLostFromEvaporation + WaterLostFromLeakage + requireWaterValue
}
