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

// https://leetcode.com/problems/count-good-nodes-in-binary-tree/?envType=study-plan-v2&envId=leetcode-75
// Intuition: keep track of state or path to the current node
// after a while, I think I only need to keep track of the maximum value in the path of the current node
// if max > current -> current is not a good node
// HACK: ->>>> we can traverse and maintain some state in while traversing the tree
func goodNodes(root *TreeNode) int {
	ans := 1
	var traverse func(*TreeNode, int)
	traverse = func(node *TreeNode, max int) {
		if node == nil {
			return
		}
		if node.Val >= max {
			ans++
			traverse(node.Left, node.Val)
			traverse(node.Right, node.Val)
		} else {
			traverse(node.Left, max)
			traverse(node.Right, max)
		}
	}
	traverse(root.Left, root.Val)
	traverse(root.Right, root.Val)
	return ans
}
