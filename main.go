package main

import (
	"os"

	flags "github.com/jessevdk/go-flags"
)

// Options contains go-flags setting.
type Options struct {
	ListStorage bool `long:"list" description:"Show installed storage"`
}

func main() {
	var opts Options
	parser := flags.NewParser(&opts, flags.Default)
	args, err := parser.Parse()
	if err != nil {
		os.Exit(1)
	}

	path := "./storage"

	// インストール済みのストレージを検索
	if opts.ListStorage {
		commandListStorage(args, path)
		return
	}

	// クエリを元に検索を実行
	if len(args) > 0 {
		commandExecute(args, path)
		return
	}
}
