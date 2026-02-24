package main

import "testing"

func TestSumRootToLeaf(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected int
	}{
		{
			name: "example 1: complete binary tree [1,0,1,0,1,0,1]",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   0,
					Left:  &TreeNode{Val: 0},
					Right: &TreeNode{Val: 1},
				},
				Right: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 0},
					Right: &TreeNode{Val: 1},
				},
			},
			expected: 22,
		},
		{
			name:     "example 2: single node zero",
			root:     &TreeNode{Val: 0},
			expected: 0,
		},
		{
			name:     "edge case: single node one",
			root:     &TreeNode{Val: 1},
			expected: 1,
		},
		{
			name: "edge case: left-skewed tree 1->0->1",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  0,
					Left: &TreeNode{Val: 1},
				},
			},
			expected: 5, // binary 101 = 5
		},
		{
			name: "edge case: all zeros",
			root: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val:  0,
					Left: &TreeNode{Val: 0},
				},
				Right: &TreeNode{
					Val: 0,
				},
			},
			expected: 0, // 000 + 00 = 0
		},
		{
			name: "edge case: right-skewed tree 1->1->1",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:   1,
					Right: &TreeNode{Val: 1},
				},
			},
			expected: 7, // binary 111 = 7
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sumRootToLeaf(tt.root)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
