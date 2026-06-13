package main

import "testing"

func TestWeightedWordMapping(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		weights  []int
		expected string
	}{
		{
			name:     "example 1: multiple multi-char words",
			words:    []string{"abcd", "def", "xyz"},
			weights:  []int{5, 3, 12, 14, 1, 2, 3, 2, 10, 6, 6, 9, 7, 8, 7, 10, 8, 9, 6, 9, 9, 8, 3, 7, 7, 2},
			expected: "rij",
		},
		{
			name:     "example 2: single-char words, all weights 1",
			words:    []string{"a", "b", "c"},
			weights:  []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			expected: "yyy",
		},
		{
			name:     "example 3: single word",
			words:    []string{"abcd"},
			weights:  []int{7, 5, 3, 4, 3, 5, 4, 9, 4, 2, 2, 7, 10, 2, 5, 10, 6, 1, 2, 2, 4, 1, 3, 4, 4, 5},
			expected: "g",
		},
		{
			name:    "edge case: weight divisible by 26 maps to 'z'",
			words:   []string{"a"},
			weights: []int{26, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			// weight = 26, 26 % 26 = 0 -> 'z'
			expected: "z",
		},
		{
			name:    "edge case: max-length word with max weights",
			words:   []string{"aaaaaaaaaa"},
			weights: []int{100, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			// weight = 10 * 100 = 1000, 1000 % 26 = 12 -> 'z' - 12 = 'n'
			expected: "n",
		},
		{
			name:    "edge case: word using last letter of alphabet",
			words:   []string{"z"},
			weights: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 25},
			// weight = 25, 25 % 26 = 25 -> 'z' - 25 = 'a'
			expected: "a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := weightedWordMapping(tt.words, tt.weights)
			if result != tt.expected {
				t.Errorf("weightedWordMapping(%v, ...) = %q, want %q", tt.words, result, tt.expected)
			}
		})
	}
}
