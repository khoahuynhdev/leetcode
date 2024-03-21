package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 type ListNode struct {
      Val int
      Next *ListNode
}

func Reverse(cur *ListNode, pre *ListNode) *ListNode {
  if cur == nil { return nil }
  node := cur.Next
  cur.Next = pre
  if node == nil { return cur}
  return Reverse(node, cur)
}
func reverseList(head *ListNode) *ListNode {
  node:=Reverse(head, nil)
  return node
}
