package main

import "testing"

func TestFindTheString(t *testing.T) {
	tests := []struct {
		name     string
		lcp      [][]int
		expected string
	}{
		{
			name:     "example 1: alternating pattern abab",
			lcp:      [][]int{{4, 0, 2, 0}, {0, 3, 0, 1}, {2, 0, 2, 0}, {0, 1, 0, 1}},
			expected: "abab",
		},
		{
			name:     "example 2: all same character aaaa",
			lcp:      [][]int{{4, 3, 2, 1}, {3, 3, 2, 1}, {2, 2, 2, 1}, {1, 1, 1, 1}},
			expected: "aaaa",
		},
		{
			name:     "example 3: invalid lcp[3][3] = 3",
			lcp:      [][]int{{4, 3, 2, 1}, {3, 3, 2, 1}, {2, 2, 2, 1}, {1, 1, 1, 3}},
			expected: "",
		},
		{
			name:     "edge case: single character",
			lcp:      [][]int{{1}},
			expected: "a",
		},
		{
			name:     "edge case: invalid single element lcp[0][0] != 1",
			lcp:      [][]int{{2}},
			expected: "",
		},
		{
			name:     "edge case: asymmetric matrix",
			lcp:      [][]int{{2, 1}, {0, 1}},
			expected: "",
		},
		{
			name:     "edge case: two distinct characters",
			lcp:      [][]int{{2, 0}, {0, 1}},
			expected: "ab",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findTheString(tt.lcp)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}
