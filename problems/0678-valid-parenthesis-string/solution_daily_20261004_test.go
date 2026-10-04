package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected bool
	}{
		{"example 1: simple pair", "()", true},
		{"example 2: star as empty", "(*)", true},
		{"example 3: star as open", "(*))", true},
		{"example 4: single open", "(", false},
		{"edge case: single close", ")", false},
		{"edge case: single star", "*", true},
		{"edge case: only stars", "***", true},
		{"edge case: close before open", ")(", false},
		{"edge case: star before open cannot close it", "*(", false},
		{"edge case: not enough stars to close", "((*", false},
		{"edge case: stars close multiple opens", "((**", true},
		{"edge case: star as close at start invalid", "*)))", false},
		{"edge case: long mixed valid", "(*()*)(*)*", true},
		{"edge case: leetcode tricky", "(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := checkValidString(tt.s)
			if result != tt.expected {
				t.Errorf("checkValidString(%q) = %v, want %v", tt.s, result, tt.expected)
			}
		})
	}
}
