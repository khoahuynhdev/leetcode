package main

import "testing"

func TestFindTheDistanceValue(t *testing.T) {
	tests := []struct {
		name     string
		arr1     []int
		arr2     []int
		d        int
		expected int
	}{
		{"example 1: mixed matches", []int{4, 5, 8}, []int{10, 9, 1, 8}, 2, 2},
		{"example 2: negatives in arr2", []int{1, 4, 2, 3}, []int{-4, -3, 6, 10, 20, 30}, 3, 2},
		{"example 3: only one far element", []int{2, 1, 100, 3}, []int{-5, -2, 10, -3, 7}, 6, 1},
		{"edge case: d is zero with exact match", []int{1, 2, 3}, []int{2}, 0, 2},
		{"edge case: distance exactly d disqualifies", []int{5}, []int{8}, 3, 0},
		{"edge case: all arr2 smaller than window", []int{1000}, []int{-1000, -500}, 100, 1},
		{"edge case: all arr2 larger than window", []int{-1000}, []int{500, 1000}, 100, 1},
		{"edge case: single elements equal", []int{7}, []int{7}, 0, 0},
		{"edge case: duplicates in arr1", []int{3, 3, 50, 50}, []int{6}, 5, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr2 := append([]int(nil), tt.arr2...)
			result := findTheDistanceValue(tt.arr1, arr2, tt.d)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
