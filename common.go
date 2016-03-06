package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/dojineko/alfred"
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

func printAlfred(items []Item) {
	var result []alfred.Item

	for _, v := range items {
		item := alfred.Item{
			Autocomplete: v.Autocomplete,
			Title:        v.Title,
			Arg:          v.Arg,
		}
		item.AddSubtitle(v.Subtitle, "")
		result = append(result, item)
	}

	fmt.Print(alfred.Marshal(result))
}
