package main

import "testing"

func TestMaxSelectedElements(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"example 1: mixed values with duplicates", []int{2, 1, 5, 1, 1}, 3},
		{"example 2: values too far apart", []int{1, 4, 7, 10}, 1},
		{"edge case: single element", []int{5}, 1},
		{"edge case: all duplicates of two only", []int{1, 1}, 2},
		{"edge case: many duplicates can fill only two slots", []int{3, 3, 3, 3}, 2},
		{"edge case: gap of one bridged by increment", []int{1, 1, 3}, 3},
		{"edge case: already consecutive", []int{1, 2, 3, 4, 5}, 5},
		{"edge case: shift whole block", []int{1, 3, 3, 4}, 4},
		{"edge case: gap of two cannot be bridged", []int{1, 4}, 1},
		{"edge case: large values", []int{1000000, 999999, 1000000}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxSelectedElements(tt.nums)
			if got != tt.expected {
				t.Errorf("maxSelectedElements(%v) = %d, want %d", tt.nums, got, tt.expected)
			}
		})
	}
}
