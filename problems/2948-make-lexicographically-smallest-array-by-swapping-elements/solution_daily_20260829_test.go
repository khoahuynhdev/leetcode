package main

import (
	"reflect"
	"testing"
)

func TestLexicographicallySmallestArrayDaily20260829(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		limit    int
		expected []int
	}{
		{
			name:     "example 1: two independent adjacent pairs swap",
			nums:     []int{1, 5, 3, 9, 8},
			limit:    2,
			expected: []int{1, 3, 5, 8, 9},
		},
		{
			name:     "example 2: duplicates share a group with 2",
			nums:     []int{1, 7, 6, 18, 2, 1},
			limit:    3,
			expected: []int{1, 6, 7, 18, 1, 2},
		},
		{
			name:     "example 3: every gap exceeds limit so nothing moves",
			nums:     []int{1, 7, 28, 19, 10},
			limit:    3,
			expected: []int{1, 7, 28, 19, 10},
		},
		{
			name:     "edge case: single element",
			nums:     []int{5},
			limit:    1,
			expected: []int{5},
		},
		{
			name:     "edge case: all values identical",
			nums:     []int{3, 3, 3},
			limit:    1,
			expected: []int{3, 3, 3},
		},
		{
			name:     "edge case: transitive chain links 1-2-3 but not 8",
			nums:     []int{3, 1, 2, 8},
			limit:    1,
			expected: []int{1, 2, 3, 8},
		},
		{
			name:     "edge case: huge limit sorts the whole array",
			nums:     []int{5, 4, 3, 2, 1},
			limit:    1000000000,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "edge case: already sorted input is unchanged",
			nums:     []int{1, 2, 3, 4},
			limit:    2,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "edge case: large values near the constraint ceiling",
			nums:     []int{1000000000, 999999999, 1},
			limit:    1,
			expected: []int{999999999, 1000000000, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]int, len(tt.nums))
			copy(input, tt.nums)

			result := lexicographicallySmallestArrayDaily20260829(input, tt.limit)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("lexicographicallySmallestArray(%v, %d) = %v, want %v",
					tt.nums, tt.limit, result, tt.expected)
			}
		})
	}
}
