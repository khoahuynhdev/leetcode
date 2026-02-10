package main

import "testing"

func TestIsMatch(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		p        string
		expected bool
	}{
		{
			name:     "example 1: pattern too short",
			s:        "aa",
			p:        "a",
			expected: false,
		},
		{
			name:     "example 2: star matches multiple chars",
			s:        "aa",
			p:        "a*",
			expected: true,
		},
		{
			name:     "example 3: dot star matches everything",
			s:        "ab",
			p:        ".*",
			expected: true,
		},
		{
			name:     "edge case: empty string with star pattern",
			s:        "",
			p:        "a*",
			expected: true,
		},
		{
			name:     "edge case: empty string with multiple star patterns",
			s:        "",
			p:        "a*b*c*",
			expected: true,
		},
		{
			name:     "edge case: single dot matches single char",
			s:        "a",
			p:        ".",
			expected: true,
		},
		{
			name:     "edge case: dot star at beginning",
			s:        "aab",
			p:        "c*a*b",
			expected: true,
		},
		{
			name:     "complex pattern with multiple stars",
			s:        "mississippi",
			p:        "mis*is*p*.",
			expected: false,
		},
		{
			name:     "complex pattern matching",
			s:        "mississippi",
			p:        "mis*is*ip*.",
			expected: true,
		},
		{
			name:     "star with no preceding match",
			s:        "ab",
			p:        ".*c",
			expected: false,
		},
		{
			name:     "exact match without wildcards",
			s:        "abc",
			p:        "abc",
			expected: true,
		},
		{
			name:     "pattern longer with stars",
			s:        "a",
			p:        "ab*",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isMatch(tt.s, tt.p)
			if result != tt.expected {
				t.Errorf("isMatch(%q, %q) = %v, want %v", tt.s, tt.p, result, tt.expected)
			}
		})
	}
}
