package main

import "testing"

func TestMinimumPushes(t *testing.T) {
	tests := []struct {
		name     string
		word     string
		expected int
	}{
		{"example 1: five distinct letters fit the cheap tier", "abcde", 5},
		{"example 2: three letters repeated four times each", "xyzxyzxyzxyz", 12},
		{"example 3: nine distinct letters spill into the second tier", "aabbccddeeffgghhiiiiii", 24},
		{"edge case: single character", "a", 1},
		{"edge case: single letter repeated", "aaaaa", 5},
		{"edge case: exactly eight distinct letters, all cost one push", "abcdefgh", 8},
		{"edge case: nine distinct letters, one lands in the second tier", "abcdefghi", 10},
		{"edge case: all 26 letters use all four tiers", "abcdefghijklmnopqrstuvwxyz", 56},
		{"edge case: order does not matter (interleaved)", "abab", 4},
		{"edge case: order does not matter (grouped)", "aabb", 4},
		{"edge case: skewed frequencies favor the heavy letter", "aaaaaaaaaabcdefghij", 21},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumPushes(tt.word); got != tt.expected {
				t.Errorf("minimumPushes(%q) = %v, want %v", tt.word, got, tt.expected)
			}
		})
	}
}
