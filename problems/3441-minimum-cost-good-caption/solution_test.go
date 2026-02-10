package main

import "testing"

func TestMinCostGoodCaption(t *testing.T) {
	tests := []struct {
		name     string
		caption  string
		expected string
	}{
		{"example 1: cdcd", "cdcd", "cccc"},
		{"example 2: aca", "aca", "aaa"},
		{"example 3: bc (too short)", "bc", ""},
		{"edge case: single character", "a", ""},
		{"edge case: two characters", "ab", ""},
		{"edge case: exactly 3 same", "aaa", "aaa"},
		{"edge case: exactly 3 different", "abc", "bbb"},
		{"edge case: length 6 two groups", "aaabbb", "aaabbb"},
		{"edge case: all same long", "aaaaaa", "aaaaaa"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minCostGoodCaption(tt.caption)
			if result != tt.expected {
				t.Errorf("minCostGoodCaption(%q) = %q, want %q", tt.caption, result, tt.expected)
			}
		})
	}
}
