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
func sumNumbers(root *TreeNode) int {
   return SumRoot(root,0)
}

func SumRoot(root *TreeNode, s int) int {
  if root == nil { return 0}
  if root.Left == nil && root.Right == nil {
    return s + root.Val
  }
  ms := (s + root.Val) * 10
  l,r := SumRoot(root.Left, ms), SumRoot(root.Right, ms)
  return l + r
}
