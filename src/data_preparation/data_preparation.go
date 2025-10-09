package data_preparation

import (
	"encoding/csv"
	"os"
	"strconv"
)

type CityData struct {
	city_id   int64
	latitude  float64
	longitude float64
}

func LoadFromCsv(
	path string,
) []CityData {
	// Load city data from csv
	var cities_data []CityData = []CityData{}

	file, err := os.Open(path)

	if err != nil {
		return cities_data
	}

	csv_reader := csv.NewReader(file)

	records, err := csv_reader.ReadAll()
	if err != nil {
		return cities_data
	}
	for _, row := range records {
		if len(row) < 7 {
			continue
		}
		city_id, _ := strconv.ParseInt(row[0], 10, 64)
		latitude, _ := strconv.ParseFloat(row[5], 64)
		longitude, _ := strconv.ParseFloat(row[6], 64)
		cities_data = append(
			cities_data,
			CityData{
				city_id,
				latitude,
				longitude,
			},
		)
	}

	return cities_data
}
