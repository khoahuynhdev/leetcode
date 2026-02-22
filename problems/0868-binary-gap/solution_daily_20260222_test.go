package main

import "testing"

func TestBinaryGap(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{"example 1: n=22 binary 10110", 22, 2},
		{"example 2: n=8 binary 1000", 8, 0},
		{"example 3: n=5 binary 101", 5, 2},
		{"edge case: n=1 single bit", 1, 0},
		{"edge case: all ones n=7 binary 111", 7, 1},
		{"edge case: two ones far apart n=33 binary 100001", 33, 5},
		{"large gap: n=1025 binary 10000000001", 1025, 10},
		{"consecutive ones: n=6 binary 110", 6, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := binaryGap(tt.n)
			if result != tt.expected {
				t.Errorf("binaryGap(%d) = %d, want %d", tt.n, result, tt.expected)
			}
		})
	}
}
