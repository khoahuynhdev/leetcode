package main

import (
	"reflect"
	"testing"
)

// buildList turns a slice of values into a singly-linked list and returns its head.
func buildList(vals []int) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for _, v := range vals {
		tail.Next = &ListNode{Val: v}
		tail = tail.Next
	}
	return dummy.Next
}

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		head     []int
		expected []int
	}{
		{"example 1: no critical points in a two node list", []int{3, 1}, []int{-1, -1}},
		{"example 2: three critical points", []int{5, 3, 1, 2, 5, 1, 2}, []int{1, 3}},
		{"example 3: plateaus are not critical points", []int{1, 3, 2, 2, 3, 2, 2, 2, 7}, []int{3, 3}},
		{"edge case: exactly one critical point", []int{2, 1, 3}, []int{-1, -1}},
		{"edge case: two equal nodes", []int{1, 1}, []int{-1, -1}},
		{"edge case: all values equal", []int{5, 5, 5, 5, 5}, []int{-1, -1}},
		{"edge case: strictly increasing list", []int{1, 2, 3, 4, 5}, []int{-1, -1}},
		{"edge case: strictly alternating list", []int{1, 5, 1, 5, 1, 5, 1}, []int{1, 4}},
		{"edge case: adjacent critical points", []int{1, 2, 1, 2, 1}, []int{1, 2}},
		{"edge case: closest pair is not the first pair", []int{1, 9, 1, 1, 1, 9, 8, 9, 1}, []int{1, 6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := nodesBetweenCriticalPoints(buildList(tt.head))
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("nodesBetweenCriticalPoints(%v) = %v, want %v", tt.head, got, tt.expected)
			}
		})
	}
}
