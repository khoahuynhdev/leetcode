package main

import "testing"

// intPtr wraps a value so test trees can spell out LeetCode's null placeholders.
func intPtr(v int) *int { return &v }

// buildTreeLevelOrder builds a binary tree from a LeetCode-style level-order
// slice, where nil entries mark absent children.
func buildTreeLevelOrder(vals []*int) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}

	root := &TreeNode{Val: *vals[0]}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]

		if i < len(vals) {
			if vals[i] != nil {
				node.Left = &TreeNode{Val: *vals[i]}
				queue = append(queue, node.Left)
			}
			i++
		}
		if i < len(vals) {
			if vals[i] != nil {
				node.Right = &TreeNode{Val: *vals[i]}
				queue = append(queue, node.Right)
			}
			i++
		}
	}

	return root
}

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		vals     []*int
		expected int
	}{
		{
			name:     "example 1: [4,8,5,0,1,null,6] counts root, 5, 0, 1 and 6",
			vals:     []*int{intPtr(4), intPtr(8), intPtr(5), intPtr(0), intPtr(1), nil, intPtr(6)},
			expected: 5,
		},
		{
			name:     "example 2: single node [1] always matches itself",
			vals:     []*int{intPtr(1)},
			expected: 1,
		},
		{
			name:     "edge case: empty input returns zero",
			vals:     nil,
			expected: 0,
		},
		{
			name:     "edge case: all zeros [0,0,0] counts every node",
			vals:     []*int{intPtr(0), intPtr(0), intPtr(0)},
			expected: 3,
		},
		{
			name:     "edge case: flooring makes [5,null,6] match at the root",
			vals:     []*int{intPtr(5), nil, intPtr(6)},
			expected: 2,
		},
		{
			name:     "edge case: left-skewed chain [3,2,null,1] counts only the leaf",
			vals:     []*int{intPtr(3), intPtr(2), nil, intPtr(1)},
			expected: 1,
		},
		{
			name:     "edge case: [1,2,3] root average of 2 does not match value 1",
			vals:     []*int{intPtr(1), intPtr(2), intPtr(3)},
			expected: 2,
		},
		{
			name:     "edge case: max constrained values [1000,1000,1000] all match",
			vals:     []*int{intPtr(1000), intPtr(1000), intPtr(1000)},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTreeLevelOrder(tt.vals)
			if result := averageOfSubtree(root); result != tt.expected {
				t.Errorf("averageOfSubtree() = %v, want %v", result, tt.expected)
			}
		})
	}
}
