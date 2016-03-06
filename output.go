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
	Arg          string
}

func findAny(query []string, records [][]string) []Item {
	var result []Item
	for i, record := range records {
		joined := strings.Join(record, " ")
		autocomplete := ":" + strconv.Itoa(i) + " "

		result = append(result, Item{
			Autocomplete: autocomplete,
			Title:        joined,
			Subtitle:     "RecordID: " + strconv.Itoa(i),
		})
	}
	return result
}

func findLine(query []string, record []string) []Item {
	var result []Item

	isSubQuery := len(query) > 1
	for i, column := range record {
		if isSubQuery && !strings.Contains(column, query[1]) {
			continue
		}
		result = append(result, Item{
			Autocomplete: query[0] + " " + column,
			Title:        column,
			Subtitle:     "ColumnNo: " + strconv.Itoa(i),
			Arg:          column,
		})
	}

	return result
}

func findColumn(query []string, records [][]string) []Item {
	var result []Item
	return result
}
