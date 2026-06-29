package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		patterns []string
		word     string
		expected int
	}{
		{
			name:     "example 1: some patterns match",
			patterns: []string{"a", "abc", "bc", "d"},
			word:     "abc",
			expected: 3,
		},
		{
			name:     "example 2: repeated characters in word",
			patterns: []string{"a", "b", "c"},
			word:     "aaaaabbbbb",
			expected: 2,
		},
		{
			name:     "example 3: duplicate patterns counted separately",
			patterns: []string{"a", "a", "a"},
			word:     "ab",
			expected: 3,
		},
		{
			name:     "edge case: no pattern matches",
			patterns: []string{"x", "yz", "qq"},
			word:     "abc",
			expected: 0,
		},
		{
			name:     "edge case: pattern longer than word",
			patterns: []string{"abcd", "ab"},
			word:     "abc",
			expected: 1,
		},
		{
			name:     "edge case: single-character word and patterns",
			patterns: []string{"a", "b"},
			word:     "a",
			expected: 1,
		},
		{
			name:     "edge case: all patterns match",
			patterns: []string{"ab", "bc", "abc"},
			word:     "abc",
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := numOfStrings(tt.patterns, tt.word)
			if result != tt.expected {
				t.Errorf("numOfStrings(%v, %q) = %d, want %d",
					tt.patterns, tt.word, result, tt.expected)
			}
		})
	}
}
