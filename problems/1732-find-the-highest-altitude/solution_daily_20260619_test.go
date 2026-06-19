package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		gain     []int
		expected int
	}{
		{"example 1: peak at an interior point", []int{-5, 1, 5, 0, -7}, 1},
		{"example 2: never rises above start", []int{-4, -3, -2, -1, 4, 3, 2}, 0},
		{"edge case: single negative gain", []int{-100}, 0},
		{"edge case: single positive gain", []int{100}, 100},
		{"edge case: all positive, increasing altitude", []int{1, 2, 3, 4}, 10},
		{"edge case: rises then falls below start", []int{5, -10, 5}, 5},
		{"edge case: returns to start altitude", []int{3, -3}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := largestAltitude(tt.gain)
			if result != tt.expected {
				t.Errorf("largestAltitude(%v) = %d, want %d", tt.gain, result, tt.expected)
			}
		})
	}
}
