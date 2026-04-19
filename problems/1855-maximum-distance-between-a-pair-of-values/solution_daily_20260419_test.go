package main

import "testing"

func TestMaxDistanceDaily20260419(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		expected int
	}{
		{
			name:     "example 1: mixed values, answer 2 at pair (2,4)",
			nums1:    []int{55, 30, 5, 4, 2},
			nums2:    []int{100, 20, 10, 10, 5},
			expected: 2,
		},
		{
			name:     "example 2: nums1 all equal, answer 1 at pair (0,1)",
			nums1:    []int{2, 2, 2},
			nums2:    []int{10, 10, 1},
			expected: 1,
		},
		{
			name:     "example 3: plateau in nums2, answer 2 at pair (2,4)",
			nums1:    []int{30, 29, 19, 5},
			nums2:    []int{25, 25, 25, 25, 25},
			expected: 2,
		},
		{
			name:     "edge case: no valid pair (nums1 strictly greater than nums2)",
			nums1:    []int{5, 4},
			nums2:    []int{3, 2, 1},
			expected: 0,
		},
		{
			name:     "edge case: single elements, valid pair with distance 0",
			nums1:    []int{1},
			nums2:    []int{1},
			expected: 0,
		},
		{
			name:     "edge case: single elements, invalid pair",
			nums1:    []int{5},
			nums2:    []int{1},
			expected: 0,
		},
		{
			name:     "edge case: nums2 much longer than nums1 maximizes distance",
			nums1:    []int{1},
			nums2:    []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: 9,
		},
		{
			name:     "edge case: equality satisfies the constraint",
			nums1:    []int{5, 5, 5},
			nums2:    []int{5, 5, 5, 5},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxDistanceDaily20260419(tt.nums1, tt.nums2)
			if got != tt.expected {
				t.Errorf("maxDistanceDaily20260419(%v, %v) = %d; want %d",
					tt.nums1, tt.nums2, got, tt.expected)
			}
		})
	}
}
