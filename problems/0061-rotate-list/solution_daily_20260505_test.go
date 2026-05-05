package main

import (
	"reflect"
	"testing"
)

func buildListDaily20260505(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	cur := head
	for _, v := range vals[1:] {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return head
}

func toSliceDaily20260505(head *ListNode) []int {
	result := []int{}
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func TestRotateRightDaily20260505(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		k        int
		expected []int
	}{
		{"example 1: rotate [1,2,3,4,5] by 2", []int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
		{"example 2: rotate [0,1,2] by 4 (k > n)", []int{0, 1, 2}, 4, []int{2, 0, 1}},
		{"edge case: empty list", []int{}, 3, []int{}},
		{"edge case: single node, k > 0", []int{1}, 5, []int{1}},
		{"edge case: k is zero", []int{1, 2, 3}, 0, []int{1, 2, 3}},
		{"edge case: k equals length (no-op)", []int{1, 2, 3, 4}, 4, []int{1, 2, 3, 4}},
		{"edge case: k is exact multiple of length", []int{1, 2, 3}, 6, []int{1, 2, 3}},
		{"edge case: two nodes rotated by 1", []int{1, 2}, 1, []int{2, 1}},
		{"edge case: large k (modulus reduction)", []int{1, 2, 3, 4, 5}, 2000000000, []int{1, 2, 3, 4, 5}},
		{"edge case: negative values", []int{-1, -2, -3}, 1, []int{-3, -1, -2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildListDaily20260505(tt.input)
			result := rotateRight(head, tt.k)
			got := toSliceDaily20260505(result)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("rotateRight(%v, %d) = %v, want %v", tt.input, tt.k, got, tt.expected)
			}
		})
	}
}
