package main

// Approach: Two-step process
// 1. Perform in-order traversal to extract sorted values from BST
// 2. Build balanced BST from sorted array by recursively choosing middle elements

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	// Step 1: Collect values in sorted order via in-order traversal
	values := []int{}
	inorder(root, &values)

	// Step 2: Build balanced BST from sorted array
	return buildBalancedBST(values, 0, len(values)-1)
}

func inorder(node *TreeNode, values *[]int) {
	if node == nil {
		return
	}
	inorder(node.Left, values)
	*values = append(*values, node.Val)
	inorder(node.Right, values)
}

func buildBalancedBST(values []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	// Choose middle element as root for balance
	mid := left + (right-left)/2
	node := &TreeNode{Val: values[mid]}

	// Recursively build left and right subtrees
	node.Left = buildBalancedBST(values, left, mid-1)
	node.Right = buildBalancedBST(values, mid+1, right)

	return node
}
