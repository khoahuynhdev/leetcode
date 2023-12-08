package solution

// https://leetcode.com/problems/construct-string-from-binary-tree/
import (
	"fmt"
)

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

// Intuition: At first I was going to store the node in an array and accessing them with 2n+1, 2n+2
// but realize the '()' will come next to the right most node as the right node cannot be null
// otherwise there will be no '()' in the string
func tree2str(root *TreeNode) string {
	if root == nil {
		return "()"
	}
	if root.Left == nil && root.Right == nil {
		return fmt.Sprintf("%d", root.Val)
	}
	if root.Right == nil && root.Left != nil {
		return fmt.Sprintf("%d(%s)", root.Val, tree2str(root.Left))
	}

	if root.Left == nil && root.Right != nil {
		return fmt.Sprintf("%d()(%s)", root.Val, tree2str(root.Right))
	}
	return fmt.Sprintf("%d(%s)(%s)", root.Val, tree2str(root.Left), tree2str(root.Right))
}
