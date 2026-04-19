package main

import (
	"os"
	"strings"
	"testing"
)

func Test_infer_type(t *testing.T) {
	tests := []struct {
		input    string
		expected any
	}{
		{"123", 123},
		{" 45 ", 45},
		{"1.1", 1.1},
		{"true", true},
		{"false", false},
		{"hello", "hello"},
		{"", nil},
	}

	for _, tt := range tests {
		got := infer_type(tt.input)

		if got != tt.expected {
			t.Errorf("expected %v, got %v", tt.expected, got)
		}
	}
}

func Test_csv_to_json(t *testing.T) {
	input := "test.csv"
	output := "test.jsonl"

	os.WriteFile(input, []byte("name,age\nTim,30\n"), 0644)

	defer os.Remove(input)
	defer os.Remove(output)

	err := csv_to_json(input, output)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, _ := os.ReadFile(output)
	if !strings.Contains(string(data), "Tim") {
		t.Errorf("expected Tim in output: %s", string(data))
	}
}
