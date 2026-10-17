package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"example 1: two elements need one swap", []int{37, 100}, 1},
		{"example 2: already sorted by digit sum", []int{22, 14, 33, 7}, 0},
		{"example 3: ties broken by value", []int{18, 43, 34, 16}, 2},
		{"edge case: single element", []int{5}, 0},
		{"edge case: single cycle of length three", []int{3, 1, 2}, 2},
		{"edge case: all same digit sum reversed", []int{70, 61, 52, 43, 34, 25, 16}, 3},
		{"edge case: large values", []int{999999999, 1000000000, 10}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minSwaps(tt.nums)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
