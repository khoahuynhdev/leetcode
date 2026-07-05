package main

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		board    []string
		expected []int
	}{
		{
			name:     "example 1: single max path",
			board:    []string{"E23", "2X2", "12S"},
			expected: []int{7, 1},
		},
		{
			name:     "example 2: two paths tie for max",
			board:    []string{"E12", "1X1", "21S"},
			expected: []int{4, 2},
		},
		{
			name:     "example 3: no path due to obstacle wall",
			board:    []string{"E11", "XXX", "11S"},
			expected: []int{0, 0},
		},
		{
			name:     "edge case: smallest 2x2 board with digits",
			board:    []string{"E1", "1S"},
			expected: []int{1, 2},
		},
		{
			name:     "edge case: 2x2 with diagonal obstacle forces one path",
			board:    []string{"E1", "XS"},
			expected: []int{1, 1},
		},
		{
			name:     "edge case: diagonal S-to-E move still reachable past corner obstacles",
			board:    []string{"EX", "XS"},
			expected: []int{0, 1},
		},
		{
			name:     "edge case: all nines, max avoids diagonals to collect more cells",
			board:    []string{"E99", "999", "99S"},
			expected: []int{27, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pathsWithMaxScore(tt.board)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("pathsWithMaxScore(%v) = %v, want %v", tt.board, got, tt.expected)
			}
		})
	}
}
