package main

import "testing"

func TestMinSumSquareDiff(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		k1       int
		k2       int
		expected int64
	}{
		{"example 1: no modifications allowed", []int{1, 2, 3, 4}, []int{2, 10, 20, 19}, 0, 0, 579},
		{"example 2: partial reduction of tied maximum", []int{1, 4, 10, 12}, []int{5, 8, 6, 9}, 1, 1, 43},
		{"edge case: budget exceeds total difference", []int{1, 2}, []int{4, 6}, 5, 10, 0},
		{"edge case: all diffs already zero", []int{3, 3, 3}, []int{3, 3, 3}, 2, 2, 0},
		{"edge case: single element partially reduced", []int{0}, []int{10}, 3, 0, 49},
		{"edge case: flatten peaks evenly", []int{0, 0}, []int{10, 4}, 0, 7, 25},
		{"edge case: large values need int64", []int{0, 100000}, []int{100000, 0}, 0, 0, 20000000000},
		{"edge case: huge budget", []int{100000}, []int{0}, 1000000000, 1000000000, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minSumSquareDiff(tt.nums1, tt.nums2, tt.k1, tt.k2)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
