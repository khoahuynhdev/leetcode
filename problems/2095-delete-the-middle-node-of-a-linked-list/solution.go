package main

// 2095. Delete the Middle Node of a Linked List
//
// Approach: fast/slow two pointers (tortoise and hare).
// The fast pointer moves two nodes per step while slow moves one, so when fast
// reaches the end, slow lands on the middle node (floor(n/2), 0-based). We keep
// a `prev` pointer trailing slow so we can splice the middle node out in O(1)
// once it is found. Single pass, constant extra space.

// ListNode is a node in a singly linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	// 0 or 1 node: the middle is the head itself; deleting it empties the list.
	if head == nil || head.Next == nil {
		return nil
	}

	var prev *ListNode
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// prev is the node just before the middle; skip over the middle node.
	prev.Next = slow.Next
	return head
}
