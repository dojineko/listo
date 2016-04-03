package main

import (
	"os"

	flags "github.com/jessevdk/go-flags"
)

// Options contains go-flags setting.
type Options struct {
}

var opts Options

func main() {
	parser := flags.NewParser(&opts, flags.Default)
	args, _ := parser.Parse()
	if len(args) == 0 {
		parser.WriteHelp(os.Stdout)
		os.Exit(1)
	}

	query := args
	path := "./storage"

	result := execute(query, path)
	printAlfred(result)
}
