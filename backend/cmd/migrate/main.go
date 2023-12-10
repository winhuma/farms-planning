package main

import (
	"encoding/csv"
	"encoding/json"
	"farms-planning/myconfig/appinit"
	"farms-planning/myconfig/models"
	"farms-planning/myconfig/myvar"
	"farms-planning/pkg/myconnect"
	"fmt"
	"log"
	"os"
	"strconv"
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
	provinceFile := "cmd/migrate/initialSetupData/province.json"
	rawFile := "cmd/migrate/input/i_r_q_e_s.csv"
	saveJsonFile := "cmd/migrate/initialSetupData/i_r_q_e_s.json"

	data, err := os.ReadFile(provinceFile)
	if err != nil {
		log.Panicf("Error reading file: %v", err)
	}

	var province = []Privince{}
	if err := json.Unmarshal(data, &province); err != nil {
		log.Panicf("Error decoding JSON: %v", err)
	}

	csvdata := ReadDataFromCSV(rawFile)

	fmt.Println(province[0].ID, province[0].ProvinceName)
	fmt.Println(csvdata[0])

	var result []models.DBProvinceWeather
	for i, v := range csvdata {
		if i == 0 {
			continue
		}

		var pid = 0
		fmt.Println(v[0])
		for _, p := range province {
			if strings.Contains(v[0], p.ProvinceName) {
				pid = p.ID
				break
			}
		}

		RainAvgPerYear, _ := strconv.ParseFloat(v[1], 64)
		EvaporationRate, _ := strconv.ParseFloat(v[2], 64)
		LeakageRate, _ := strconv.ParseFloat(v[3], 64)
		RunoffRate, _ := strconv.ParseFloat(v[4], 64)
		HighestFloodRate, _ := strconv.ParseFloat(v[4], 64)

		result = append(result, models.DBProvinceWeather{
			ProvinceID:       pid,
			RainAvgPerYear:   RainAvgPerYear,
			EvaporationRate:  EvaporationRate,
			LeakageRate:      LeakageRate,
			RunoffRate:       RunoffRate,
			HighestFloodRate: HighestFloodRate,
		})
	}

	WriteJSONToFile(saveJsonFile, result)

}

func UpdateDataToDB() {

	myvar.SetEnv()
	db := myconnect.NewPostgres(myvar.ENV_DB_CONNECT)
	appinit.MiGrateAllDB(db)

	fileOri := "cmd/migrate/initialSetupData/data_e.json"

	data, err := os.ReadFile(fileOri)
	if err != nil {
		log.Panicf("Error reading file: %v", err)
	}

	var result = []map[string]interface{}{}
	if err := json.Unmarshal(data, &result); err != nil {
		log.Panicf("Error decoding JSON: %v", err)
	}

	for i := range result {

		pid, _ := strconv.Atoi(fmt.Sprint(result[i]["province_id"]))
		dis := fmt.Sprint(result[i]["discription"])
		symbol := fmt.Sprint(result[i]["symbol"])
		unitTH := fmt.Sprint(result[i]["unit_th"])
		unitThFull := fmt.Sprint(result[i]["unit_th_fullname"])
		myvalue, _ := strconv.ParseFloat(fmt.Sprint(result[i]["value"]), 64)

		var addData = models.DBWaterEvaporationRate{
			Discription:   dis,
			ProvinceID:    pid,
			Symbol:        symbol,
			UnitTH:        unitTH,
			UnitFulnameTH: unitThFull,
			Value:         myvalue,
		}
		err := db.Table(models.DBWaterEvaporationRate.TableName(models.DBWaterEvaporationRate{})).
			Create(&addData).Error
		if err != nil {
			fmt.Println(result[i])
			log.Panicf("Error create data: %v", err)
		}
	}
}

func UpdatePrivince() {
	fileOri := "cmd/migrate/initialSetupData/data_e.json"

	data, err := os.ReadFile(fileOri)
	if err != nil {
		log.Panicf("Error reading file: %v", err)
	}

	var result = []map[string]interface{}{}
	if err := json.Unmarshal(data, &result); err != nil {
		log.Panicf("Error decoding JSON: %v", err)
	}

	myvar.SetEnv()
	db := myconnect.NewPostgres(myvar.ENV_DB_CONNECT)
	appinit.MiGrateAllDB(db)
	for i := range result {
		fmt.Println(result[i]["province_id"], result[i]["province_name"])

		myid, _ := strconv.Atoi(fmt.Sprint(result[i]["province_id"]))
		ProvinceName := fmt.Sprint(result[i]["province_name"])
		db.Table(models.DBProvince.TableName(models.DBProvince{})).
			Where("id=?", myid).Update("province_name", ProvinceName)
	}
}

func SetFileFinal() {

	_ = `
	{
        "province_id": 43,
        "province_name": "นนทบุรี ",
        "discription": "อัตราน้ำท่า พ.ค.-ต.ค.",
        "symbol": "",
        "unit": "ม.3/กม.2",
        "unit_fullname": "",
        "value": "1247"
    }
	`

	fileOri := "cmd/migrate/fileOutput/data_h.json"
	fileOut := "cmd/migrate/fileFinal/data_h.json"

	var result = []map[string]interface{}{}
	data, err := os.ReadFile(fileOri)
	if err != nil {
		log.Panicf("Error reading file: %v", err)
	}

	if err := json.Unmarshal(data, &result); err != nil {
		log.Panicf("Error decoding JSON: %v", err)
	}

	var filterMap = []map[string]interface{}{}
	for _, v := range result {
		filterMap = append(filterMap, map[string]interface{}{
			"province_id":   v["province_id"],
			"province_name": v["province_name"],
			"value":         v["value"],
			"discription":   "อัตราน้ำท่า พ.ค.-ต.ค.",
			"symbol":        "",
			"unit":          "ม.3/กม.2",
			"unit_fullname": "",
		})
	}

	WriteJSONToFile(fileOut, filterMap)
}

func SetFileInput() {

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
	data, err := os.ReadFile(filename)
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
