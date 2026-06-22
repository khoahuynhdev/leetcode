package main

import "testing"

func TestMaxNumberOfBalloons(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected int
	}{
		{"example 1: single balloon from scrambled letters", "nlaebolko", 1},
		{"example 2: two balloons with extra letters", "loonbalxballpoon", 2},
		{"example 3: no balloon possible", "leetcode", 0},
		{"edge case: exact balloon spelling", "balloon", 1},
		{"edge case: missing required letter n", "balloo", 0},
		{"edge case: single character", "b", 0},
		{"edge case: enough singles but only one l", "balon", 0},
		{"edge case: three balloons", "balloonballoonballoon", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxNumberOfBalloons(tt.text)
			if result != tt.expected {
				t.Errorf("maxNumberOfBalloons(%q) = %d, want %d", tt.text, result, tt.expected)
			}
		})
	}
}
