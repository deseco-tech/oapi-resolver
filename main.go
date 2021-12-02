package main

import (
	"flag"
)

func main() {
	var inputFile, outputFile string

	flag.StringVar(&inputFile, "i", "", "input file path")
	flag.StringVar(&outputFile, "o", "", "output file path")
	flag.Parse()
}
