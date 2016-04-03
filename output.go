package main

import (
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

func findStorage(query []string, path string) []Item {
	var result []Item

	var queryFilename string
	if len(query) > 0 && len(query[0]) > 1 {
		queryFilename = query[0][1:]
	}

	filelist := getFileList(path, true)
	for _, filename := range filelist {
		if len(queryFilename) > 0 && !strings.Contains(filename, queryFilename) {
			continue
		}

		result = append(result, Item{
			Autocomplete: "@" + filename + " ",
			Title:        filename,
			Subtitle:     "Type: " + filepath.Ext(filename),
		})
	}

	return result
}

func findAnyStorage(query []string, path string) []Item {
	var result []Item

	filelist := getFileList(path, true)
	for _, filename := range filelist {
		records, err := loadCSV(path+"/"+filename, '\t')
		if err != nil {
			log.Fatal(err)
		}

		prefix := AlfredItemModifier{
			AutoComplete: "@" + filename,
			Subtitle:     "Storage: " + filename,
		}

		result = append(result, findAny(query, records, prefix)...)
	}

	return result
}

func findAny(query []string, records [][]string, prefix AlfredItemModifier) []Item {
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

func findLine(query []string, record []string, prefix AlfredItemModifier) []Item {
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
