package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		expected bool
	}{
		{"example 1: mixed pair, all odd via 2-3", []int{2, 3}, true},
		{"example 2: both even, keep both", []int{4, 6}, true},
		{"edge case: single element, odd", []int{7}, true},
		{"edge case: single element, even", []int{8}, true},
		{"edge case: exactly one odd among many evens", []int{1, 2, 4, 6}, true},
		{"edge case: exactly one even among many odds", []int{2, 3, 5, 7}, true},
		{"all odd", []int{1, 3, 5, 9}, true},
		{"all even", []int{2, 10, 42, 100}, true},
		{"mixed, odd not first", []int{4, 7, 10, 12}, true},
		{"constraint bounds: min and max values", []int{1, 100}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := canConstructUniformParityArray(tt.nums1)
			if result != tt.expected {
				t.Errorf("canConstructUniformParityArray(%v) = %v, want %v", tt.nums1, result, tt.expected)
			}
		})
	}
}

// TestBuildUniformParityArrayIsLegal checks that the witness nums2 the solution
// constructs actually obeys the rules: every entry equals nums1[i] or
// nums1[i] - nums1[j] for some j != i, and all entries share one parity.
func TestBuildUniformParityArrayIsLegal(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
	}{
		{"example 1", []int{2, 3}},
		{"example 2", []int{4, 6}},
		{"single element", []int{7}},
		{"one odd among evens", []int{1, 2, 4, 6}},
		{"one even among odds", []int{2, 3, 5, 7}},
		{"mixed, odd not first", []int{4, 7, 10, 12}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums2 := buildUniformParityArray(tt.nums1)

			if len(nums2) != len(tt.nums1) {
				t.Fatalf("len(nums2) = %d, want %d", len(nums2), len(tt.nums1))
			}
			if !isUniformParity(nums2) {
				t.Errorf("nums2 = %v is not uniform in parity", nums2)
			}

			for i, got := range nums2 {
				if got == tt.nums1[i] {
					continue
				}
				legal := false
				for j := range tt.nums1 {
					if j != i && got == tt.nums1[i]-tt.nums1[j] {
						legal = true
						break
					}
				}
				if !legal {
					t.Errorf("nums2[%d] = %d is not nums1[%d] nor nums1[%d]-nums1[j] for any j != %d",
						i, got, i, i, i)
				}
			}
		})
	}
}

func TestIsUniformParity(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{"empty is trivially uniform", []int{}, true},
		{"single element", []int{5}, true},
		{"all odd", []int{1, 3, 5}, true},
		{"all even", []int{2, 4, 6}, true},
		{"negatives stay odd", []int{-1, 3, -5}, true},
		{"mixed is not uniform", []int{1, 2}, false},
		{"mixed with negatives", []int{-3, 4}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUniformParity(tt.nums); got != tt.expected {
				t.Errorf("isUniformParity(%v) = %v, want %v", tt.nums, got, tt.expected)
			}
		})
	}
}
