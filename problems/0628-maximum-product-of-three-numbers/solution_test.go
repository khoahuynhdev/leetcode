package main

import "testing"

func TestMaximumProduct(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"example 1: three ascending positives", []int{1, 2, 3}, 6},
		{"example 2: top three positives win", []int{1, 2, 3, 4}, 24},
		{"example 3: all negative, closest to zero", []int{-1, -2, -3}, -6},
		{"edge case: two large negatives beat top three", []int{-100, -98, -1, 2, 3, 4}, 39200},
		{"edge case: four negatives", []int{-1, -2, -3, -4}, -6},
		{"edge case: all zeros", []int{0, 0, 0}, 0},
		{"edge case: zero is the maximum product", []int{-1, 0, 1, 2}, 0},
		{"edge case: negative pair ties top three", []int{-1, -2, 1, 2, 3}, 6},
		{"edge case: duplicate values", []int{1, 1, 1, 1}, 1},
		{"edge case: constraint bounds, all max", []int{1000, 1000, 1000}, 1000000000},
		{"edge case: constraint bounds, negative pair", []int{-1000, -1000, 1000}, 1000000000},
		{"edge case: single negative among positives", []int{-5, 1, 2, 3}, 6},
		{"edge case: mixed with duplicated minimum", []int{-4, -4, 2, 3}, 48},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumProduct(tt.nums)
			if got != tt.expected {
				t.Errorf("maximumProduct(%v) = %v, want %v", tt.nums, got, tt.expected)
			}
		})
	}
}
