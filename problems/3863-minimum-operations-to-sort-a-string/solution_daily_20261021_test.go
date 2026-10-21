package main

import "testing"

func TestMinOperations(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{"example 1: first char is minimum", "dog", 1},
		{"example 2: needs two operations", "card", 2},
		{"example 3: length two unsorted is impossible", "gf", -1},
		{"edge case: single character", "a", 0},
		{"edge case: already sorted", "abcz", 0},
		{"edge case: length two sorted", "ab", 0},
		{"edge case: all same characters", "zzzz", 0},
		{"edge case: last char is maximum", "bcaz", 1},
		{"edge case: min only at end, max only at start", "cba", 3},
		{"edge case: longer reversed string", "zyxwa", 3},
		{"edge case: duplicate max allows two ops", "zbza", 2},
		{"edge case: duplicate min allows two ops", "dabca", 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.s); got != tt.expected {
				t.Errorf("minOperations(%q) = %d, want %d", tt.s, got, tt.expected)
			}
		})
	}
}
