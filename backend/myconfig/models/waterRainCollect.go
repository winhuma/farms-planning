package models

type (
	BodyCalSurfaceArea struct {
		AreaPool    float64 `json:"area_pool"`
		PercentArea float64 `json:"percent_area"`
	}

	BodyGetWeatherByProvinceID struct {
		ProvinceID int `json:"province_id"`
	}

	BodyCalAreaReceivesRainwater struct {
		SurfaceArea float64 `json:"surface_area"`
	}
	BodyCalAverageRunoffPerYear struct {
		ProvinceID            int     `json:"province_id"`
		AreaReceivesRainwater float64 `json:"area_receive_water"`
	}
	BodyCalWaterLostFromEvaporation struct {
		ProvinceID  int     `json:"province_id"`
		SurfaceArea float64 `json:"surface_area"`
	}

	BodyCalWaterCapacity struct {
		WaterRequire        float64 `json:"water_require"`
		LostFromEvaporation float64 `json:"lost_from_evaporation"`
		LostFromLeakage     float64 `json:"lost_from_leakage"`
	}
)

type (
	DBWaterEvaporationRate struct {
		ID            int     `json:"id"`
		Discription   string  `json:"discription"`
		ProvinceID    int     `json:"province_id"`
		Symbol        string  `json:"symbol"`
		UnitTH        string  `json:"unit_th"`
		UnitFulnameTH string  `json:"unit_fullname_th"`
		Value         float64 `json:"value"`
	}

	DBWaterRunOffRate struct {
		ID            int     `json:"id"`
		Discription   string  `json:"discription"`
		ProvinceID    int     `json:"province_id"`
		Symbol        string  `json:"symbol"`
		UnitTH        string  `json:"unit_th"`
		UnitFulnameTH string  `json:"unit_fullname_th"`
		Value         float64 `json:"value"`
	}

	DBWaterAverageRainPerYear struct {
		ID            int     `json:"id"`
		Discription   string  `json:"discription"`
		ProvinceID    int     `json:"province_id"`
		Symbol        string  `json:"symbol"`
		UnitTH        string  `json:"unit_th"`
		UnitFulnameTH string  `json:"unit_fullname_th"`
		Value         float64 `json:"value"`
	}

	DBWaterFloodPeakRate struct {
		ID            int     `json:"id"`
		Discription   string  `json:"discription"`
		ProvinceID    int     `json:"province_id"`
		Symbol        string  `json:"symbol"`
		UnitTH        string  `json:"unit_th"`
		UnitFulnameTH string  `json:"unit_fullname_th"`
		Value         float64 `json:"value"`
	}

	DBWaterLeakageRate struct {
		ID            int     `json:"id"`
		Discription   string  `json:"discription"`
		ProvinceID    int     `json:"province_id"`
		Symbol        string  `json:"symbol"`
		UnitTH        string  `json:"unit_th"`
		UnitFulnameTH string  `json:"unit_fullname_th"`
		Value         float64 `json:"value"`
	}
)

func (DBWaterEvaporationRate) TableName() string {
	return "data_water_evaporation_rate"
}

func (DBWaterRunOffRate) TableName() string {
	return "data_water_run_off_rate"
}

func (DBWaterAverageRainPerYear) TableName() string {
	return "data_water_average_rain_per_year"
}

func (DBWaterFloodPeakRate) TableName() string {
	return "data_water_flood_peak_rate"
}

func (DBWaterLeakageRate) TableName() string {
	return "data_water_leakage_rate"
}
