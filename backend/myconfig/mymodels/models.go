package mymodels

type (
	BodyWaterAreaCal struct {
		AreaID       int     `json:"area_id"`
		NumberPerson float64 `json:"number_person"`
		NumberDay    float64 `json:"number_day"`
	}
	BodyWaterIndustryCal struct {
		IndustryID       int     `json:"industry_id"`
		IndustryAreaSize float64 `json:"industry_area_size"`
	}
	BodyWaterPlantCal struct {
		PlantName  string  `json:"plant_name"`
		ProvinceID int     `json:"province_id"`
		FarmArea   float64 `json:"farm_area"`
	}
)

type (
	WaterRequiryArea struct {
		Unit  string  `json:"unit"`
		Value float64 `json:"value"`
		Area  string  `json:"area"`
	}
	WaterRequiryIndustry struct {
		Value        float64 `json:"value"`
		Unit         string  `json:"unit"`
		IndustryName string  `json:"industry_name"`
	}
	WaterRequirePerson struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
		Area  string  `json:"area"`
	}
	WaterRequirePlant struct {
		ProvinceID int    `json:"province_id"`
		PlantData  string `json:"data"`
	}
)

type (
	DBWaterRequiryArea struct {
		ID    int     `json:"id"`
		Unit  string  `json:"unit"`
		Value float64 `json:"value"`
		Area  string  `json:"area"`
	}
	DBWaterRequiryIndustry struct {
		ID           int     `json:"id"`
		Unit         string  `json:"unit"`
		Value        float64 `json:"value"`
		IndustryName string  `json:"industry_name"`
	}
	DBWaterRequirePerson struct {
		ID    int     `json:"id"`
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
		Area  string  `json:"area"`
	}
	DBWaterRequirePlant struct {
		ID         int    `json:"id"`
		ProvinceID int    `json:"province_id"`
		PlantData  string `json:"data"`
	}
)

type DBProvince struct {
	ID           int    `json:"id"`
	ProvinceName string `json:"province_name"`
}

func (DBWaterRequiryArea) TableName() string {
	return "water_require_area"
}

func (DBWaterRequiryIndustry) TableName() string {
	return "water_require_industry"
}

func (DBWaterRequirePerson) TableName() string {
	return "water_require_person"
}

func (DBWaterRequirePlant) TableName() string {
	return "water_require_plant"
}

func (DBProvince) TableName() string {
	return "province"
}
