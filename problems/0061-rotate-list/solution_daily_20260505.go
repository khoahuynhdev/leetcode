package main

// Rotate List — rotate the list to the right by k places.
//
// Approach: measure the list length n and grab the tail in one pass,
// reduce k modulo n, splice the tail to the head to form a cycle, walk
// n-k-1 steps from head to land on the new tail, then break the cycle.
// O(n) time, O(1) space.

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	n := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		n++
	}

	k %= n
	if k == 0 {
		return head
	}

	tail.Next = head

	newTail := head
	for i := 0; i < n-k-1; i++ {
		newTail = newTail.Next
	}

	newHead := newTail.Next
	newTail.Next = nil
	return newHead
}
