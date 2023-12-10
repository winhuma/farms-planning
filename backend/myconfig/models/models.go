package models

type (
	DBProvince struct {
		ID           int    `json:"id"`
		ProvinceName string `json:"province_name"`
	}
	DBProvinceWeather struct {
		ID               int     `json:"id"`
		ProvinceID       int     `json:"province_id"`
		RainAvgPerYear   float64 `json:"rain_avg_per_year"`
		EvaporationRate  float64 `json:"evaporation_rate"`
		LeakageRate      float64 `json:"leakage_rate"`
		RunoffRate       float64 `json:"runoff_rate"`
		HighestFloodRate float64 `json:"highest_flood_rate"`
	}
)

func (DBProvince) TableName() string {
	return "province"
}

func (DBProvinceWeather) TableName() string {
	return "province_weather"
}
