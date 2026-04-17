package main

import "testing"

func TestMinAbsoluteDistance(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"example 1: adjacent mirror pair wins", []int{12, 21, 45, 33, 54}, 1},
		{"example 2: trailing-zero reverse", []int{120, 21}, 1},
		{"example 3: asymmetric order, no pair", []int{21, 120}, -1},
		{"edge case: single element", []int{5}, -1},
		{"edge case: palindrome numbers pair with themselves", []int{11, 11}, 1},
		{"edge case: no mirror pair at all", []int{1, 2, 3}, -1},
		{"edge case: closer pair appears later in scan", []int{12, 99, 21, 12, 21}, 1},
		{"edge case: two-element matching pair", []int{10, 1}, 1},
		{"edge case: large palindrome", []int{121, 121}, 1},
		{"edge case: only distant pair matches", []int{12, 56, 78, 89, 21}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minAbsoluteDistance(tt.nums)
			if got != tt.expected {
				t.Errorf("minAbsoluteDistance(%v) = %d, want %d", tt.nums, got, tt.expected)
			}
		})
	}
}

func TestReverseDigits(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"single digit", 7, 7},
		{"two digits", 12, 21},
		{"trailing zero dropped", 120, 21},
		{"palindrome", 121, 121},
		{"max 9-digit", 987654321, 123456789},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseDigits(tt.input)
			if got != tt.expected {
				t.Errorf("reverseDigits(%d) = %d, want %d", tt.input, got, tt.expected)
			}
		})
	}
}
