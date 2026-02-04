package main

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
func sumOfLeftLeaves(root *TreeNode) int {
  return SumLeave(root, false)
}

func SumLeave(root *TreeNode, left bool) int {
  if root == nil { return 0 }
  if root.Left == nil && root.Right == nil {
    if left { return root.Val} else { return 0 }
  }
  l, r := SumLeave(root.Left, true), SumLeave(root.Right, false)
  return  l + r
}
