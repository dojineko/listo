package main

import (
	"log"
	"regexp"
	"strconv"
)

func execute(query []string, path string) []Item {
	// queryが空の場合は終了
	if len(query) == 0 {
		return nil
	}

	// 1つめのクエリが@から始まり、クエリが1つの場合はストレージ絞り込み検索
	var isFileSelect bool
	isFileSelect, _ = regexp.MatchString("^@.*?$", query[0])
	if isFileSelect && len(query) == 1 {
		var storageFilter string
		if len(query[0]) > 1 {
			storageFilter = query[0][1:]
		}
		return findStorage(path, storageFilter)
	}

	// 1つめのクエリが@から始まり、クエリが1より多い場合はストレージ内絞り込み検索
	if isFileSelect && len(query) > 1 {
		// ファイル名のみ抽出して1つ目のクエリを破棄
		filename := query[0][1:]
		query = query[1:]
		prefix := AlfredItemModifier{
			AutoComplete: "@" + filename,
			Subtitle:     "Storage: " + filename,
		}

		// クエリがなくなった場合はプレースホルダを返す
		if len(query) == 0 {
			return []Item{
				Item{
					Title:    "キーワードを入力",
					Subtitle: filename + "で検索するキーワードを指定してください",
				},
			}
		}

		// ストレージを読み込む
		records, err := loadCSV(path+"/"+filename, '\t')
		if err != nil {
			log.Fatal(err)
		}

		// レコード指定がある場合はレコード内絞り込み検索
		linePattern := "^:([0-9]+?)$"
		if isLineSelect, _ := regexp.MatchString(linePattern, query[0]); isLineSelect {
			regexLine, _ := regexp.Compile(linePattern)
			match := regexLine.FindStringSubmatch(query[0])[1]
			matchInt, _ := strconv.Atoi(match)
			prefix.Subtitle = prefix.Subtitle + ", RecordNo: " + match

			// レコード内絞り込み検索
			return findInRecord(query, records[matchInt], prefix)
		}

		// ストレージ内絞り込み検索
		return findInStorage(query, records, prefix)
	}

	// それ以外の場合はストレージ横断検索
	return findInAllStorage(query, path)
}
