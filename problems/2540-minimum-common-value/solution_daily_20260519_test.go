package main

import "testing"

func TestGetCommonDaily20260519(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		expected int
	}{
		{
			name:     "example 1: smallest common is 2",
			nums1:    []int{1, 2, 3},
			nums2:    []int{2, 4},
			expected: 2,
		},
		{
			name:     "example 2: two common values, pick the minimum",
			nums1:    []int{1, 2, 3, 6},
			nums2:    []int{2, 3, 4, 5},
			expected: 2,
		},
		{
			name:     "edge case: no common element returns -1",
			nums1:    []int{1, 2, 3},
			nums2:    []int{4, 5, 6},
			expected: -1,
		},
		{
			name:     "edge case: single-element arrays that match",
			nums1:    []int{7},
			nums2:    []int{7},
			expected: 7,
		},
		{
			name:     "edge case: single-element arrays with no match",
			nums1:    []int{5},
			nums2:    []int{6},
			expected: -1,
		},
		{
			name:     "edge case: common value at the very end",
			nums1:    []int{1, 2, 9},
			nums2:    []int{3, 5, 9},
			expected: 9,
		},
		{
			name:     "edge case: duplicates within arrays",
			nums1:    []int{1, 1, 2, 2, 3},
			nums2:    []int{2, 2, 2, 4},
			expected: 2,
		},
		{
			name:     "edge case: disjoint ranges with very different sizes",
			nums1:    []int{1, 2, 3, 4, 5},
			nums2:    []int{1000000000},
			expected: -1,
		},
		{
			name:     "edge case: large value matches at the boundary",
			nums1:    []int{1, 999999999, 1000000000},
			nums2:    []int{1000000000},
			expected: 1000000000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getCommon(tt.nums1, tt.nums2)
			if got != tt.expected {
				t.Errorf("getCommon(%v, %v) = %d, want %d", tt.nums1, tt.nums2, got, tt.expected)
			}
		})
	}
}
