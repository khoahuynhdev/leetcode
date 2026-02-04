package solution

/**
 * Definition for a binary tree node.

 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// simple searchBST
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	left := searchBST(root.Left, val)
	if left != nil && left.Val == val {
		return left
	}
	right := searchBST(root.Right, val)
	if right != nil && right.Val == val {
		return right
	}
	return nil
}
