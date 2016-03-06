package main

import (
	"flag"
	"regexp"
	"strconv"
)

func main() {
	storage := *(flag.String("storage", "./storage.csv", "Set storage path."))
	format := *(flag.String("format", "alfred", "Set output format."))
	flag.Parse()

	records, err := loadCSV(storage, ',')
	if err != nil {
		initStorage(storage)
		records = nil
	}

	query := flag.Args()
	if len(query) == 0 {
		return
	}

	var result []Item
	linePattern := "^:([0-9]+?)$"
	if m, _ := regexp.MatchString(linePattern, query[0]); m {
		regexLine, _ := regexp.Compile(linePattern)
		match := regexLine.FindStringSubmatch(query[0])[1]
		matchInt, _ := strconv.Atoi(match)

		result = findLine(query, records[matchInt])
	} else {
		result = findAny(query, records)
	}

	switch format {
	case "alfred":
		printAlfred(result)
	default:
	}
}
