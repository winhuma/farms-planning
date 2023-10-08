package models

type (
	BodyWaterAreaCal struct {
		AreaID       int `json:"area_id"`
		NumberPerson int `json:"number_person"`
		NumberDay    int `json:"number_day"`
	}
	BodyWaterIndustryCal struct {
		IndustryID       int     `json:"industry_id"`
		IndustryAreaSize float64 `json:"industry_area_size"`
		NumberDay        int     `json:"number_day"`
	}
	BodyWaterPlantCal struct {
		PlantValue float64 `json:"plant_value"`
		FarmArea   float64 `json:"farm_area"`
	}
	BodyWaterAnimalCal struct {
		AnimalID  int `json:"animal_id"`
		NumAnimal int `json:"number_animal"`
		NumDay    int `json:"number_day"`
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
	DBWaterRequireAnimal struct {
		ID         int     `json:"id"`
		Unit       string  `json:"unit"`
		Value      float64 `json:"value"`
		AnimalName string  `json:"animal_name"`
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

func (DBWaterRequireAnimal) TableName() string {
	return "water_require_animal"
}

func (DBProvince) TableName() string {
	return "province"
}
