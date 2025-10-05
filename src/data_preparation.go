package data_preparation

import (
	"encoding/csv"
	"os"
)

type CityData struct {
	city_id   int64
	latitude  float32
	longitude float64
}

type Error struct {
	msg string
}

func load_from_csv(
	path string,
	city_ind,
	lat_ind int16,
	long_ind int16,
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
		cities_data = append(
			cities_data,
			CityData{
				row[city_ind],
				row[lat_ind],
				row[long_ind],
			},
		)
	}

	return cities_data
}
