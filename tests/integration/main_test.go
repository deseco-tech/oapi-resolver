package main

import (
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	err := run()
	if err != nil {
		t.Fatalf("expected no errors: got %v", err)
	}
	expectedBytes, err := os.ReadFile(expectedFile)
	if err != nil {
		t.Fatalf("expected no errors: got %v", err)
	}
	resultBytes, err := os.ReadFile(outputFile)
	if err != nil {
		t.Fatalf("expected no errors: got %v", err)
	}
	result := string(resultBytes)
	expected := string(expectedBytes)
	if result != expected {
		t.Fatalf("expected to have %s: got %s", expected, result)
	}
}
