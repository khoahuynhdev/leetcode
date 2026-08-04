package main

import (
	"reflect"
	"testing"
)

func TestFindMissingElements(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "example 1: single gap in the middle",
			nums:     []int{1, 4, 2, 5},
			expected: []int{3},
		},
		{
			name:     "example 2: nothing missing, range does not start at 1",
			nums:     []int{7, 8, 6, 9},
			expected: []int{},
		},
		{
			name:     "example 3: only the two endpoints are present",
			nums:     []int{5, 1},
			expected: []int{2, 3, 4},
		},
		{
			name:     "edge case: two adjacent values, nothing missing",
			nums:     []int{3, 4},
			expected: []int{},
		},
		{
			name:     "edge case: widest possible range 1..100",
			nums:     []int{1, 100},
			expected: rangeExcept(1, 100, map[int]bool{1: true, 100: true}),
		},
		{
			name:     "edge case: already sorted with multiple separate gaps",
			nums:     []int{2, 3, 6, 8, 9},
			expected: []int{4, 5, 7},
		},
		{
			name:     "edge case: descending input is handled without sorting",
			nums:     []int{9, 6, 3},
			expected: []int{4, 5, 7, 8},
		},
		{
			name:     "edge case: full contiguous range given out of order",
			nums:     []int{12, 10, 13, 11},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMissingElements(tt.nums)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("findMissingElements(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}

// rangeExcept builds the sorted list of integers in [lo, hi] that are not in skip.
func rangeExcept(lo, hi int, skip map[int]bool) []int {
	out := make([]int, 0)
	for v := lo; v <= hi; v++ {
		if !skip[v] {
			out = append(out, v)
		}
	}
	return out
}
