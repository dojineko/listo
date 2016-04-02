package main

import (
	"flag"
	"log"
	"regexp"
	"strconv"
)

func main() {
	flag.Parse()
	query := flag.Args()

	storageDir := "./storage"

	var result []Item
	linePattern := "^:([0-9]+?)$"

	var isFileSelect bool
	if len(query) > 0 {
		isFileSelect, _ = regexp.MatchString("^@.*?$", query[0])
	}
	if !isFileSelect || len(query) == 0 {
		result = findAnyStorage(query, storageDir)
	} else if isFileSelect && len(query) == 1 {
		result = findStorage(query, storageDir)
	} else {
		filename := query[0][1:]
		query = query[1:]
		prefix := AlfredItemPrefix{
			AutoComplete: "@" + filename,
			Subtitle:     "Storage: " + filename,
		}

		records, err := loadCSV(storageDir+"/"+filename, '\t')
		if err != nil {
			log.Fatal(err)
		}

		if len(query) == 0 {
			result = append(result, Item{
				Title:    "キーワードを入力",
				Subtitle: filename + "で検索するキーワードを指定してください",
			})
		} else {
			if m, _ := regexp.MatchString(linePattern, query[0]); m {
				regexLine, _ := regexp.Compile(linePattern)
				match := regexLine.FindStringSubmatch(query[0])[1]
				matchInt, _ := strconv.Atoi(match)
				prefix.Subtitle = prefix.Subtitle + ", RecordNo: " + match

				result = findLine(query, records[matchInt], prefix)
			} else {
				result = findAny(query, records, prefix)
			}
		}

	}

	printAlfred(result)
}
