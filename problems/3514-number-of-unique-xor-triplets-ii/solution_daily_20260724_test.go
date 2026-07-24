package main

import "testing"

func TestUniqueXorTriplets(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"example 1: two elements {1,3} -> {1,3}", []int{1, 3}, 2},
		{"example 2: {6,7,8,9} -> {6,7,8,9}", []int{6, 7, 8, 9}, 4},
		{"edge case: single element yields only itself", []int{5}, 1},
		{"edge case: all identical collapses to one value", []int{7, 7, 7, 7}, 1},
		{"edge case: duplicates do not add new values", []int{1, 1, 3, 3}, 2},
		{"edge case: value 1 alone", []int{1}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := uniqueXorTriplets(tt.nums)
			if got != tt.expected {
				t.Errorf("uniqueXorTriplets(%v) = %d, want %d", tt.nums, got, tt.expected)
			}
		})
	}
}

// TestUniqueXorTripletsBruteForce cross-checks the optimized solution against a
// straightforward O(n^3) reference on small random-ish inputs.
func TestUniqueXorTripletsBruteForce(t *testing.T) {
	inputs := [][]int{
		{1, 3},
		{6, 7, 8, 9},
		{5},
		{2, 4, 6, 8, 10},
		{1, 2, 3, 4, 5, 6},
		{1500, 1, 1499, 2},
	}

	brute := func(nums []int) int {
		set := map[int]struct{}{}
		n := len(nums)
		for i := 0; i < n; i++ {
			for j := i; j < n; j++ {
				for k := j; k < n; k++ {
					set[nums[i]^nums[j]^nums[k]] = struct{}{}
				}
			}
		}
		return len(set)
	}

	for _, in := range inputs {
		want := brute(in)
		got := uniqueXorTriplets(in)
		if got != want {
			t.Errorf("uniqueXorTriplets(%v) = %d, brute force = %d", in, got, want)
		}
	}
}
