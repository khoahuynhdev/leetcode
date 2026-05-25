package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		minJump  int
		maxJump  int
		expected bool
	}{
		{"example 1: reachable via 0->3->5", "011010", 2, 3, true},
		{"example 2: blocked by run of ones", "01101110", 2, 3, false},
		{"edge case: minimal length single jump", "00", 1, 1, true},
		{"edge case: last index is '1' is unreachable", "01", 1, 1, false},
		{"edge case: all zeros always reachable", "00000", 1, 2, true},
		{"edge case: minJump equals maxJump exact landing", "0000", 3, 3, true},
		{"edge case: minJump equals maxJump misses target", "0010", 2, 2, false},
		{"edge case: target is '0' but path is blocked", "0100", 1, 1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canReach(tt.s, tt.minJump, tt.maxJump)
			if got != tt.expected {
				t.Errorf("canReach(%q, %d, %d) = %v, want %v",
					tt.s, tt.minJump, tt.maxJump, got, tt.expected)
			}
		})
	}
}
