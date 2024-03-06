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

func hasCycle(head *ListNode) bool {
   sl, ft := head, head
   for ft != nil && ft.Next != nil {
     sl, ft = sl.Next, ft.Next.Next
     if sl == ft { return true}
   }
   return false
}
