package main

import (
	"github.com/deseco-tech/oapi-resolver/internal/resolver"
)

var (
	inputFile    = "../data/api.yaml"
	outputFile   = "../data/combined.yaml"
	expectedFile = "../data/expected.yaml"
)

func run() error {
	processor, err := resolver.New(inputFile, outputFile)
	if err != nil {
		return err
	}
	if err := processor.Combaine(); err != nil {
		return err
	}
	return nil
}
