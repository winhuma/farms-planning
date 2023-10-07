package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/tealeg/xlsx"
)

type ProvinceData struct {
	Result []Privince `json:"result"`
}

type Privince struct {
	ID           int    `json:"id"`
	ProvinceName string `json:"province_name"`
}

func main() {

	fileProvince := "cmd/migrate/input/province.json"
	fileInput := "cmd/migrate/input/hydrograph.xlsx"
	fileOutput := "cmd/migrate/output/hydrograph.json"

	pData := ReadProvinceFromJSON(fileProvince)
	// targetData := ReadDataFromCSV(fileInput)
	targetData := ReadDataFromXlsx(fileInput)

	var mapProvince = map[string]int{}
	for _, v := range pData.Result {
		mapProvince[v.ProvinceName] = v.ID
	}

	var result = []map[string]interface{}{}
	for i, v := range targetData {

		if i == 0 {
			continue
		}

		var found bool
		for _, p := range pData.Result {
			if strings.Contains(p.ProvinceName, v[0]) {
				found = true
				result = append(result, map[string]interface{}{
					"id":            p.ID,
					"province_name": p.ProvinceName,
					"value":         v[1],
					"unit":          "อัตราน้ำท่า พ.ค.-ต.ค.®ม.3/กม.2",
				})
			}
		}
		if !found {
			result = append(result, map[string]interface{}{
				"id":            0,
				"province_name": v[0],
				"value":         v[1],
				"unit":          "อัตราน้ำท่า พ.ค.-ต.ค.®ม.3/กม.2",
			})
		}
	}
	WriteJSONToFile(fileOutput, result)
}

func ReadProvinceFromJSON(filename string) ProvinceData {
	var pData ProvinceData
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Panicf("Error reading file: %v", err)
	}

	if err := json.Unmarshal(data, &pData); err != nil {
		log.Panicf("Error decoding JSON: %v", err)
	}

	return pData
}

func ReadDataFromCSV(filename string) [][]string {
	// Open the CSV file for reading
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return [][]string{}
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read all records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return [][]string{}
	}
	return records
}

// writeJSONToFile writes JSON data to a file.
func WriteJSONToFile(filename string, rawData interface{}) {
	jsonData, err := json.Marshal(rawData)
	if err != nil {
		log.Panicf("error marshaling JSON: %v", err)
	}

	file, err := os.Create(filename)
	if err != nil {
		log.Panicf("Create: %v", err)
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		log.Panicf("write: %v", err)
	}
}

func ReadDataFromXlsx(excelFileName string) [][]string {
	var result = [][]string{}
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.Fatalf("Error opening XLSX file: %v", err)
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			result = append(result, []string{
				row.Cells[0].String(),
				row.Cells[1].String(),
			})
		}
	}
	return result
}
