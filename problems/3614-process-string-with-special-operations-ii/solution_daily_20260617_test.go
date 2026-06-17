package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		k        int64
		expected byte
	}{
		// Examples from the problem statement.
		{"example 1: remove after reverse, result \"ba\"", "a#b%*", 1, 'a'},
		{"example 2: duplicate/remove/duplicate, result \"dcddcd\"", "cd%#*#", 3, 'd'},
		{"example 3: result empty, k out of bounds", "z*#", 0, '.'},

		// Additional edge cases.
		{"edge case: single letter, k in bounds", "x", 0, 'x'},
		{"edge case: single letter, k out of bounds", "x", 1, '.'},
		{"edge case: remove on empty result stays empty", "*", 0, '.'},
		{"edge case: duplicate then index into second copy", "ab#", 2, 'a'},
		{"edge case: reverse a simple string", "abc%", 0, 'c'},
		{"edge case: large k via duplication", "a##", 3, 'a'},
		{"edge case: every char removed", "abc***", 0, '.'},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := processStr(tt.s, tt.k)
			if result != tt.expected {
				t.Errorf("processStr(%q, %d) = %q, want %q",
					tt.s, tt.k, result, tt.expected)
			}
		})
	}
}
