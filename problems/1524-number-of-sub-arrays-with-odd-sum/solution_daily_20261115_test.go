package main

import "testing"

func TestNumOfSubarrays(t *testing.T) {
	// All odd values, n = 100000: prefix parity alternates odd/even, giving
	// 50000 odd prefixes and 50001 even prefixes (including empty).
	// Count = 50000 * 50001 = 2500050000; mod 1e9+7 = 500049986.
	allOddLarge := make([]int, 100000)
	for i := range allOddLarge {
		allOddLarge[i] = 1
	}

	tests := []struct {
		name     string
		arr      []int
		expected int
	}{
		{"example 1: all odd small", []int{1, 3, 5}, 4},
		{"example 2: all even", []int{2, 4, 6}, 0},
		{"example 3: mixed sequence", []int{1, 2, 3, 4, 5, 6, 7}, 16},
		{"edge case: single odd element", []int{7}, 1},
		{"edge case: single even element", []int{100}, 0},
		{"edge case: odd surrounded by evens", []int{2, 1, 2}, 4},
		{"edge case: large input requires modulo", allOddLarge, 500049986},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := numOfSubarrays(tt.arr)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
