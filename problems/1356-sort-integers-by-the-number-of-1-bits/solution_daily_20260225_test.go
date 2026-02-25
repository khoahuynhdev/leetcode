package main

import (
	"reflect"
	"testing"
)

func TestSortByBits(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "example 1: mixed bit counts 0-8",
			input:    []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
			expected: []int{0, 1, 2, 4, 8, 3, 5, 6, 7},
		},
		{
			name:     "example 2: powers of two descending",
			input:    []int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1},
			expected: []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024},
		},
		{
			name:     "edge case: single element",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "edge case: all same value",
			input:    []int{7, 7, 7},
			expected: []int{7, 7, 7},
		},
		{
			name:     "edge case: zero only",
			input:    []int{0},
			expected: []int{0},
		},
		{
			name:     "edge case: same bit count different values",
			input:    []int{3, 5, 6, 9, 10, 12},
			expected: []int{3, 5, 6, 9, 10, 12},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// copy input to avoid mutating test data for debugging
			inp := make([]int, len(tt.input))
			copy(inp, tt.input)
			result := sortByBits(inp)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
