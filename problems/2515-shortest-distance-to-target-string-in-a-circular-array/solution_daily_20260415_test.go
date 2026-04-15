package main

import "testing"

func TestClosetTarget(t *testing.T) {
	tests := []struct {
		name       string
		words      []string
		target     string
		startIndex int
		expected   int
	}{
		{
			name:       "example 1: target appears twice, shortest is one step left",
			words:      []string{"hello", "i", "am", "leetcode", "hello"},
			target:     "hello",
			startIndex: 1,
			expected:   1,
		},
		{
			name:       "example 2: single occurrence reachable in one step",
			words:      []string{"a", "b", "leetcode"},
			target:     "leetcode",
			startIndex: 0,
			expected:   1,
		},
		{
			name:       "example 3: target not present returns -1",
			words:      []string{"i", "eat", "leetcode"},
			target:     "ate",
			startIndex: 0,
			expected:   -1,
		},
		{
			name:       "edge case: startIndex already at target",
			words:      []string{"a", "b", "c"},
			target:     "b",
			startIndex: 1,
			expected:   0,
		},
		{
			name:       "edge case: single element array matches",
			words:      []string{"only"},
			target:     "only",
			startIndex: 0,
			expected:   0,
		},
		{
			name:       "edge case: single element array no match",
			words:      []string{"x"},
			target:     "y",
			startIndex: 0,
			expected:   -1,
		},
		{
			name:       "edge case: wrap around is shorter than direct",
			words:      []string{"a", "b", "c", "d", "e"},
			target:     "a",
			startIndex: 4,
			expected:   1,
		},
		{
			name:       "edge case: equidistant from start, picks one step",
			words:      []string{"a", "b", "a"},
			target:     "a",
			startIndex: 1,
			expected:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := closetTarget(tt.words, tt.target, tt.startIndex)
			if got != tt.expected {
				t.Errorf("closetTarget(%v, %q, %d) = %d, want %d",
					tt.words, tt.target, tt.startIndex, got, tt.expected)
			}
		})
	}
}
