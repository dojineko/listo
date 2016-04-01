package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// Item contains a result record
type Item struct {
	Autocomplete string
	Title        string
	Subtitle     string
	Icon         string
	Arg          string
}

func existsStrings(src string, query []string) bool {
	for _, v := range query {
		if !strings.Contains(src, v) {
			return false
		}
	}
	return true
}

func findStorage(query []string, path string) []Item {
	var result []Item
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		filename := file.Name()
		if m, _ := regexp.MatchString("^\\.", filename); m {
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

func findAny(query []string, records [][]string, prefix string) []Item {
	var result []Item
	for i, record := range records {
		joined := strings.Join(record, " ")
		if !existsStrings(joined, query) {
			continue
		}

		result = append(result, Item{
			Autocomplete: prefix + " " + ":" + strconv.Itoa(i) + " ",
			Title:        joined,
			Subtitle:     "RecordID: " + strconv.Itoa(i),
		})
	}
	return result
}

func findLine(query []string, record []string, prefix string) []Item {
	var result []Item

	withSubQuery := len(query) > 1
	for i, column := range record {
		if withSubQuery && !existsStrings(column, query[1:]) {
			continue
		}

		result = append(result, Item{
			Autocomplete: prefix + " " + query[0] + " " + column,
			Title:        column,
			Subtitle:     "ColumnNo: " + strconv.Itoa(i),
			Arg:          column,
		})
	}

	return result
}
