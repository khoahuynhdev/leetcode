package main

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{"example 1: target appears twice", []int{1, 2, 5, 2, 3}, 2, []int{1, 2}},
		{"example 2: target appears once in middle", []int{1, 2, 5, 2, 3}, 3, []int{3}},
		{"example 3: target is the maximum", []int{1, 2, 5, 2, 3}, 5, []int{4}},
		{"edge case: target not present", []int{1, 2, 5, 2, 3}, 4, []int{}},
		{"edge case: target larger than all", []int{1, 2, 3}, 100, []int{}},
		{"edge case: target smaller than all", []int{5, 6, 7}, 1, []int{}},
		{"edge case: all elements equal target", []int{7, 7, 7, 7}, 7, []int{0, 1, 2, 3}},
		{"edge case: single element match", []int{42}, 42, []int{0}},
		{"edge case: single element no match", []int{42}, 41, []int{}},
		{"edge case: target is the minimum", []int{9, 3, 1, 1, 8}, 1, []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := targetIndices(tt.nums, tt.target)
			if result == nil {
				t.Fatalf("got nil, want %v", tt.expected)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
