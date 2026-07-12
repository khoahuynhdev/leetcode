package main

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "example 1: distinct increasing ranks",
			input:    []int{40, 10, 20, 30},
			expected: []int{4, 1, 2, 3},
		},
		{
			name:     "example 2: all equal share rank",
			input:    []int{100, 100, 100},
			expected: []int{1, 1, 1},
		},
		{
			name:     "example 3: duplicates with mixed order",
			input:    []int{37, 12, 28, 9, 100, 56, 80, 5, 12},
			expected: []int{5, 3, 4, 2, 8, 6, 7, 1, 3},
		},
		{
			name:     "edge case: empty input",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "edge case: single element",
			input:    []int{42},
			expected: []int{1},
		},
		{
			name:     "edge case: negative values",
			input:    []int{-5, -10, 0, -10, 5},
			expected: []int{2, 1, 3, 1, 4},
		},
		{
			name:     "edge case: reverse sorted",
			input:    []int{9, 7, 5, 3, 1},
			expected: []int{5, 4, 3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := arrayRankTransform(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("arrayRankTransform(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
