package main

// Good question about pointer
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
 }
 
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
  if depth == 1 {
    r := &TreeNode{ Val: val, Left: root}
    return r
  }
  addDepth(root, 1, depth, val)
  return root
}

func addDepth(root *TreeNode, depth, target, val int) {
  if root == nil { return }
  if depth == target - 1 {
    l, r := root.Left, root.Right
    nl, nr := &TreeNode{ Val: val, Left: l}, &TreeNode{ Val: val, Right: r}
    root.Left, root.Right = nl, nr
    return
  } else {
    addDepth(root.Left, depth + 1, target, val)
    addDepth(root.Right, depth + 1, target, val)
  }
}
