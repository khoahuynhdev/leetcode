package main

import "testing"

func TestProcessStr(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected string
	}{
		{"example 1: a#b%* -> ba", "a#b%*", "ba"},
		{"example 2: z*# empties then duplicates empty", "z*#", ""},
		{"edge case: empty result on leading remove", "*abc", "abc"},
		{"edge case: only special chars stays empty", "*%#", ""},
		{"edge case: single letter", "q", "q"},
		{"edge case: duplicate then reverse palindrome-ish", "ab#", "abab"},
		{"edge case: reverse of empty is empty", "%", ""},
		{"edge case: remove all then append", "ab**c", "c"},
		{"edge case: reverse single char unchanged", "a%", "a"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processStr(tt.s); got != tt.expected {
				t.Errorf("processStr(%q) = %q, want %q", tt.s, got, tt.expected)
			}
		})
	}
}
