package main

import (
	"sort"
	"testing"
)

// levelOrder serializes a binary tree into the LeetCode-style level-order slice,
// using nil to mark absent children (trailing nils trimmed), so trees can be
// compared structurally.
func levelOrder(root *TreeNode) []interface{} {
	var out []interface{}
	if root == nil {
		return out
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			out = append(out, nil)
			continue
		}
		out = append(out, node.Val)
		queue = append(queue, node.Left, node.Right)
	}
	// Trim trailing nils to match LeetCode's compact representation.
	for len(out) > 0 && out[len(out)-1] == nil {
		out = out[:len(out)-1]
	}
	return out
}

func equalSlices(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// collectVals gathers all node values in the tree (order independent) so we can
// confirm structural integrity beyond just the root.
func collectVals(root *TreeNode) []int {
	var vals []int
	var dfs func(*TreeNode)
	dfs = func(n *TreeNode) {
		if n == nil {
			return
		}
		vals = append(vals, n.Val)
		dfs(n.Left)
		dfs(n.Right)
	}
	dfs(root)
	sort.Ints(vals)
	return vals
}

func TestSolution(t *testing.T) {
	tests := []struct {
		name         string
		descriptions [][]int
		wantLevel    []interface{} // expected level-order serialization
		wantVals     []int         // expected sorted set of values
	}{
		{
			name:         "example 1: standard tree rooted at 50",
			descriptions: [][]int{{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}, {80, 19, 1}},
			wantLevel:    []interface{}{50, 20, 80, 15, 17, 19},
			wantVals:     []int{15, 17, 19, 20, 50, 80},
		},
		{
			name:         "example 2: left-leaning chain rooted at 1",
			descriptions: [][]int{{1, 2, 1}, {2, 3, 0}, {3, 4, 1}},
			wantLevel:    []interface{}{1, 2, nil, nil, 3, 4},
			wantVals:     []int{1, 2, 3, 4},
		},
		{
			name:         "edge case: single edge, parent is root with one left child",
			descriptions: [][]int{{1, 2, 1}},
			wantLevel:    []interface{}{1, 2},
			wantVals:     []int{1, 2},
		},
		{
			name:         "edge case: single edge, right child only",
			descriptions: [][]int{{10, 5, 0}},
			wantLevel:    []interface{}{10, nil, 5},
			wantVals:     []int{5, 10},
		},
		{
			name:         "edge case: descriptions out of order (child seen before its own edge)",
			descriptions: [][]int{{80, 19, 1}, {50, 80, 0}, {50, 20, 1}, {20, 15, 1}, {20, 17, 0}},
			wantLevel:    []interface{}{50, 20, 80, 15, 17, 19},
			wantVals:     []int{15, 17, 19, 20, 50, 80},
		},
		{
			name:         "edge case: right-leaning chain",
			descriptions: [][]int{{1, 2, 0}, {2, 3, 0}},
			wantLevel:    []interface{}{1, nil, 2, nil, 3},
			wantVals:     []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := createBinaryTree(tt.descriptions)
			got := levelOrder(root)
			if !equalSlices(got, tt.wantLevel) {
				t.Errorf("level order = %v, want %v", got, tt.wantLevel)
			}
			if gotVals := collectVals(root); !equalIntSlices(gotVals, tt.wantVals) {
				t.Errorf("values = %v, want %v", gotVals, tt.wantVals)
			}
		})
	}
}

func equalIntSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
