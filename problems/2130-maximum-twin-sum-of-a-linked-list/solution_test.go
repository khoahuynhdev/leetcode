package main

import "testing"

// buildList constructs a linked list from a slice of values and returns its head.
func buildList(vals []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, v := range vals {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return dummy.Next
}

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"example 1: [5,4,2,1] outer pairs all sum to 6", []int{5, 4, 2, 1}, 6},
		{"example 2: [4,2,2,3] max of 7 and 4", []int{4, 2, 2, 3}, 7},
		{"example 3: [1,100000] single twin pair", []int{1, 100000}, 100001},
		{"edge case: minimum size two equal nodes", []int{7, 7}, 14},
		{"edge case: all equal values", []int{5, 5, 5, 5}, 10},
		{"edge case: max sum from inner pair", []int{1, 100, 100, 1}, 200},
		{"edge case: longer list six nodes", []int{1, 2, 3, 4, 5, 6}, 7},
		{"edge case: descending values", []int{6, 5, 4, 3, 2, 1}, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildList(tt.input)
			result := pairSum(head)
			if result != tt.expected {
				t.Errorf("pairSum(%v) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}
