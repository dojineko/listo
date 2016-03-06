package main

import (
	"strconv"
	"strings"
)

type Item struct {
	Autocomplete string
	Title        string
	Subtitle     string
	Icon         string
}

func findAny(query []string, records [][]string) []Item {
	var result []Item
	for key, record := range records {
		joined := strings.Join(record, " ")
		autocomplete := ":" + strconv.Itoa(key) + " "

		result = append(result, Item{
			Autocomplete: autocomplete,
			Title:        joined,
			Subtitle:     "RecordID: " + strconv.Itoa(key),
		})
	}
	return result
}

func findLine(query []string, record []string) []Item {
	var result []Item

	for _, column := range record {
		result = append(result, Item{
			Title: column,
		})
	}

	return result
}

func findColumn(query []string, records [][]string) []Item {
	var result []Item
	return result
}
