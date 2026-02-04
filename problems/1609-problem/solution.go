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
func Walk(t *TreeNode, deg int, db map[int][]int) {
  if t == nil { return }
  if db[deg] == nil {
    db[deg] = []int{t.Val}
  } else {
    db[deg] = append(db[deg], t.Val)
  }
  Walk(t.Left, deg+1, db)
  Walk(t.Right, deg+1, db)
}

func isEvenOddTree(root *TreeNode) bool {
  if root.Val % 2 == 0 { return false }
  db := make(map[int][]int)
  Walk(root, 0, db)
  for k,v := range db {
  if k % 2 == 0 {
      if len(v) == 1 && v[0] % 2 == 0 { return false }
      for i:=0;i<len(v);i++ {
        if v[i] % 2 == 0 { return false}
        if i == len(v) - 1 { break }
        if v[i] >= v[i+1] { return false}
      }
    } else {
      // fmt.Println(k,v)
      if len(v) == 1 && v[0] % 2 != 0 { return false }
      for i:=0;i<len(v);i++ {
        if v[i] % 2 != 0 { return false}
        if i == len(v) - 1 { break }
        if v[i] <= v[i+1] { return false}
      }
    }
  }
  return true
}
