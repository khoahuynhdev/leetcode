package main

import "testing"

func TestFindGCD(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"example 1: min 2, max 10, gcd 2", []int{2, 5, 6, 9, 10}, 2},
		{"example 2: min 3, max 8, coprime gcd 1", []int{7, 5, 6, 8, 3}, 1},
		{"example 3: all equal, gcd is the value", []int{3, 3}, 3},
		{"edge case: contains 1 forces gcd 1", []int{1, 6, 9, 12}, 1},
		{"edge case: min divides max evenly", []int{4, 8, 16, 12}, 4},
		{"edge case: two-element coprime pair", []int{13, 17}, 1},
		{"edge case: two-element with divisor", []int{15, 5}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findGCD(tt.nums)
			if result != tt.expected {
				t.Errorf("findGCD(%v) = %d, want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
