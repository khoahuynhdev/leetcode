package main

import (
	"reflect"
	"testing"
)

func TestConstructTransformedArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "example 1: mixed positive and negative values",
			nums:     []int{3, -2, 1, 1},
			expected: []int{1, 1, 1, 3},
		},
		{
			name:     "example 2: negative and positive movements",
			nums:     []int{-1, 4, -1},
			expected: []int{-1, -1, 4},
		},
		{
			name:     "edge case: single element",
			nums:     []int{5},
			expected: []int{5},
		},
		{
			name:     "edge case: all zeros",
			nums:     []int{0, 0, 0},
			expected: []int{0, 0, 0},
		},
		{
			name:     "edge case: large positive steps",
			nums:     []int{10, -5, 7},
			expected: []int{-5, 7, 10},
		},
		{
			name:     "edge case: large negative steps",
			nums:     []int{1, -10, 3},
			expected: []int{-10, 1, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := constructTransformedArray(tt.nums)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
