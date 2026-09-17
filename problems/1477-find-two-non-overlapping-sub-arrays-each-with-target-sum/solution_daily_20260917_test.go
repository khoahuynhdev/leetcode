package main

import "testing"

func TestMinSumOfLengthsDaily20260917(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{"example 1: only two valid sub-arrays, both length 1", []int{3, 2, 2, 4, 3}, 3, 2},
		{"example 2: three candidates, pick the two shortest non-overlapping", []int{7, 3, 4, 7}, 7, 2},
		{"example 3: only one valid sub-array", []int{4, 3, 2, 6, 2, 3, 4}, 6, -1},
		{"edge case: single element equal to target", []int{1}, 1, -1},
		{"edge case: two candidates that overlap", []int{1, 2, 1}, 3, -1},
		{"edge case: adjacent sub-arrays are allowed", []int{2, 2, 2}, 2, 2},
		{"edge case: whole array is the only candidate", []int{1, 1, 1, 1}, 4, -1},
		{"edge case: target exceeds the total sum", []int{1, 2, 3}, 100, -1},
		{"edge case: shortest overall must pair with a longer one", []int{1, 1, 1, 2, 1}, 2, 3},
		{"edge case: candidates overlap pairwise, best pair is disjoint", []int{2, 1, 2, 1, 2}, 3, 4},
		{"edge case: best pair sits at both ends", []int{5, 1, 1, 1, 5}, 5, 2},
		{"edge case: short prefix candidate plus later candidate", []int{3, 1, 1, 1, 5, 1, 2, 1}, 3, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minSumOfLengthsDaily20260917(tt.arr, tt.target)
			if result != tt.expected {
				t.Errorf("minSumOfLengthsDaily20260917(%v, %d) = %v, want %v", tt.arr, tt.target, result, tt.expected)
			}
		})
	}
}
