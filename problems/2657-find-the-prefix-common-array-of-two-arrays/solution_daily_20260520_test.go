package main

import (
	"reflect"
	"testing"
)

func TestFindThePrefixCommonArrayDaily20260520(t *testing.T) {
	tests := []struct {
		name     string
		a        []int
		b        []int
		expected []int
	}{
		{
			name:     "example 1: A=[1,3,2,4] B=[3,1,2,4]",
			a:        []int{1, 3, 2, 4},
			b:        []int{3, 1, 2, 4},
			expected: []int{0, 2, 3, 4},
		},
		{
			name:     "example 2: A=[2,3,1] B=[3,1,2]",
			a:        []int{2, 3, 1},
			b:        []int{3, 1, 2},
			expected: []int{0, 1, 3},
		},
		{
			name:     "edge case: single element",
			a:        []int{1},
			b:        []int{1},
			expected: []int{1},
		},
		{
			name:     "edge case: identical arrays",
			a:        []int{1, 2, 3, 4, 5},
			b:        []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "edge case: fully reversed",
			a:        []int{1, 2, 3, 4},
			b:        []int{4, 3, 2, 1},
			expected: []int{0, 0, 2, 4},
		},
		{
			name:     "edge case: no early commonality",
			a:        []int{1, 2, 3},
			b:        []int{3, 2, 1},
			expected: []int{0, 1, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findThePrefixCommonArray(tt.a, tt.b)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("findThePrefixCommonArray(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
