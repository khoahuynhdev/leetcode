package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil && n == 1 {
		return nil
	}
	size, ptr := 1, head
	for ptr.Next != nil {
		size++
		ptr = ptr.Next
	}
	pos := size - n + 1
	if pos == 1 {
		return head.Next
	}
	ptr, size = head, 1
	for size < pos {
		if size == pos-1 {
			//  fmt.Println(ptr)
			tmp := ptr.Next
			ptr.Next = tmp.Next
		} else {
			ptr = ptr.Next
		}
		size++
	}
	return head
}
