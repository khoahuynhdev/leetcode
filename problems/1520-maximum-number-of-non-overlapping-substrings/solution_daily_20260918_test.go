package main

import (
	"reflect"
	"sort"
	"testing"
)

// sorted returns a sorted copy so results can be compared regardless of order,
// as the problem allows returning the substrings in any order.
func sorted(in []string) []string {
	out := append([]string(nil), in...)
	sort.Strings(out)
	return out
}

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected []string
	}{
		{
			name:     "example 1: nested candidate is skipped in favor of two short ones",
			s:        "adefaddaccc",
			expected: []string{"e", "f", "ccc"},
		},
		{
			name:     "example 2: tie broken by minimum total length",
			s:        "abbaccd",
			expected: []string{"bb", "cc", "d"},
		},
		{
			name:     "edge case: single character",
			s:        "a",
			expected: []string{"a"},
		},
		{
			name:     "edge case: all characters identical",
			s:        "aaaa",
			expected: []string{"aaaa"},
		},
		{
			name:     "edge case: all characters distinct",
			s:        "abcd",
			expected: []string{"a", "b", "c", "d"},
		},
		{
			name:     "edge case: interleaved letters force a single substring",
			s:        "abab",
			expected: []string{"abab"},
		},
		{
			name:     "edge case: fully nested letters, innermost wins",
			s:        "abcba",
			expected: []string{"c"},
		},
		{
			name:     "edge case: enclosing candidate discarded for two inner ones",
			s:        "cbacdcbc",
			expected: []string{"a", "d"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxNumOfSubstrings(tt.s)
			if !reflect.DeepEqual(sorted(result), sorted(tt.expected)) {
				t.Errorf("maxNumOfSubstrings(%q) = %v, want %v (any order)", tt.s, result, tt.expected)
			}
		})
	}
}
