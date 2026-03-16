package main

import (
	"reflect"
	"testing"
)

func TestGetBiggestThree(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected []int
	}{
		{
			name: "example 1: 5x5 grid with large rhombus sums",
			grid: [][]int{
				{3, 4, 5, 1, 3},
				{3, 3, 4, 2, 3},
				{20, 30, 200, 40, 10},
				{1, 5, 5, 4, 1},
				{4, 3, 2, 2, 5},
			},
			expected: []int{228, 216, 211},
		},
		{
			name: "example 2: 3x3 grid",
			grid: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: []int{20, 9, 8},
		},
		{
			name: "example 3: single row with identical values",
			grid: [][]int{
				{7, 7, 7},
			},
			expected: []int{7},
		},
		{
			name: "edge case: 1x1 grid",
			grid: [][]int{
				{42},
			},
			expected: []int{42},
		},
		{
			name: "edge case: single column",
			grid: [][]int{
				{1},
				{2},
				{3},
				{4},
			},
			expected: []int{4, 3, 2},
		},
		{
			name: "edge case: 2x2 grid no large rhombus fits",
			grid: [][]int{
				{10, 20},
				{30, 40},
			},
			expected: []int{40, 30, 20},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getBiggestThree(tt.grid)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
