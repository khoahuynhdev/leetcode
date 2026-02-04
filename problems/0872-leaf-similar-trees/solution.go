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

func inorderTraverse(root *TreeNode, arr []int) []int {
	if root == nil {
		return arr
	}
	if root.Left == nil && root.Right == nil {
		arr = append(arr, root.Val)
		return arr
	}

	lArr := inorderTraverse(root.Left, arr)
	rArr := inorderTraverse(root.Right, arr)
	return append(lArr, rArr...)
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	// h1: leaf sequence value should not contain other values ✅
	arr1 := inorderTraverse(root1, []int{})
	arr2 := inorderTraverse(root2, []int{})
	if len(arr1) != len(arr2) {
		return false
	}
	for idx, v := range arr1 {
		if arr2[idx] != v {
			return false
		}
	}
	return true
}
