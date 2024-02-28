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

type BL struct {
  Val int
  Deg int
}

func Walk(t *TreeNode, deg int, s *BL) {
  if t == nil {
    return
  }
   Walk(t.Left, deg + 1, s)
   if s.Deg < deg {
     s.Deg = deg
     s.Val = t.Val
   }

   Walk(t.Right, deg + 1, s)
}

func findBottomLeftValue(root *TreeNode) int {
  ans := &BL{Deg: 0, Val: root.Val}
  Walk(root,0,ans)
  return ans.Val
}
