package adventofcode2023

import (
	"testing"
)

func TestWeekThree(t *testing.T) {
	// Test cases
	tests := []struct {
		input  []string
		output int
	}{
		//go:embed  input.txt
		var inputFile string
	}

	for _, tt := range tests {
		result := WeekThree(tt.input)
		if result != tt.output {
			t.Errorf("Expected %d, but got %d", tt.output, result)
		}
	}
}


