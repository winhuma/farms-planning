package models

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
