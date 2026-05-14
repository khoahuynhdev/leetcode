package main

import "testing"

func TestIsGood(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected bool
	}{
		{"example 1: length too short for max=3", []int{2, 1, 3}, false},
		{"example 2: permutation of base[3]", []int{1, 3, 3, 2}, true},
		{"example 3: base[1]", []int{1, 1}, true},
		{"example 4: length too long for max=4", []int{3, 4, 4, 1, 2, 1}, false},
		{"edge case: missing middle value", []int{1, 3, 3, 3}, false},
		{"edge case: duplicate non-max value", []int{1, 1, 2, 3}, false},
		{"edge case: all max repeated, missing 1..n-1", []int{3, 3, 3, 3}, false},
		{"edge case: base[2]", []int{1, 2, 2}, true},
		{"edge case: base[5] shuffled", []int{5, 1, 4, 2, 3, 5}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]int, len(tt.input))
			copy(input, tt.input)
			got := isGood(input)
			if got != tt.expected {
				t.Errorf("isGood(%v) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
