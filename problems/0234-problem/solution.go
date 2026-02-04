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
func isPalindrome(head *ListNode) bool {
  if head.Next == nil { return true}
  tail, mid := head, head
  for tail != nil && tail.Next != nil {
    mid = mid.Next
    tail = tail.Next
    if tail.Next != nil {
      tail = tail.Next
    }
  }
  if mid == tail {
    if head.Val == head.Next.Val { return true}
    return false
  }
  r := Reverse(mid.Next, mid)
  for head != mid {
    if head.Val != r.Val {
      return false
    }
    head = head.Next
    r = r.Next
  }
  return true
}
