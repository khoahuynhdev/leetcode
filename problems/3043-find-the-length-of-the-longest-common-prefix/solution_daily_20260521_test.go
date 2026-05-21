package main

import "testing"

func TestLongestCommonPrefixDaily20260521(t *testing.T) {
	tests := []struct {
		name     string
		arr1     []int
		arr2     []int
		expected int
	}{
		{
			name:     "example 1: prefixes of 1000 match 100",
			arr1:     []int{1, 10, 100},
			arr2:     []int{1000},
			expected: 3,
		},
		{
			name:     "example 2: no common prefix for any pair",
			arr1:     []int{1, 2, 3},
			arr2:     []int{4, 4, 4},
			expected: 0,
		},
		{
			name:     "edge case: identical numbers share full length",
			arr1:     []int{12345},
			arr2:     []int{12345},
			expected: 5,
		},
		{
			name:     "edge case: shorter number caps the prefix length",
			arr1:     []int{100},
			arr2:     []int{1},
			expected: 1,
		},
		{
			name:     "edge case: same-array prefixes do not count",
			arr1:     []int{123, 456},
			arr2:     []int{789},
			expected: 0,
		},
		{
			name:     "edge case: max-size nine-digit values match fully",
			arr1:     []int{100000000},
			arr2:     []int{100000000},
			expected: 9,
		},
		{
			name:     "edge case: best pair comes from a later element",
			arr1:     []int{5, 56, 567},
			arr2:     []int{9, 5678},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestCommonPrefix(tt.arr1, tt.arr2)
			if result != tt.expected {
				t.Errorf("longestCommonPrefix(%v, %v) = %d, want %d", tt.arr1, tt.arr2, result, tt.expected)
			}
		})
	}
}
