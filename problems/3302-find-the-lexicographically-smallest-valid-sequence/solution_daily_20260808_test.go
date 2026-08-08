package main

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		word1    string
		word2    string
		expected []int
	}{
		{
			name:     "example 1: change spent immediately on index 0",
			word1:    "vbcca",
			word2:    "abc",
			expected: []int{0, 1, 2},
		},
		{
			name:     "example 2: change spent in the middle",
			word1:    "bacdc",
			word2:    "abc",
			expected: []int{1, 2, 4},
		},
		{
			name:     "example 3: no valid sequence exists",
			word1:    "aaaaaa",
			word2:    "aaabc",
			expected: []int{},
		},
		{
			name:     "example 4: exact subsequence, change unused",
			word1:    "abc",
			word2:    "ab",
			expected: []int{0, 1},
		},
		{
			name:     "edge case: paying the change beats a free but later index",
			word1:    "abcde",
			word2:    "ace",
			expected: []int{0, 1, 4},
		},
		{
			name:     "edge case: change must be saved for the last character",
			word1:    "abz",
			word2:    "abc",
			expected: []int{0, 1, 2},
		},
		{
			name:     "edge case: two mismatches needed, impossible",
			word1:    "ab",
			word2:    "cd",
			expected: []int{},
		},
		{
			name:     "edge case: single character with no match anywhere",
			word1:    "zz",
			word2:    "a",
			expected: []int{0},
		},
		{
			name:     "edge case: plain prefix match consumes no change",
			word1:    "abcd",
			word2:    "abc",
			expected: []int{0, 1, 2},
		},
		{
			name:     "edge case: change at index 0 unlocks an earlier start",
			word1:    "bbabc",
			word2:    "abc",
			expected: []int{0, 1, 4},
		},
		{
			name:     "edge case: single character exact match at the front",
			word1:    "ab",
			word2:    "a",
			expected: []int{0},
		},
		{
			name:     "edge case: repeated characters, change lands on the tail",
			word1:    "aaaa",
			word2:    "aab",
			expected: []int{0, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validSequence(tt.word1, tt.word2)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("validSequence(%q, %q) = %v, want %v", tt.word1, tt.word2, got, tt.expected)
			}
		})
	}
}
