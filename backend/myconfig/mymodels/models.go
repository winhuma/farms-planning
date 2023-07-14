package mymodels

type BodyWaterAreaCal struct {
	AreaName     string `json:"area_name"`
	NumberPerson string `json:"number_person"`
	NumberDay    string `json:"number_day"`
}

type BodyWaterIndustryCal struct {
	IndustryType     string  `json:"industry_type"`
	IndustryAreaSize float64 `json:"industry_area_size"`
}

type BodyWaterPlantCal struct {
	Plant    string  `json:"plant"`
	Province string  `json:"province"`
	FarmArea float64 `json:"farm_area"`
}
