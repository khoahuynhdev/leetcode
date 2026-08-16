package main

import "testing"

func TestStoneGameIX(t *testing.T) {
	tests := []struct {
		name     string
		stones   []int
		expected bool
	}{
		{"example 1: one stone of each nonzero residue, Alice wins", []int{2, 1}, true},
		{"example 2: single safe stone but pile empties, Bob wins", []int{2}, false},
		{"example 3: balanced residues with one zero stone, Bob wins", []int{5, 1, 2, 4, 3}, false},

		{"edge case: single stone divisible by 3, Alice loses immediately", []int{3}, false},
		{"edge case: every stone divisible by 3, no legal opening move", []int{3, 6, 9}, false},
		{"edge case: odd zero count with large surplus of residue 1", []int{3, 1, 1, 1, 1}, true},
		{"edge case: even zero count but residue 2 missing", []int{1, 1, 1, 1}, false},
		{"edge case: odd zero count with difference exactly 2, Bob wins", []int{3, 1, 1, 1, 2}, false},
		{"edge case: odd zero count with difference exactly 3, Alice wins", []int{3, 1, 1, 1, 1, 2}, true},
		{"edge case: three zero stones still behave as one", []int{3, 6, 9, 1, 1, 1, 1}, true},
		{"edge case: large values reduced mod 3", []int{10000, 9999, 9998}, false},
		{"edge case: even zero count with both residues present", []int{3, 6, 1, 2}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := stoneGameIX(tt.stones); result != tt.expected {
				t.Errorf("stoneGameIX(%v) = %v, want %v", tt.stones, result, tt.expected)
			}
		})
	}
}

func TestStoneGameIXLargeInput(t *testing.T) {
	// 10^5 stones, all with remainder 1: c0 = 0 (even) and c2 = 0, so Bob wins.
	stones := make([]int, 100000)
	for i := range stones {
		stones[i] = 1
	}
	if result := stoneGameIX(stones); result != false {
		t.Errorf("stoneGameIX(100000 stones of value 1) = %v, want false", result)
	}
}
