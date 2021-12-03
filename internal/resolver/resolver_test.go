package resolver

import (
	"testing"
)

var (
	inputFile  = "../../tests/data/api.yaml"
	outputFile = "../../tests/data/combained.yaml"
)

func TestResolverNew(t *testing.T) {
	_, err := New("", outputFile)
	if err == nil {
		t.Logf("expect to get error, got nil")
	}
}

func TestExtractCwdFromInputFile(t *testing.T) {
	var result string
	extractCwdFromInputFile("")
	if cwd != "" {
		t.Fatalf("expected to have: %s, got: %s", "", result)
	}
	expected := "test/api"
	extractCwdFromInputFile("test/api/api.yaml")
	if cwd != expected {
		t.Fatalf("expected to have: test/api, got: %s", result)
	}
}
