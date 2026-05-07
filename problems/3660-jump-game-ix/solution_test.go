package main

import (
	"reflect"
	"testing"
)

func TestMaxValue(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "example 1: [2,1,3] new max isolates last index",
			nums:     []int{2, 1, 3},
			expected: []int{2, 2, 3},
		},
		{
			name:     "example 2: [2,3,1] valley bridges all",
			nums:     []int{2, 3, 1},
			expected: []int{3, 3, 3},
		},
		{
			name:     "edge case: single element",
			nums:     []int{42},
			expected: []int{42},
		},
		{
			name:     "edge case: strictly increasing",
			nums:     []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "edge case: strictly decreasing",
			nums:     []int{5, 4, 3, 2, 1},
			expected: []int{5, 5, 5, 5, 5},
		},
		{
			name:     "edge case: all equal values yield no edges",
			nums:     []int{5, 5, 5},
			expected: []int{5, 5, 5},
		},
		{
			name:     "plateau then dip: [3,1,4,2]",
			nums:     []int{3, 1, 4, 2},
			expected: []int{4, 4, 4, 4},
		},
		{
			name:     "two components: [1,3,2]",
			nums:     []int{1, 3, 2},
			expected: []int{1, 3, 3},
		},
		{
			name:     "two components with later cluster: [1,5,2,4,3]",
			nums:     []int{1, 5, 2, 4, 3},
			expected: []int{1, 5, 5, 5, 5},
		},
		{
			name:     "small valley merges all: [1,3,0]",
			nums:     []int{1, 3, 0},
			expected: []int{3, 3, 3},
		},
		{
			name:     "isolated peak at end: [5,1,4,2,6]",
			nums:     []int{5, 1, 4, 2, 6},
			expected: []int{5, 5, 5, 5, 6},
		},
		{
			name:     "two-element ascending: [1,2]",
			nums:     []int{1, 2},
			expected: []int{1, 2},
		},
		{
			name:     "two-element descending: [2,1]",
			nums:     []int{2, 1},
			expected: []int{2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxValue(tt.nums)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("maxValue(%v) = %v, want %v", tt.nums, got, tt.expected)
			}
		})
	}
}
