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
			name:     "example 1: third element joins arr1 because 2 > 1",
			nums:     []int{2, 1, 3},
			expected: []int{2, 3, 1},
		},
		{
			name:     "example 2: distribution alternates between both arrays",
			nums:     []int{5, 4, 3, 8},
			expected: []int{5, 3, 4, 8},
		},
		{
			name:     "edge case: minimum length with ascending seed sends element to arr2",
			nums:     []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "edge case: minimum length with descending seed sends element to arr1",
			nums:     []int{3, 2, 1},
			expected: []int{3, 1, 2},
		},
		{
			name:     "edge case: every remaining element lands in arr1",
			nums:     []int{100, 1, 2, 3, 4},
			expected: []int{100, 2, 3, 4, 1},
		},
		{
			name:     "edge case: every remaining element lands in arr2",
			nums:     []int{1, 100, 2, 3, 4},
			expected: []int{1, 100, 2, 3, 4},
		},
		{
			name:     "edge case: compares last appended element, not array maximum",
			nums:     []int{7, 4, 9, 2, 8},
			expected: []int{7, 9, 2, 4, 8},
		},
		{
			name:     "edge case: boundary values 1 and 100 from the constraints",
			nums:     []int{100, 1, 50},
			expected: []int{100, 50, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := append([]int(nil), tt.nums...)
			got := resultArray(input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("resultArray(%v) = %v, want %v", tt.nums, got, tt.expected)
			}
		})
	}
}
