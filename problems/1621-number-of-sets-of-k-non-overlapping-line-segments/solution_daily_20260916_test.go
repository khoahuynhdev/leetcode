package main

import "testing"

func TestNumberOfSets(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		k        int
		expected int
	}{
		{"example 1: n=4, k=2 gives 5 arrangements", 4, 2, 5},
		{"example 2: n=3, k=1 counts every index pair", 3, 1, 3},
		{"example 3: n=30, k=7 needs the modulus", 30, 7, 796297179},
		{"edge case: smallest input n=2, k=1", 2, 1, 1},
		{"edge case: k=n-1 forces the unit-segment chain", 6, 5, 1},
		{"edge case: k=n-1 at the upper bound", 1000, 999, 1},
		{"edge case: k=1 reduces to C(n,2)", 1000, 1, 499500},
		{"edge case: n=5, k=2", 5, 2, 15},
		{"edge case: large n and k exercise modular reduction", 1000, 500, 70047606},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := numberOfSets(tt.n, tt.k); result != tt.expected {
				t.Errorf("numberOfSets(%d, %d) = %v, want %v", tt.n, tt.k, result, tt.expected)
			}
		})
	}
}

// TestNumberOfSetsAgainstBruteForce cross-checks the closed form against an
// explicit enumeration of every valid chain of segments for small inputs.
func TestNumberOfSetsAgainstBruteForce(t *testing.T) {
	for n := 2; n <= 9; n++ {
		for k := 1; k <= n-1; k++ {
			want := bruteForceSets1621(n, k)
			if got := numberOfSets(n, k); got != want {
				t.Errorf("numberOfSets(%d, %d) = %v, want %v (brute force)", n, k, got, want)
			}
		}
	}
}

// bruteForceSets1621 enumerates chains a1 < b1 <= a2 < b2 <= ... < bk <= n-1,
// counting each set of segments exactly once via its left-to-right ordering.
func bruteForceSets1621(n, k int) int {
	var count func(start, remaining int) int
	count = func(start, remaining int) int {
		if remaining == 0 {
			return 1
		}
		total := 0
		for a := start; a < n; a++ {
			for b := a + 1; b < n; b++ {
				total += count(b, remaining-1)
			}
		}
		return total
	}
	return count(0, k)
}
