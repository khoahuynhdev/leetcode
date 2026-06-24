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
		{"example 1: n=3 range [4,5]", 3, 4, 5, 2},
		{"example 2: n=3 range [1,3]", 3, 1, 3, 10},
		{"edge case: minimum range m=2 forces single alternation", 5, 1, 2, 2},
		{"edge case: huge n with two values stays at 2", 1000000000, 1, 2, 2},
		{"edge case: n=3 over range [1,4]", 3, 1, 4, 28},
		{"edge case: length 5 over range [1,5]", 5, 1, 5, 492},
		{"edge case: offset invariance, range [2,6]", 5, 2, 6, 492},
		{"edge case: huge n over full range [1,75]", 1000000000, 1, 75, 16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := zigZagArrays(tt.n, tt.l, tt.r)
			if got != tt.expected {
				t.Errorf("zigZagArrays(%d, %d, %d) = %d, want %d",
					tt.n, tt.l, tt.r, got, tt.expected)
			}
		})
	}
}
