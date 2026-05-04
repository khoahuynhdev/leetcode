package main

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected [][]int
	}{
		{
			name: "example 1: 3x3 matrix",
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			name: "example 2: 4x4 matrix",
			input: [][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			},
			expected: [][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11},
			},
		},
		{
			name: "edge case: 1x1 matrix is unchanged",
			input: [][]int{
				{42},
			},
			expected: [][]int{
				{42},
			},
		},
		{
			name: "edge case: 2x2 smallest non-trivial size",
			input: [][]int{
				{1, 2},
				{3, 4},
			},
			expected: [][]int{
				{3, 1},
				{4, 2},
			},
		},
		{
			name: "edge case: 5x5 with negatives, zeros, and duplicates",
			input: [][]int{
				{-1, -2, -3, -4, -5},
				{1, 2, 3, 4, 5},
				{0, 0, 0, 0, 0},
				{6, 7, 8, 9, 10},
				{-6, -7, -8, -9, -10},
			},
			expected: [][]int{
				{-6, 6, 0, 1, -1},
				{-7, 7, 0, 2, -2},
				{-8, 8, 0, 3, -3},
				{-9, 9, 0, 4, -4},
				{-10, 10, 0, 5, -5},
			},
		},
		{
			name: "edge case: 3x3 with all identical values stays the same",
			input: [][]int{
				{7, 7, 7},
				{7, 7, 7},
				{7, 7, 7},
			},
			expected: [][]int{
				{7, 7, 7},
				{7, 7, 7},
				{7, 7, 7},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.input)
			if !reflect.DeepEqual(tt.input, tt.expected) {
				t.Errorf("got %v, want %v", tt.input, tt.expected)
			}
		})
	}
}
