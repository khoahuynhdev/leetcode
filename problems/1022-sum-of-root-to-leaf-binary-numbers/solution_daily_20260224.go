package main

// Approach: DFS traversal carrying a running value that accumulates the
// binary number from root to leaf. At each node, shift the current value
// left by 1 (multiply by 2) and add the node's bit. At a leaf, return
// the accumulated value. Sum results from all root-to-leaf paths.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(node *TreeNode, currentVal int) int {
	if node == nil {
		return 0
	}
	currentVal = currentVal*2 + node.Val
	if node.Left == nil && node.Right == nil {
		return currentVal
	}
	return dfs(node.Left, currentVal) + dfs(node.Right, currentVal)
}
