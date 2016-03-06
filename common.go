package main

import (
	"encoding/csv"
	"log"
	"os"
)

func initStorage(filename string) error {
	_, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func loadCSV(filename string, demiliter rune) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(file)
	r.Comma = demiliter

	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}
