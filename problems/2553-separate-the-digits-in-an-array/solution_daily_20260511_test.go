package main

import (
	"reflect"
	"testing"
)

func TestSeparateDigits(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "example 1: mixed two-digit numbers",
			nums:     []int{13, 25, 83, 77},
			expected: []int{1, 3, 2, 5, 8, 3, 7, 7},
		},
		{
			name:     "example 2: all single digits stay the same",
			nums:     []int{7, 1, 3, 9},
			expected: []int{7, 1, 3, 9},
		},
		{
			name:     "edge case: single multi-digit element",
			nums:     []int{10921},
			expected: []int{1, 0, 9, 2, 1},
		},
		{
			name:     "edge case: maximum value with internal zeros",
			nums:     []int{100000},
			expected: []int{1, 0, 0, 0, 0, 0},
		},
		{
			name:     "edge case: single-element input of value 1",
			nums:     []int{1},
			expected: []int{1},
		},
		{
			name:     "edge case: mix of 1, 2, and 3 digit numbers",
			nums:     []int{5, 42, 678, 9},
			expected: []int{5, 4, 2, 6, 7, 8, 9},
		},
		{
			name:     "edge case: repeated identical numbers preserve order",
			nums:     []int{77, 77, 77},
			expected: []int{7, 7, 7, 7, 7, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := separateDigits(tt.nums)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("separateDigits(%v) = %v, want %v", tt.nums, got, tt.expected)
			}
		})
	}
}
