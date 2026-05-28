package main

import (
	"reflect"
	"testing"
)

func TestStringIndices(t *testing.T) {
	tests := []struct {
		name           string
		wordsContainer []string
		wordsQuery     []string
		expected       []int
	}{
		{
			name:           "example 1: shared cd suffix tiebreak by length",
			wordsContainer: []string{"abcd", "bcd", "xbcd"},
			wordsQuery:     []string{"cd", "bcd", "xyz"},
			expected:       []int{1, 1, 1},
		},
		{
			name:           "example 2: longest suffix beats shorter length",
			wordsContainer: []string{"abcdefgh", "poiuygh", "ghghgh"},
			wordsQuery:     []string{"gh", "acbfgh", "acbfegh"},
			expected:       []int{2, 0, 2},
		},
		{
			name:           "edge case: single container word and no suffix overlap",
			wordsContainer: []string{"abc"},
			wordsQuery:     []string{"xyz"},
			expected:       []int{0},
		},
		{
			name:           "edge case: no overlap falls back to earliest of equal-length",
			wordsContainer: []string{"abc", "xyz", "pqr"},
			wordsQuery:     []string{"hhh"},
			expected:       []int{0},
		},
		{
			name:           "edge case: progressively longer suffix matches",
			wordsContainer: []string{"a", "ba", "bba"},
			wordsQuery:     []string{"a", "ba", "cba"},
			expected:       []int{0, 1, 1},
		},
		{
			name:           "edge case: query equals a container word exactly",
			wordsContainer: []string{"hello", "world"},
			wordsQuery:     []string{"world"},
			expected:       []int{1},
		},
		{
			name:           "edge case: duplicate container words pick earliest index",
			wordsContainer: []string{"abc", "abc", "xabc"},
			wordsQuery:     []string{"abc", "zzz"},
			expected:       []int{0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stringIndices(tt.wordsContainer, tt.wordsQuery)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("stringIndices(%v, %v) = %v, want %v",
					tt.wordsContainer, tt.wordsQuery, got, tt.expected)
			}
		})
	}
}
