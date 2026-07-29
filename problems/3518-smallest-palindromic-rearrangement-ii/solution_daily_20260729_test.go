package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		k        int
		expected string
	}{
		{"example 1: two palindromic rearrangements, second one", "abba", 2, "baab"},
		{"example 2: k exceeds the single rearrangement", "aa", 2, ""},
		{"example 3: odd length, smallest rearrangement", "bacab", 1, "abcba"},

		{"edge case: single character, k = 1", "a", 1, "a"},
		{"edge case: single character, k out of range", "a", 2, ""},
		{"edge case: all identical characters, only one arrangement", "aaaa", 1, "aaaa"},
		{"edge case: all identical characters, k out of range", "aaaa", 2, ""},
		{"edge case: k exactly one past the total", "abba", 3, ""},

		{"even length: first of three arrangements", "aabbaa", 1, "aabbaa"},
		{"even length: middle of three arrangements", "aabbaa", 2, "abaaba"},
		{"even length: last of three arrangements", "aabbaa", 3, "baaaab"},
		{"even length: k past the last arrangement", "aabbaa", 4, ""},

		{"odd length: middle character preserved, k = 1", "abcacba", 1, "abcacba"},
		{"odd length: middle character preserved, k = 3", "abcacba", 3, "bacacab"},
		{"odd length: largest arrangement", "abcacba", 6, "cbaaabc"},
		{"odd length: k past the last arrangement", "abcacba", 7, ""},

		{"large k: second of C(8,4) = 70 arrangements", "aaaabbbbbbbbaaaa", 2, "aaababbbbbbabaaa"},
		{"large k: last of C(8,4) = 70 arrangements", "aaaabbbbbbbbaaaa", 70, "bbbbaaaaaaaabbbb"},
		{"large k: one past C(8,4) = 70 arrangements", "aaaabbbbbbbbaaaa", 71, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := smallestPalindrome(tt.s, tt.k)
			if result != tt.expected {
				t.Errorf("smallestPalindrome(%q, %d) = %q, want %q", tt.s, tt.k, result, tt.expected)
			}
		})
	}
}
