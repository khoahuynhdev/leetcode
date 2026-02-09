package main

import (
	"math"
	"testing"
)

func TestBalanceBST(t *testing.T) {
	tests := []struct {
		name  string
		input *TreeNode
	}{
		{
			name:  "example 1: skewed right tree",
			input: buildTree([]interface{}{1, nil, 2, nil, 3, nil, 4}),
		},
		{
			name:  "example 2: already balanced tree",
			input: buildTree([]interface{}{2, 1, 3}),
		},
		{
			name:  "edge case: single node",
			input: &TreeNode{Val: 5},
		},
		{
			name:  "edge case: completely left-skewed tree",
			input: buildSkewedLeft([]int{5, 4, 3, 2, 1}),
		},
		{
			name:  "edge case: larger unbalanced tree",
			input: buildTree([]interface{}{1, nil, 2, nil, 3, nil, 4, nil, 5, nil, 6}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := balanceBST(tt.input)

			// Verify result is a valid BST
			if !isBST(result, math.MinInt64, math.MaxInt64) {
				t.Errorf("result is not a valid BST")
			}

			// Verify result is balanced
			if !isBalanced(result) {
				t.Errorf("result tree is not balanced")
			}

			// Verify same number of nodes
			inputSize := countNodes(tt.input)
			resultSize := countNodes(result)
			if inputSize != resultSize {
				t.Errorf("node count mismatch: input has %d nodes, result has %d nodes", inputSize, resultSize)
			}

			// Verify same values (via in-order traversal)
			inputValues := inorderValues(tt.input)
			resultValues := inorderValues(result)
			if !equalSlices(inputValues, resultValues) {
				t.Errorf("values mismatch: input has %v, result has %v", inputValues, resultValues)
			}
		})
	}
}

// Helper: build tree from level-order array (nil represents missing nodes)
func buildTree(values []interface{}) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	root := &TreeNode{Val: values[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(values) {
		node := queue[0]
		queue = queue[1:]

		// Left child
		if i < len(values) && values[i] != nil {
			node.Left = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		// Right child
		if i < len(values) && values[i] != nil {
			node.Right = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

// Helper: build left-skewed tree from descending values
func buildSkewedLeft(values []int) *TreeNode {
	if len(values) == 0 {
		return nil
	}
	root := &TreeNode{Val: values[0]}
	current := root
	for i := 1; i < len(values); i++ {
		current.Left = &TreeNode{Val: values[i]}
		current = current.Left
	}
	return root
}

// Helper: check if tree is a valid BST
func isBST(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	if node.Val <= min || node.Val >= max {
		return false
	}
	return isBST(node.Left, min, node.Val) && isBST(node.Right, node.Val, max)
}

// Helper: check if tree is balanced
func isBalanced(node *TreeNode) bool {
	_, balanced := checkBalance(node)
	return balanced
}

func checkBalance(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}

	leftHeight, leftBalanced := checkBalance(node.Left)
	if !leftBalanced {
		return 0, false
	}

	rightHeight, rightBalanced := checkBalance(node.Right)
	if !rightBalanced {
		return 0, false
	}

	if abs(leftHeight-rightHeight) > 1 {
		return 0, false
	}

	return max(leftHeight, rightHeight) + 1, true
}

// Helper: count nodes in tree
func countNodes(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + countNodes(node.Left) + countNodes(node.Right)
}

// Helper: get in-order values
func inorderValues(node *TreeNode) []int {
	var values []int
	inorder(node, &values)
	return values
}

// Helper: compare two slices
func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
