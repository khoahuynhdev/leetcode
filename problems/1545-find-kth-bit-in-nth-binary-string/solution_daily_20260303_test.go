package main

import "testing"

func TestFindKthBit(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		k        int
		expected byte
	}{
		{"example 1: S3 k=1 first bit", 3, 1, '0'},
		{"example 2: S4 k=11", 4, 11, '1'},
		{"edge case: n=1 base case", 1, 1, '0'},
		{"edge case: middle bit of S2", 2, 2, '1'},
		{"edge case: last bit of S3", 3, 7, '1'},
		{"edge case: large n first bit always 0", 20, 1, '0'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findKthBit(tt.n, tt.k)
			if result != tt.expected {
				t.Errorf("findKthBit(%d, %d) = %c, want %c", tt.n, tt.k, result, tt.expected)
			}
		})
	}
}
