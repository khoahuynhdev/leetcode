package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"example 1: rearrange to [1,2,2,2,1]", []int{2, 2, 1, 2, 1}, 2},
		{"example 2: decrease to [1,2,3]", []int{100, 1, 1000}, 3},
		{"example 3: already valid", []int{1, 2, 3, 4, 5}, 5},
		{"edge case: single element greater than 1", []int{17}, 1},
		{"edge case: single element equal to 1", []int{1}, 1},
		{"edge case: all equal values climb by one", []int{5, 5, 5, 5}, 4},
		{"edge case: large gaps capped at length", []int{1000000000, 1, 2}, 3},
		{"edge case: first element not one", []int{3, 4, 5}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximumElementAfterDecreasing(tt.input)
			if result != tt.expected {
				t.Errorf("maximumElementAfterDecreasing(%v) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}
