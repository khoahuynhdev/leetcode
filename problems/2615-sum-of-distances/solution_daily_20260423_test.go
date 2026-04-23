package main

import (
	"reflect"
	"testing"
)

func TestDistance(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int64
	}{
		{
			name:     "example 1: mixed duplicates and unique values",
			nums:     []int{1, 3, 1, 1, 2},
			expected: []int64{5, 0, 3, 4, 0},
		},
		{
			name:     "example 2: all distinct values",
			nums:     []int{0, 5, 3},
			expected: []int64{0, 0, 0},
		},
		{
			name:     "edge case: single element",
			nums:     []int{1},
			expected: []int64{0},
		},
		{
			name:     "edge case: two identical elements",
			nums:     []int{7, 7},
			expected: []int64{1, 1},
		},
		{
			name:     "edge case: all identical elements",
			nums:     []int{2, 2, 2, 2},
			expected: []int64{6, 4, 4, 6},
		},
		{
			name:     "edge case: two interleaved groups",
			nums:     []int{1, 2, 1, 2, 1, 2},
			expected: []int64{6, 6, 4, 4, 6, 6},
		},
		{
			name:     "edge case: large values used only as keys",
			nums:     []int{1000000000, 1000000000, 5},
			expected: []int64{1, 1, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := distance(tt.nums)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("distance(%v) = %v, want %v", tt.nums, got, tt.expected)
			}
		})
	}
}
