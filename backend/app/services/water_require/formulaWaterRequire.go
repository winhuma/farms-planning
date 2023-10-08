package services

import (
	"github.com/rs/zerolog/log"
)

func FormulaWaterDay(areaValue float64, numberDay float64, numberPerson float64) float64 {
	log.Info().Msgf("[FormulaWaterDay]: %f * %f * %f", areaValue, numberDay, numberPerson)
	return areaValue * numberDay * numberPerson
}

func FormulaWaterIndustry(industryValue float64, IndustryAreaSize float64, numberDay float64) float64 {
	log.Info().Msgf("[FormulaWaterIndustry]: %f * %f * %f", industryValue, IndustryAreaSize, numberDay)
	return industryValue * IndustryAreaSize * numberDay
}

func FormulaWaterPlant(plantValue float64, areaPlant float64) float64 {
	// ETc=Kc*ETo
	log.Info().Msgf("[FormulaWaterPlant]: %f * %f", plantValue, areaPlant)
	return plantValue * areaPlant
}

func FormulaWaterAnimal(animalWaterValue float64, numAnimal float64, numDay float64) float64 {
	// ปริมาณความต้องการน้ําปศุสัตว์(ลบ.ม.) = สปส.การใช้น้ําของสัตว์(ลบ.ม./ตัว/วัน)*จํานวนปศุสัตว์(ตัว)*จํานวนวันที่เก็บน้ําไว้ให้ปศุสัตว์(วัน)
	log.Info().Msgf("[FormulaWaterAnimal]: %f * %f * %f", animalWaterValue, numAnimal, numDay)
	return animalWaterValue * numAnimal * numDay
}
