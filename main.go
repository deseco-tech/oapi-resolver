package main

import (
	"flag"
	"log"

	"github.com/deseco-tech/oapi-resolver/internal/resolver"
)

func main() {
	var inputFile, outputFile string

	flag.StringVar(&inputFile, "i", "", "input file path")
	flag.StringVar(&outputFile, "o", "", "output file path")
	flag.Parse()

	processor, err := resolver.New(inputFile, outputFile)
	if err != nil {
		log.Fatalf("unable to create manager: %v", err)
	}

	if err := processor.Combaine(); err != nil {
		log.Fatalf("unable to combine spec: %v", err)
	}
}
