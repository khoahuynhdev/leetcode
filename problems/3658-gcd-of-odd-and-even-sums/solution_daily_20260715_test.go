package main

import "testing"

func TestGcdOfOddEvenSums(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{"example 1: n=4 -> gcd(16,20)=4", 4, 4},
		{"example 2: n=5 -> gcd(25,30)=5", 5, 5},
		{"edge case: smallest n=1 -> gcd(1,2)=1", 1, 1},
		{"edge case: n=2 -> gcd(4,6)=2", 2, 2},
		{"edge case: upper bound n=1000 -> 1000", 1000, 1000},
		{"additional: n=7 -> gcd(49,56)=7", 7, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := gcdOfOddEvenSums(tt.n)
			if result != tt.expected {
				t.Errorf("gcdOfOddEvenSums(%d) = %d, want %d", tt.n, result, tt.expected)
			}
		})
	}
}
