package main

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		queries  [][]int
		expected []int
	}{
		{
			name:     "example 1: mixed digits with interspersed zeros",
			s:        "10203004",
			queries:  [][]int{{0, 7}, {1, 3}, {4, 6}},
			expected: []int{12340, 4, 9},
		},
		{
			name:     "example 2: trailing zeros and single zero",
			s:        "1000",
			queries:  [][]int{{0, 3}, {1, 1}},
			expected: []int{1, 0},
		},
		{
			name:     "example 3: modulo wrap on large product",
			s:        "9876543210",
			queries:  [][]int{{0, 9}},
			expected: []int{444444137},
		},
		{
			name:     "edge case: all zeros yields zero",
			s:        "0000",
			queries:  [][]int{{0, 3}, {2, 2}},
			expected: []int{0, 0},
		},
		{
			name:     "edge case: single non-zero digit",
			s:        "5",
			queries:  [][]int{{0, 0}},
			expected: []int{25}, // x = 5, sum = 5, 5*5 = 25
		},
		{
			name:     "edge case: contiguous non-zero digits and subranges",
			s:        "12345",
			queries:  [][]int{{0, 4}, {1, 3}, {2, 2}},
			expected: []int{185175, 2106, 9}, // 12345*15, 234*9, 3*3
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := processQueries(tt.s, tt.queries)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("processQueries(%q, %v) = %v, want %v", tt.s, tt.queries, result, tt.expected)
			}
		})
	}
}
