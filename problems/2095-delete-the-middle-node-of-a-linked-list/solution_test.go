package main

import (
	"reflect"
	"testing"
)

// buildList constructs a linked list from a slice and returns its head.
func buildList(vals []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

// toSlice converts a linked list back into a slice for comparison.
func toSlice(head *ListNode) []int {
	var out []int
	for node := head; node != nil; node = node.Next {
		out = append(out, node.Val)
	}
	return out
}

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"example 1: odd length, removes index 3", []int{1, 3, 4, 7, 1, 2, 6}, []int{1, 3, 4, 1, 2, 6}},
		{"example 2: even length, removes index 2", []int{1, 2, 3, 4}, []int{1, 2, 4}},
		{"example 3: two nodes, removes second", []int{2, 1}, []int{2}},
		{"edge case: single node becomes empty", []int{42}, nil},
		{"edge case: three nodes removes middle", []int{10, 20, 30}, []int{10, 30}},
		{"edge case: five nodes removes index 2", []int{5, 4, 3, 2, 1}, []int{5, 4, 2, 1}},
		{"edge case: empty input", nil, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toSlice(deleteMiddle(buildList(tt.input)))
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("deleteMiddle(%v) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
