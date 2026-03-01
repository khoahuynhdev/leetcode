package main

import "testing"

func TestMinPartitions(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"example 1: two digits summing to 32", "32", 3},
		{"example 2: five digits with max 8", "82734", 8},
		{"example 3: large number with max digit 9", "27346209830709182346", 9},
		{"edge case: single digit 1", "1", 1},
		{"edge case: single digit 9", "9", 9},
		{"edge case: all ones", "111", 1},
		{"edge case: number is itself deci-binary", "10", 1},
		{"edge case: all nines", "999", 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minPartitions(tt.input)
			if result != tt.expected {
				t.Errorf("minPartitions(%q) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}
