package main

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		parent   []int
		s        string
		expected []int
	}{
		{"example 1: node 3 moves to root", []int{-1, 0, 0, 1, 1, 1}, "abaabc", []int{6, 3, 1, 1, 1, 1}},
		{"example 2: simultaneous reparenting", []int{-1, 0, 4, 0, 1}, "abbba", []int{5, 2, 1, 1, 1}},
		{"edge case: single node", []int{-1}, "a", []int{1}},
		{"edge case: all same chars in a chain stays a chain", []int{-1, 0, 1, 2}, "aaaa", []int{4, 3, 2, 1}},
		{"edge case: all distinct chars, no changes", []int{-1, 0, 1, 2}, "abcd", []int{4, 3, 2, 1}},
		{"edge case: alternating chain flattens into two chains", []int{-1, 0, 1, 2, 3}, "ababa", []int{5, 2, 2, 1, 1}},
		{"edge case: star tree all same char", []int{-1, 0, 0, 0}, "aaaa", []int{4, 1, 1, 1}},
		{"edge case: deep match skips intermediate nodes", []int{-1, 0, 1, 2}, "abca", []int{4, 2, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findSubtreeSizes(tt.parent, tt.s)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
