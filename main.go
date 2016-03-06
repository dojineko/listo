package main

import (
	"flag"
	"strconv"
	"strings"

	"github.com/harasou/alfred"
)

func main() {
	wf := alfred.Workflow()
	defer wf.Print()

	storage := flag.String("storage", "./storage.csv", "Set storage path.")
	// format := flag.String("format", "alfred", "Set output format.")
	flag.Parse()

	records, err := loadCSV(*storage, ',')
	if err != nil {
		initStorage(*storage)
		records = nil
	}

	for key, record := range records {
		joined := strings.Join(record, " ")
		autocomplete := ":" + strconv.Itoa(key)
		wf.AddItem(&alfred.Item{
			Uid:          joined,
			Autocomplete: autocomplete,
			Title:        joined,
			Subtitle:     "RecordID " + autocomplete,
		})
	}
}
