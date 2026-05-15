package main

import "testing"

func TestFindMin(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"example 1: rotated three times", []int{3, 4, 5, 1, 2}, 1},
		{"example 2: rotated four times", []int{4, 5, 6, 7, 0, 1, 2}, 0},
		{"example 3: full rotation (looks unrotated)", []int{11, 13, 15, 17}, 11},
		{"edge case: single element", []int{1}, 1},
		{"edge case: single negative element", []int{-5000}, -5000},
		{"edge case: two elements rotated once", []int{2, 1}, 1},
		{"edge case: two elements not rotated effectively", []int{1, 2}, 1},
		{"edge case: minimum is negative in rotated array", []int{3, 4, 5, -1, 0, 1, 2}, -1},
		{"edge case: rotation pivot near the end", []int{2, 3, 4, 5, 6, 7, 1}, 1},
		{"edge case: rotation pivot near the start", []int{7, 1, 2, 3, 4, 5, 6}, 1},
		{"edge case: all negative numbers rotated", []int{-2, -1, -5, -4, -3}, -5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMin(tt.input)
			if got != tt.expected {
				t.Errorf("findMin(%v) = %d, want %d", tt.input, got, tt.expected)
			}
		})
	}
}
