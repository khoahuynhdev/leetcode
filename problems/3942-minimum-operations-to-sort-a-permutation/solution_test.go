package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"example 1: rotate then reverse", []int{0, 2, 1}, 2},
		{"example 2: reverse then rotate", []int{1, 0, 2}, 2},
		{"example 3: not a rotation of sorted or reversed", []int{2, 0, 1, 3}, -1},
		{"edge case: single element", []int{0}, 0},
		{"edge case: already sorted", []int{0, 1, 2, 3}, 0},
		{"edge case: two elements swapped", []int{1, 0}, 1},
		{"edge case: fully reversed", []int{3, 2, 1, 0}, 1},
		{"edge case: one left rotation needed", []int{3, 0, 1, 2}, 1},
		{"edge case: reverse-rotate-reverse beats rotating", []int{1, 2, 3, 4, 0}, 3},
		{"edge case: impossible larger permutation", []int{0, 2, 1, 3, 4}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minOperations(tt.nums)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
