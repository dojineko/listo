package main

import "flag"

func main() {
	flag.Parse()
	query := flag.Args()
	path := "./storage"

	result := execute(query, path)
	printAlfred(result)
}
