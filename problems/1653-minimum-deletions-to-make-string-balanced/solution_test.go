package main

import "testing"

func TestMinimumDeletions(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "example 1: mixed a's and b's",
			s:        "aababbab",
			expected: 2,
		},
		{
			name:     "example 2: b's at start and end",
			s:        "bbaaaaabb",
			expected: 2,
		},
		{
			name:     "edge case: all a's",
			s:        "aaaa",
			expected: 0,
		},
		{
			name:     "edge case: all b's",
			s:        "bbbb",
			expected: 0,
		},
		{
			name:     "edge case: single character a",
			s:        "a",
			expected: 0,
		},
		{
			name:     "edge case: single character b",
			s:        "b",
			expected: 0,
		},
		{
			name:     "edge case: already balanced",
			s:        "aaabbb",
			expected: 0,
		},
		{
			name:     "edge case: completely reversed",
			s:        "bbbaaa",
			expected: 3,
		},
		{
			name:     "edge case: alternating",
			s:        "ababab",
			expected: 2,
		},
		{
			name:     "edge case: single b followed by a",
			s:        "ba",
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minimumDeletions(tt.s)
			if result != tt.expected {
				t.Errorf("minimumDeletions(%q) = %d, want %d", tt.s, result, tt.expected)
			}
		})
	}
}
