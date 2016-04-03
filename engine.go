package main

import (
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

func findStorage(path string, filter string) []Item {
	var result []Item

	filelist := getStorageList(path, true, filter)
	for _, filename := range filelist {
		result = append(result, Item{
			Autocomplete: "@" + filename + " ",
			Title:        filename,
			Subtitle:     "Type: " + filepath.Ext(filename),
		})
	}

	return result
}

func findInStorage(query []string, records [][]string, prefix AlfredItemModifier) []Item {
	var result []Item

	for i, record := range records {
		joined := strings.Join(record, " ")
		if !existsStrings(joined, query) {
			continue
		}

		result = append(result, Item{
			Autocomplete: prefix.AutoComplete + " " + ":" + strconv.Itoa(i) + " ",
			Title:        joined,
			Subtitle:     prefix.Subtitle + ", RecordNo: " + strconv.Itoa(i),
		})
	}

	return result
}

func findInRecord(query []string, record []string, prefix AlfredItemModifier) []Item {
	var result []Item

	withSubQuery := len(query) > 1
	for i, column := range record {
		if withSubQuery && !existsStrings(column, query[1:]) {
			continue
		}

		result = append(result, Item{
			Autocomplete: prefix.AutoComplete + " " + query[0] + " " + column,
			Title:        column,
			Subtitle:     prefix.Subtitle + ", ColumnNo: " + strconv.Itoa(i),
			Arg:          column,
		})
	}

	return result
}

func findInAllStorage(query []string, path string) []Item {
	var result []Item

	filelist := getStorageList(path, true, "")
	for _, filename := range filelist {
		records, err := loadCSV(path+"/"+filename, '\t')
		if err != nil {
			log.Fatal(err)
		}

		prefix := AlfredItemModifier{
			AutoComplete: "@" + filename,
			Subtitle:     "Storage: " + filename,
		}

		result = append(result, findInStorage(query, records, prefix)...)
	}

	return result
}
