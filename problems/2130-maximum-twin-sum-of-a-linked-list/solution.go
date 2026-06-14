package main

// Maximum Twin Sum of a Linked List
//
// Approach: find the middle of the list with slow/fast pointers, reverse the
// second half in place, then walk both halves in lockstep. Each step pairs
// node i with node n-1-i (its twin); track the maximum pair sum.
// Time: O(n), Space: O(1).

// ListNode is a node in a singly linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	// Step 1: find the start of the second half (index n/2).
	// For an even-length list, when fast reaches nil, slow is at node n/2.
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Step 2: reverse the second half starting at slow.
	var prev *ListNode
	curr := slow
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	// Step 3: walk the first half (from head) and the reversed second half
	// (from prev) together, summing each twin pair and keeping the max.
	maxSum := 0
	first, second := head, prev
	for second != nil {
		if sum := first.Val + second.Val; sum > maxSum {
			maxSum = sum
		}
		first = first.Next
		second = second.Next
	}
	return maxSum
}
