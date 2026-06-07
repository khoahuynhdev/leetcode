package main

// Create Binary Tree From Descriptions
//
// Approach: hash map from value -> *TreeNode. Each description is an edge
// (parent, child, isLeft). We lazily create a single node per value and reuse it
// so every reference points to the same object. We track which values appear as
// children; the one value that never appears as a child is the root.
//
// Time:  O(n) over the descriptions.
// Space: O(n) for the node map and children set.

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	nodes := make(map[int]*TreeNode)
	children := make(map[int]bool)

	// getNode returns the node for val, creating it on first use.
	getNode := func(val int) *TreeNode {
		if n, ok := nodes[val]; ok {
			return n
		}
		n := &TreeNode{Val: val}
		nodes[val] = n
		return n
	}

	for _, d := range descriptions {
		parentVal, childVal, isLeft := d[0], d[1], d[2]
		parent := getNode(parentVal)
		child := getNode(childVal)
		if isLeft == 1 {
			parent.Left = child
		} else {
			parent.Right = child
		}
		children[childVal] = true
	}

	// The root is the only value that is never a child.
	for val, node := range nodes {
		if !children[val] {
			return node
		}
	}
	return nil
}
