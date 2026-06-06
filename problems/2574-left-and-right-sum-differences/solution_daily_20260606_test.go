package main

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "example 1: [10,4,8,3]",
			nums:     []int{10, 4, 8, 3},
			expected: []int{15, 1, 11, 22},
		},
		{
			name:     "example 2: single element",
			nums:     []int{1},
			expected: []int{0},
		},
		{
			name:     "edge case: two elements",
			nums:     []int{5, 9},
			expected: []int{9, 5},
		},
		{
			name:     "edge case: all equal elements",
			nums:     []int{3, 3, 3},
			expected: []int{6, 0, 6},
		},
		{
			name:     "edge case: symmetric array",
			nums:     []int{1, 2, 1},
			expected: []int{3, 0, 3},
		},
		{
			name:     "edge case: max single value",
			nums:     []int{100000},
			expected: []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := leftRightDifference(tt.nums)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("leftRightDifference(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}
