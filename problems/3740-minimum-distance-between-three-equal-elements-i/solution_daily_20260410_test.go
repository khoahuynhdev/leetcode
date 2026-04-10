package main

import "testing"

func TestMinimumDistance(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"example 1: three 1s close together", []int{1, 2, 1, 1, 3}, 6},
		{"example 2: three 2s spread out", []int{1, 1, 2, 3, 2, 1, 2}, 8},
		{"example 3: single element", []int{1}, -1},
		{"edge case: no value appears 3 times", []int{1, 2, 1, 2}, -1},
		{"edge case: all elements same", []int{3, 3, 3}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minimumDistance(tt.nums)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
