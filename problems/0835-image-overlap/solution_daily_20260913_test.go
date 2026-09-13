package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		img1     [][]int
		img2     [][]int
		expected int
	}{
		{
			name:     "example 1: translate right 1 and down 1",
			img1:     [][]int{{1, 1, 0}, {0, 1, 0}, {0, 1, 0}},
			img2:     [][]int{{0, 0, 0}, {0, 1, 1}, {0, 0, 1}},
			expected: 3,
		},
		{
			name:     "example 2: single cell, both ones",
			img1:     [][]int{{1}},
			img2:     [][]int{{1}},
			expected: 1,
		},
		{
			name:     "example 3: single cell, both zeros",
			img1:     [][]int{{0}},
			img2:     [][]int{{0}},
			expected: 0,
		},
		{
			name:     "edge case: n=1 with mismatched cells",
			img1:     [][]int{{1}},
			img2:     [][]int{{0}},
			expected: 0,
		},
		{
			name:     "edge case: img1 is all zeros",
			img1:     [][]int{{0, 0}, {0, 0}},
			img2:     [][]int{{1, 1}, {1, 1}},
			expected: 0,
		},
		{
			name:     "edge case: img2 is all zeros",
			img1:     [][]int{{1, 1}, {1, 1}},
			img2:     [][]int{{0, 0}, {0, 0}},
			expected: 0,
		},
		{
			name:     "edge case: identical dense images need no translation",
			img1:     [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			img2:     [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			expected: 9,
		},
		{
			name:     "edge case: negative offset, translate up and left",
			img1:     [][]int{{0, 0, 0}, {0, 1, 1}, {0, 0, 1}},
			img2:     [][]int{{1, 1, 0}, {0, 1, 0}, {0, 1, 0}},
			expected: 3,
		},
		{
			name:     "edge case: single one each, opposite corners",
			img1:     [][]int{{1, 0}, {0, 0}},
			img2:     [][]int{{0, 0}, {0, 1}},
			expected: 1,
		},
		{
			name:     "edge case: 4x4 pure horizontal translation",
			img1:     [][]int{{0, 0, 0, 0}, {1, 1, 0, 0}, {0, 1, 0, 0}, {0, 0, 0, 0}},
			img2:     [][]int{{0, 0, 0, 0}, {0, 0, 1, 1}, {0, 0, 0, 1}, {0, 0, 0, 0}},
			expected: 3,
		},
		{
			name:     "edge case: identity shift beats every other offset",
			img1:     [][]int{{1, 0, 1}, {0, 1, 0}, {1, 0, 1}},
			img2:     [][]int{{1, 0, 1}, {0, 1, 0}, {1, 0, 1}},
			expected: 5,
		},
		{
			name:     "edge case: disjoint single ones in a 3x3 grid",
			img1:     [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 1}},
			img2:     [][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := largestOverlap(tt.img1, tt.img2)
			if result != tt.expected {
				t.Errorf("largestOverlap(%v, %v) = %v, want %v", tt.img1, tt.img2, result, tt.expected)
			}
		})
	}
}

// TestSolutionSymmetry checks the documented symmetry of the problem: sliding
// img1 one way is equivalent to sliding img2 the other way, so swapping the
// arguments must not change the answer.
func TestSolutionSymmetry(t *testing.T) {
	tests := []struct {
		name string
		img1 [][]int
		img2 [][]int
	}{
		{
			name: "example 1 swapped",
			img1: [][]int{{1, 1, 0}, {0, 1, 0}, {0, 1, 0}},
			img2: [][]int{{0, 0, 0}, {0, 1, 1}, {0, 0, 1}},
		},
		{
			name: "sparse corners",
			img1: [][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 1}},
			img2: [][]int{{0, 0, 1}, {0, 1, 0}, {0, 0, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			forward := largestOverlap(tt.img1, tt.img2)
			backward := largestOverlap(tt.img2, tt.img1)
			if forward != backward {
				t.Errorf("overlap not symmetric: largestOverlap(img1, img2) = %v, largestOverlap(img2, img1) = %v", forward, backward)
			}
		})
	}
}
