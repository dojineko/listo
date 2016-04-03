package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/dojineko/alfred"
)

// AlfredItemModifier constains prefix strings for Alfred result items.
type AlfredItemModifier struct {
	AutoComplete string
	Subtitle     string
}

// Item contains a result record
type Item struct {
	Autocomplete string
	Title        string
	Subtitle     string
	Icon         string
	Arg          string
}

func getStorageList(path string, excludeDotfile bool, filter string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	var result []string

	for _, file := range files {
		filename := file.Name()
		isDotfile, _ := regexp.MatchString("^\\.", filename)
		if excludeDotfile && isDotfile {
			continue
		}
		if len(filter) > 0 && !strings.Contains(filename, filter) {
			continue
		}
		result = append(result, filename)
	}

	return result
}

func existsStrings(src string, query []string) bool {
	for _, v := range query {
		if !strings.Contains(src, v) {
			return false
		}
	}
	return true
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
