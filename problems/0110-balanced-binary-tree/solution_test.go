package main

import "testing"

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected bool
	}{
		{
			name: "example 1: balanced tree",
			root: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			expected: true,
		},
		{
			name: "example 2: unbalanced tree with deep left branch",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 4},
						Right: &TreeNode{Val: 4},
					},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{Val: 2},
			},
			expected: false,
		},
		{
			name:     "example 3: empty tree",
			root:     nil,
			expected: true,
		},
		{
			name: "edge case: single node",
			root: &TreeNode{Val: 1},
			expected: true,
		},
		{
			name: "edge case: left-skewed tree with 2 nodes",
			root: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
			},
			expected: true,
		},
		{
			name: "edge case: left-skewed tree with 3 nodes (unbalanced)",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 3},
				},
			},
			expected: false,
		},
		{
			name: "edge case: perfect binary tree",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
				Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isBalanced(tt.root)
			if result != tt.expected {
				t.Errorf("isBalanced() = %v, want %v", result, tt.expected)
			}
		})
	}
}
