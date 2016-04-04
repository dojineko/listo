package main

import (
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"
)

// Options contains go-flags setting.
type Options struct {
	ListStorage    bool   `long:"list" description:"Show installed storage"`
	InstallStorage string `long:"install" description:"Install a storage"`
	RemoveStorage  string `long:"remove" description:"Remove a installed storage"`
}

func main() {
	var opts Options
	parser := flags.NewParser(&opts, flags.Default)
	args, err := parser.Parse()
	if err != nil {
		os.Exit(1)
	}

	storagePath := "./storage"

	// インストール済みのストレージを検索
	if opts.ListStorage {
		commandListStorage(args, storagePath)
		return
	}

	// ストレージを追加
	if parser.FindOptionByLongName("install").IsSet() {
		err := commandInstallStorage(opts.InstallStorage, storagePath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return
	}

	// ストレージを削除
	if parser.FindOptionByLongName("remove").IsSet() {
		err := commandRemoveStorage(opts.RemoveStorage, storagePath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return
	}

	// クエリを元に検索を実行
	if len(args) > 0 {
		commandExecute(args, storagePath)
		return
	}
}
