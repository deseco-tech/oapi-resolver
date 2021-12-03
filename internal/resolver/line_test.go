package resolver

import (
	"strings"
	"testing"
)

func TestLineSetIndentation(t *testing.T) {
	var line *Line
	line = NewLine([]byte("    $ref"))
	if line.indent != "    " {
		t.Fatalf("expected to have: 4 spaces, got: %d", len(line.indent))
	}
	line = NewLine([]byte("    $ref: components"))
	if line.indent != "    " {
		t.Fatalf("expected to have: 4 spaces, got: %d", len(line.indent))
	}
	line = NewLine([]byte("$ref"))
	if line.indent != "" {
		t.Fatalf("expected to have: 0 spaces, got: %d", len(line.indent))
	}
}

func TestLineIsRef(t *testing.T) {
	var line *Line
	var result bool
	line = NewLine([]byte("    $ref: components/types.yaml#/Item"))
	result = line.isRef()
	if result != true {
		t.Fatalf("expected to have: true, got: %v", result)
	}
	line = NewLine([]byte("    $ref: #/components/schemas/Item"))
	result = line.isRef()
	if result != false {
		t.Fatalf("expected to have: false, got: %v", result)
	}
}

func TestLineExtractRef(t *testing.T) {
	expctedKey := "Item"
	expctedFname := "components/types.yaml"

	line := NewLine([]byte("$ref: components/types.yaml#/Item"))
	line.extractRef()
	if line.ref.key != expctedKey {
		t.Fatalf("expected to have: %s, got: %s", expctedKey, line.ref.key)
	}
	if line.ref.fname != expctedFname {
		t.Fatalf("expected to have: %s, got: %s", expctedFname, line.ref.fname)
	}
}

func TestLineSetString(t *testing.T) {
	str := "Item:\n    type: object"
	expected := "    Item:\n        type: object"

	line := NewLine([]byte("    $ref"))
	line.setString(str)
	if strings.Trim(line.String(), "\n") != expected {
		t.Fatalf("expected to have: %s, got: %s", line.String(), expected)
	}
}
