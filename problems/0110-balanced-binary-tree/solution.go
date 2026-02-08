package main

// Approach: Bottom-up DFS that computes height and checks balance simultaneously.
// Returns -1 to signal an imbalanced subtree, which allows early termination.
// Time: O(n), Space: O(h) for recursion stack.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return height(root) != -1
}

// height returns the height of the tree if balanced, -1 if not balanced
func height(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftHeight := height(node.Left)
	if leftHeight == -1 {
		return -1
	}

	rightHeight := height(node.Right)
	if rightHeight == -1 {
		return -1
	}

	// Check if current node is balanced
	if abs(leftHeight-rightHeight) > 1 {
		return -1
	}

	// Return height of current node
	return max(leftHeight, rightHeight) + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
