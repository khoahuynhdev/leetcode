package main

import (
	"reflect"
	"testing"
)

func TestTwoEditWords(t *testing.T) {
	tests := []struct {
		name       string
		queries    []string
		dictionary []string
		expected   []string
	}{
		{
			name:       "example 1: mixed matches within two edits",
			queries:    []string{"word", "note", "ants", "wood"},
			dictionary: []string{"wood", "joke", "moat"},
			expected:   []string{"word", "note", "wood"},
		},
		{
			name:       "example 2: no query reachable in two edits",
			queries:    []string{"yes"},
			dictionary: []string{"not"},
			expected:   []string{},
		},
		{
			name:       "edge case: single character words always match",
			queries:    []string{"a", "z", "m"},
			dictionary: []string{"b"},
			expected:   []string{"a", "z", "m"},
		},
		{
			name:       "edge case: exact match counts as zero edits",
			queries:    []string{"abc", "xyz"},
			dictionary: []string{"abc"},
			expected:   []string{"abc"},
		},
		{
			name:       "edge case: exactly two edits is accepted",
			queries:    []string{"abcd"},
			dictionary: []string{"abef"},
			expected:   []string{"abcd"},
		},
		{
			name:       "edge case: exactly three edits is rejected",
			queries:    []string{"abcd"},
			dictionary: []string{"aefg"},
			expected:   []string{},
		},
		{
			name:       "edge case: duplicates preserve order and repetition",
			queries:    []string{"abc", "abc", "xyz"},
			dictionary: []string{"abd"},
			expected:   []string{"abc", "abc"},
		},
		{
			name:       "edge case: matches any one of multiple dictionary words",
			queries:    []string{"hello"},
			dictionary: []string{"world", "hella", "jumps"},
			expected:   []string{"hello"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoEditWords(tt.queries, tt.dictionary)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("twoEditWords(%v, %v) = %v, want %v",
					tt.queries, tt.dictionary, got, tt.expected)
			}
		})
	}
}
