package main

import "testing"

func TestMinimumPushes(t *testing.T) {
	tests := []struct {
		name     string
		word     string
		expected int
	}{
		{"example 1: five letters all fit in single-push slots", "abcde", 5},
		{"example 2: ten letters spill into second-push slots", "xycdefghij", 12},
		{"edge case: single letter", "a", 1},
		{"edge case: exactly eight letters, one per key", "abcdefgh", 8},
		{"edge case: nine letters, first rollover to two pushes", "abcdefghi", 10},
		{"edge case: full alphabet reaches the fourth tier", "abcdefghijklmnopqrstuvwxyz", 56},
		{"edge case: sixteen letters fill two full tiers", "abcdefghijklmnop", 24},
		{"edge case: letter identity does not matter", "xyz", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumPushes(tt.word); got != tt.expected {
				t.Errorf("minimumPushes(%q) = %v, want %v", tt.word, got, tt.expected)
			}
		})
	}
}
