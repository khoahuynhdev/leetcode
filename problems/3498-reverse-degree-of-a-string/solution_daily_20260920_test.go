package main

import (
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{"example 1: abc, ascending letters", "abc", 148},
		{"example 2: zaza, alternating extremes", "zaza", 160},
		{"edge case: single 'a' maps to 26", "a", 26},
		{"edge case: single 'z' maps to 1", "z", 1},
		{"edge case: empty string sums to zero", "", 0},
		{"edge case: repeated letter, position still varies", "zzz", 6},
		{"edge case: both alphabet endpoints", "az", 28},
		{"edge case: mixed word", "leetcode", 682},
		{"edge case: maximum length of 1000 'a's", strings.Repeat("a", 1000), 13013000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := reverseDegree(tt.s)
			if result != tt.expected {
				t.Errorf("reverseDegree(%q) = %v, want %v", tt.s, result, tt.expected)
			}
		})
	}
}
