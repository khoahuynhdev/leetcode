package main

import "testing"

func TestShortestBeautifulSubstring(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		k        int
		expected string
	}{
		{"example 1: shortest window wins over an earlier longer one", "100011001", 3, "11001"},
		{"example 2: trailing pair of ones is shortest", "1011", 2, "11"},
		{"example 3: no ones at all", "000", 1, ""},
		{"edge case: fewer ones than k", "0110", 3, ""},
		{"edge case: single character string", "1", 1, "1"},
		{"edge case: tie on length needs lexicographic compare", "1101011", 3, "1011"},
		{"edge case: all ones, k smaller than length", "11111", 2, "11"},
		{"edge case: k equals total number of ones", "001101", 3, "1101"},
		{"edge case: leading zeros are never absorbed", "001101", 2, "11"},
		{"edge case: k equals string length", "111", 3, "111"},
		{"edge case: k = 1 always yields a single one", "0001000", 1, "1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := shortestBeautifulSubstring(tt.s, tt.k)
			if result != tt.expected {
				t.Errorf("shortestBeautifulSubstring(%q, %d) = %q, want %q", tt.s, tt.k, result, tt.expected)
			}
		})
	}
}
