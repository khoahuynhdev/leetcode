package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		k        int
		expected int
	}{
		{"example 1: split into ab and cac", "abcac", 2, 1},
		{"example 2: split into abc and def", "abcdef", 2, 2},
		{"example 3: already semi-palindromes", "aabbaa", 3, 0},
		{"edge case: minimum length two equal chars", "aa", 1, 0},
		{"edge case: minimum length two different chars", "ab", 1, 1},
		{"edge case: whole string semi-palindrome with d=3", "abcabc", 1, 0},
		{"edge case: all pairs forced", "abcdef", 3, 3},
		{"edge case: single part uses best divisor", "abab", 1, 0},
		{"edge case: uniform string many parts", "aaaaaaaa", 4, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minimumChanges(tt.s, tt.k)
			if result != tt.expected {
				t.Errorf("minimumChanges(%q, %d) = %d, want %d", tt.s, tt.k, result, tt.expected)
			}
		})
	}
}
