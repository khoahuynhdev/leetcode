package main

import "testing"

func TestMaximumTrionicSubarraySum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int64
	}{
		{
			name: "Example 1: mixed positive and negative",
			nums: []int{0, -2, -1, -3, 0, 2, -1},
			want: -4,
		},
		{
			name: "Example 2: simple positive case",
			nums: []int{1, 4, 2, 7},
			want: 14,
		},
		{
			name: "Minimal valid trionic",
			nums: []int{1, 2, 1, 2},
			want: 6,
		},
		{
			name: "Longer trionic pattern",
			nums: []int{5, 6, 4, 3, 2, 3, 4},
			want: 27, // entire array: [5,6] up, [6,4,3,2] down, [2,3,4] up
		},
		{
			name: "All negative numbers",
			nums: []int{-5, -4, -6, -3},
			want: -18,
		},
		{
			name: "Multiple potential trionic subarrays",
			nums: []int{1, 3, 5, 4, 2, 1, 3, 5, 7},
			want: 31, // entire array: [1,3,5] up, [5,4,2,1] down, [1,3,5,7] up
		},
		{
			name: "Multiple peaks and valleys",
			nums: []int{1, 2, 3, 2, 1, 0, 1, 2, 3},
			want: 15,
		},
		{
			name: "Large positive values",
			nums: []int{10, 11, 9, 8, 9, 10, 11, 12},
			want: 80, // entire array: [10,11] up, [11,9,8] down, [8,9,10,11,12] up
		},
		{
			name: "Oscillating pattern",
			nums: []int{1, 2, 1, 2, 1, 2},
			want: 6, // best trionic is any [1,2,1,2] = 6
		},
		{
			name: "Boundary values - large negatives",
			nums: []int{-1000000000, -999999999, -1000000000, -999999999},
			want: -3999999998,
		},
		{
			name: "Long array with optimal in middle",
			nums: []int{1, 2, 3, 4, 3, 2, 1, 2, 3, 4, 5},
			want: 30, // entire array: [1,2,3,4] up, [4,3,2,1] down, [1,2,3,4,5] up
		},
		{
			name: "Simple four element case",
			nums: []int{1, 5, 4, 8},
			want: 18,
		},
		{
			name: "Entire array is trionic",
			nums: []int{1, 2, 1, 2, 3},
			want: 9,
		},
		{
			name: "Sharp peak and valley",
			nums: []int{1, 10, 2, 20},
			want: 33,
		},
		{
			name: "Multiple segments with negatives",
			nums: []int{-1, 0, -2, -1, 0, 1},
			want: -3,
		},
		{
			name: "Leetcode 388",
			nums: []int{2, 993, -791, -635, -569},
			want: -431,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumTrionicSubarraySum(tt.nums)
			if got != tt.want {
				t.Errorf("maximumTrionicSubarraySum(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
