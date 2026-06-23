package main

import "testing"

func TestZigZagArrays(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		l        int
		r        int
		expected int
	}{
		{
			name:     "example 1: n=3, range [4,5] -> [4,5,4] and [5,4,5]",
			n:        3,
			l:        4,
			r:        5,
			expected: 2,
		},
		{
			name:     "example 2: n=3, range [1,3] has 10 zigzag arrays",
			n:        3,
			l:        1,
			r:        3,
			expected: 10,
		},
		{
			name:     "edge case: minimal range m=2 forces strict alternation",
			n:        3,
			l:        1,
			r:        2,
			expected: 2,
		},
		{
			name:     "edge case: range is shift-invariant, [2,4] matches [1,3]",
			n:        3,
			l:        2,
			r:        4,
			expected: 10,
		},
		{
			name:     "edge case: longer array n=4 over range [1,3]",
			n:        4,
			l:        1,
			r:        3,
			expected: 16,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := zigZagArrays(tt.n, tt.l, tt.r)
			if result != tt.expected {
				t.Errorf("zigZagArrays(%d, %d, %d) = %d, want %d",
					tt.n, tt.l, tt.r, result, tt.expected)
			}
		})
	}
}
