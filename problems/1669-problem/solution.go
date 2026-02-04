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
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
  db, p1, p2,cnt := make([]*ListNode, 2), list1, list2,0
  for p1 != nil {
    if cnt == a - 1 {
      db[0] = p1
    }
    if cnt == b {
      db[1] = p1
      break
    }
    p1 = p1.Next
    cnt++
  }
  for p2.Next != nil {
    p2 = p2.Next
  }
  // fmt.Println(db)
  db[0].Next = list2
  p2.Next = db[1].Next
  return list1
}
