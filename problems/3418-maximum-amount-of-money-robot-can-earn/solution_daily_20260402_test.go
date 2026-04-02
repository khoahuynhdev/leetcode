package main

import "testing"

func TestMaximumAmount(t *testing.T) {
	tests := []struct {
		name     string
		coins    [][]int
		expected int
	}{
		{
			name:     "example 1: mixed values with robbers",
			coins:    [][]int{{0, 1, -1}, {1, -2, 3}, {2, -3, 4}},
			expected: 8,
		},
		{
			name:     "example 2: all positive values",
			coins:    [][]int{{10, 10, 10}, {10, 10, 10}},
			expected: 40,
		},
		{
			name:     "edge case: single cell positive",
			coins:    [][]int{{5}},
			expected: 5,
		},
		{
			name:     "edge case: single cell negative",
			coins:    [][]int{{-5}},
			expected: 0,
		},
		{
			name:     "edge case: single row with robbers",
			coins:    [][]int{{1, -10, -20, 1}},
			expected: 2,
		},
		{
			name:     "edge case: all negative values",
			coins:    [][]int{{-1, -2, -3}, {-4, -5, -6}, {-7, -8, -9}},
			expected: -6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximumAmount(tt.coins)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
