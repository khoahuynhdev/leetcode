package solution

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// SOLUTION: this is a complete binary tree
// the node of a complete binary tree should be (2^h - 1) if the node is perfect binary tree
// otherwise recursively incremental count nodes
func leftHeight(node *TreeNode) int {
	cnt := 0
	for node != nil {
		cnt++
		node = node.Left
	}
	return cnt
}

func rightHeight(node *TreeNode) int {
	cnt := 0
	for node != nil {
		cnt++
		node = node.Right
	}
	return cnt
}

func countNodes(root *TreeNode) int {
	// we can DFS/BFS
	// lower than O(n) -> O(logn)

	if root == nil {
		return 0
	}

	left := leftHeight(root)
	right := rightHeight(root)

	if left == right {
		return (1 << left) - 1
	}

	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
