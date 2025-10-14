package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"src/src/data_preparation"
)

func main() {
	var input string
	var initialSeq int = 5

	flag.StringVar(&input, "input", "", "Input file to parse")
	flag.IntVar(
		&initialSeq,
		"initialSeq",
		5,
		"Count randomly generate city sequences",
	)

	flag.Parse()

	if input == "" {
		log.Fatalln("Fail input path is missing")
		os.Exit(1)
	}

	if _, err := os.Stat(input); errors.Is(err, os.ErrNotExist) {
		log.Fatalln("Fail doesn't exist")
		os.Exit(1)
	}

	var sequences [][]data_preparation.CityData
	var data []data_preparation.CityData = data_preparation.LoadFromCsv(input)

	for i := 0; i < initialSeq; i++ {
		sequences = append(
			sequences,
			data,
		)
	}
}
