package services

func FormulaWaterArea(areaValue float64, numberDay float64, numberPerson float64) float64 {
	return areaValue * numberDay * numberPerson
}

func FormulaWaterIndustry(industryValue float64, IndustryAreaSize float64) float64 {
	return industryValue * IndustryAreaSize
}

func FormulaWaterPlant(plantValue float64, areaPlant float64) float64 {
	// ETc=Kc*ETo
	return plantValue * areaPlant
}

func FormulaWaterAnimal(animalWaterValue float64, numAnimal float64, numDay float64) float64 {
	// ปริมาณความต้องการน้ําปศุสัตว์(ลบ.ม.) = สปส.การใช้น้ําของสัตว์(ลบ.ม./ตัว/วัน)*จํานวนปศุสัตว์(ตัว)*จํานวนวันที่เก็บน้ําไว้ให้ปศุสัตว์(วัน)
	return animalWaterValue * numAnimal * numDay
}
