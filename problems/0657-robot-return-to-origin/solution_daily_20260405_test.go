package main

import "testing"

func TestJudgeCircle(t *testing.T) {
	tests := []struct {
		name     string
		moves    string
		expected bool
	}{
		{"example 1: up then down returns to origin", "UD", true},
		{"example 2: two lefts does not return", "LL", false},
		{"edge case: single move", "R", false},
		{"edge case: balanced on both axes", "LRUD", true},
		{"edge case: balanced on one axis only", "LLRR", true},
		{"edge case: long non-returning path", "UUDDLLR", false},
		{"edge case: complex returning path", "UDLRUDLR", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := judgeCircle(tt.moves)
			if result != tt.expected {
				t.Errorf("judgeCircle(%q) = %v, want %v", tt.moves, result, tt.expected)
			}
		})
	}
}
