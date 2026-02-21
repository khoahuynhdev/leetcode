package main

import "testing"

func TestCountPrimeSetBits(t *testing.T) {
	tests := []struct {
		name     string
		left     int
		right    int
		expected int
	}{
		{
			name:     "example 1: 6 to 10",
			left:     6,
			right:    10,
			expected: 4,
		},
		{
			name:     "example 2: 10 to 15",
			left:     10,
			right:    15,
			expected: 5,
		},
		{
			name:     "edge case: single number with prime bit count",
			left:     3,
			right:    3,
			expected: 1, // 3 = 11 -> 2 set bits, 2 is prime
		},
		{
			name:     "edge case: single number with non-prime bit count",
			left:     1,
			right:    1,
			expected: 0, // 1 = 1 -> 1 set bit, 1 is not prime
		},
		{
			name:     "edge case: power of two (1 set bit, not prime)",
			left:     8,
			right:    8,
			expected: 0, // 8 = 1000 -> 1 set bit
		},
		{
			name:     "edge case: left equals right at minimum",
			left:     1,
			right:    1,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countPrimeSetBits(tt.left, tt.right)
			if result != tt.expected {
				t.Errorf("countPrimeSetBits(%d, %d) = %d, want %d", tt.left, tt.right, result, tt.expected)
			}
		})
	}
}
