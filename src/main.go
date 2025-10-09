package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"src/src/data_preparation"
)

func main() {
	var input string

	flag.StringVar(&input, "input", "", "Input file to parse")
	flag.Parse()

	if input == "" {
		log.Fatalln("Fail input path is missing")
		os.Exit(1)
	}

	if _, err := os.Stat(input); errors.Is(err, os.ErrNotExist) {
		log.Fatalln("Fail doesn't exist")
		os.Exit(1)
	}

	fmt.Println(len(data_preparation.LoadFromCsv(input)))

}
