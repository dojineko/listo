package main

import (
	"io"
	"log"
	"os"
	"path"
	"regexp"
	"strconv"
)

func commandListStorage(args []string, storagePath string) {
	// 引数があればフィルタに使う
	var filter string
	if len(args) > 0 {
		filter = args[0]
	}

	var result []Item
	storages := getStorageList(storagePath, true, filter)
	for _, filename := range storages {
		result = append(result, Item{
			Autocomplete: filename,
			Title:        filename,
			Arg:          filename,
		})
	}

	printAlfred(result)
}

func commandInstallStorage(srcPath string, storagePath string) error {
	filename := path.Base(srcPath)
	destPath := storagePath + "/" + filename

	// コピー元のファイルを開く
	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer src.Close()

	// 出力先のファイルを作る
	dest, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer dest.Close()

	// データのコピーを行う
	_, err = io.Copy(dest, src)
	if err != nil {
		return err
	}

	return nil
}

func commandRemoveStorage(storageName string, storagePath string) error {
	targetPath := storagePath + "/" + storageName
	err := os.Remove(targetPath)
	return err
}

func commandExecute(args []string, storagePath string) {
	printAlfred(commandExecuteInternal(args, storagePath))
}

func commandExecuteInternal(query []string, storagePath string) []Item {
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
		return findStorage(storagePath, storageFilter)
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
		records, err := loadCSV(storagePath+"/"+filename, '\t')
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
	return findInAllStorage(query, storagePath)
}
