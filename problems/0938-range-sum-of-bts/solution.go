package solution

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// intuition: simple, just traverse the tree
// can be optimize since it's bst
// TODO: optimize solution
func rangeSumBST(root *TreeNode, low int, high int) int {
	ans := 0
	ans = traverse(root, low, high)
	return ans
}

func traverse(root *TreeNode, low, high int) int {
	if root == nil {
		return 0
	}
	if root.Val >= low && root.Val <= high {
		return root.Val + traverse(root.Left, low, high) + traverse(root.Right, low, high)
	}
	return traverse(root.Left, low, high) + traverse(root.Right, low, high)
}
