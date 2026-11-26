package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected bool
	}{
		{"example 1: Alice removes 10, Bob cannot remove 9", 12, true},
		{"example 2: Alice cannot make her first move", 1, false},
		{"edge case: one short of Alice's first move", 9, false},
		{"edge case: exactly 10 stones, Alice empties the pile", 10, true},
		{"edge case: one short of Bob's first move", 18, true},
		{"edge case: exactly 19 stones, Bob empties the pile", 19, false},
		{"edge case: exactly 27 stones, Alice's second move empties the pile", 27, true},
		{"edge case: 26 stones, Alice cannot remove 8", 26, false},
		{"edge case: maximum constraint", 50, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := canAliceWin(tt.n)
			if result != tt.expected {
				t.Errorf("canAliceWin(%d) = %v, want %v", tt.n, result, tt.expected)
			}
		})
	}
}
