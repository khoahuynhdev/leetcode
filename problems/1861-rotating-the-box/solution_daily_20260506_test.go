package main

import (
	"reflect"
	"testing"
)

func TestRotateTheBoxDaily20260506(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]byte
		expected [][]byte
	}{
		{
			name: "example 1: single row with stone, empty, stone",
			input: [][]byte{
				{'#', '.', '#'},
			},
			expected: [][]byte{
				{'.'},
				{'#'},
				{'#'},
			},
		},
		{
			name: "example 2: 2x4 grid with obstacle blocking some stones",
			input: [][]byte{
				{'#', '.', '*', '.'},
				{'#', '#', '*', '.'},
			},
			expected: [][]byte{
				{'#', '.'},
				{'#', '#'},
				{'*', '*'},
				{'.', '.'},
			},
		},
		{
			name: "example 3: 3x6 grid with multiple obstacles",
			input: [][]byte{
				{'#', '#', '*', '.', '*', '.'},
				{'#', '#', '#', '*', '.', '.'},
				{'#', '#', '#', '.', '#', '.'},
			},
			expected: [][]byte{
				{'.', '#', '#'},
				{'.', '#', '#'},
				{'#', '#', '*'},
				{'#', '*', '.'},
				{'#', '.', '*'},
				{'#', '.', '.'},
			},
		},
		{
			name: "edge case: single cell empty",
			input: [][]byte{
				{'.'},
			},
			expected: [][]byte{
				{'.'},
			},
		},
		{
			name: "edge case: single cell stone",
			input: [][]byte{
				{'#'},
			},
			expected: [][]byte{
				{'#'},
			},
		},
		{
			name: "edge case: single column with stones and empties",
			input: [][]byte{
				{'#'},
				{'.'},
				{'#'},
				{'.'},
			},
			expected: [][]byte{
				{'.', '#', '.', '#'},
			},
		},
		{
			name: "edge case: row of all empties",
			input: [][]byte{
				{'.', '.', '.', '.'},
			},
			expected: [][]byte{
				{'.'},
				{'.'},
				{'.'},
				{'.'},
			},
		},
		{
			name: "edge case: obstacle at the right edge blocks the stone",
			input: [][]byte{
				{'#', '.', '*'},
			},
			expected: [][]byte{
				{'.'},
				{'#'},
				{'*'},
			},
		},
		{
			name: "edge case: stones already settled to the right",
			input: [][]byte{
				{'.', '.', '#', '#'},
				{'.', '#', '#', '#'},
			},
			expected: [][]byte{
				{'.', '.'},
				{'#', '.'},
				{'#', '#'},
				{'#', '#'},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rotateTheBox(deepCopyBytes(tt.input))
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("rotateTheBox() =\n%s\nwant\n%s", formatGrid(got), formatGrid(tt.expected))
			}
		})
	}
}

func deepCopyBytes(in [][]byte) [][]byte {
	out := make([][]byte, len(in))
	for i := range in {
		out[i] = make([]byte, len(in[i]))
		copy(out[i], in[i])
	}
	return out
}

func formatGrid(g [][]byte) string {
	s := ""
	for _, row := range g {
		s += string(row) + "\n"
	}
	return s
}
