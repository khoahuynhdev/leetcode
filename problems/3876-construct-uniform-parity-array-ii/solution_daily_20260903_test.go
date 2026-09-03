package main

import "testing"

func TestCanConstructUniformParity(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		expected bool
	}{
		{"example 1: [1,4,7] smallest odd 1 repairs the even 4", []int{1, 4, 7}, true},
		{"example 2: [2,3] smallest even has no smaller odd partner", []int{2, 3}, false},
		{"example 3: [4,6] already all even", []int{4, 6}, true},
		{"edge case: single odd element, subtraction unavailable", []int{5}, true},
		{"edge case: single even element, subtraction unavailable", []int{4}, true},
		{"edge case: all odd stays as-is", []int{9, 3, 7, 1}, true},
		{"edge case: all even stays as-is", []int{1000000000, 2, 48}, true},
		{"edge case: min odd is exactly one below min even", []int{1, 2}, true},
		{"edge case: min even below every odd", []int{2, 5, 7, 9}, false},
		{"edge case: many evens repaired by a single small odd", []int{3, 4, 6, 8, 100}, true},
		{"edge case: one unreachable even blocks an otherwise fixable array", []int{2, 3, 10, 20}, false},
		{"edge case: large values with odd minimum", []int{999999999, 1000000000}, true},
		{"edge case: large values with even minimum", []int{1000000000, 1000000001}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstructUniformParity(tt.nums1); got != tt.expected {
				t.Errorf("canConstructUniformParity(%v) = %v, want %v", tt.nums1, got, tt.expected)
			}
		})
	}
}
