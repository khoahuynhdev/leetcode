package main

import "testing"

func TestNumSteps(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{"example 1: 1101 (13)", "1101", 6},
		{"example 2: 10 (2)", "10", 1},
		{"example 3: already one", "1", 0},
		{"edge case: all ones 1111 (15)", "1111", 5},
		{"edge case: long power of two", "100000", 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := numSteps(tt.s)
			if result != tt.expected {
				t.Errorf("numSteps(%q) = %d, want %d", tt.s, result, tt.expected)
			}
		})
	}
}
