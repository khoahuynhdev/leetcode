package solution

import (
	"math"
)

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// max root = 1 + max(max(left) + max(right))
	return 1 + int(math.Max(float64(MaxDepth(root.Left)), float64(MaxDepth(root.Right))))
}
