package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{"example 1: bcbbbcba, best window is the trailing bcba", "bcbbbcba", 4},
		{"example 2: aaaa, no window can exceed two a's", "aaaa", 2},
		{"edge case: minimum length with distinct characters", "ab", 2},
		{"edge case: minimum length with repeated character", "aa", 2},
		{"edge case: entire string already valid with pairs", "aabbcc", 6},
		{"edge case: all distinct characters, answer is full length", "abcdefg", 7},
		{"edge case: single character repeated many times", "zzzzzzzzzz", 2},
		{"edge case: repeating abc cycle, no full-string window", "abcabcabc", 6},
		{"edge case: run of three then a fresh run", "aaabbb", 4},
		{"edge case: best window sits at the front", "aabaa", 3},
		{"edge case: violation appears only at the final character", "abcdeaa", 6},
		{"edge case: full 26-letter alphabet twice over", "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz", 52},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximumLengthSubstring(tt.s)
			if result != tt.expected {
				t.Errorf("maximumLengthSubstring(%q) = %v, want %v", tt.s, result, tt.expected)
			}
		})
	}
}
