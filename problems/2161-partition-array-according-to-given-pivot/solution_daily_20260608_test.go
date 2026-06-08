package main

import (
	"reflect"
	"testing"
)

func TestPivotArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		pivot    int
		expected []int
	}{
		{
			name:     "example 1: mixed values with two pivots",
			nums:     []int{9, 12, 5, 10, 14, 3, 10},
			pivot:    10,
			expected: []int{9, 5, 3, 10, 10, 12, 14},
		},
		{
			name:     "example 2: includes negatives",
			nums:     []int{-3, 4, 3, 2},
			pivot:    2,
			expected: []int{-3, 2, 4, 3},
		},
		{
			name:     "edge case: single element equal to pivot",
			nums:     []int{5},
			pivot:    5,
			expected: []int{5},
		},
		{
			name:     "edge case: all elements equal to pivot",
			nums:     []int{7, 7, 7, 7},
			pivot:    7,
			expected: []int{7, 7, 7, 7},
		},
		{
			name:     "edge case: all less than pivot except one pivot",
			nums:     []int{4, 1, 3, 5},
			pivot:    5,
			expected: []int{4, 1, 3, 5},
		},
		{
			name:     "edge case: all greater than pivot except one pivot",
			nums:     []int{5, 9, 6, 8},
			pivot:    5,
			expected: []int{5, 9, 6, 8},
		},
		{
			name:     "edge case: negatives on both sides",
			nums:     []int{-5, 0, -2, 1, 0, -1},
			pivot:    0,
			expected: []int{-5, -2, -1, 0, 0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pivotArray(tt.nums, tt.pivot)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("pivotArray(%v, %d) = %v, want %v", tt.nums, tt.pivot, result, tt.expected)
			}
		})
	}
}
