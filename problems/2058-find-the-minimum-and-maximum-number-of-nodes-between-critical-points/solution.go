package main

// Single pass with a sliding (prev, cur, next) window.
//
// A node is a critical point when it is strictly greater or strictly smaller
// than both of its neighbours, so detection only needs three nodes at a time.
// Once a critical point is found at index i we do not store it: the maximum
// distance is always last-first, and the minimum distance can only occur
// between two consecutive critical points (any wider pair straddles a closer
// one). Tracking first, previous and the running minimum keeps space at O(1).

// ListNode is the singly-linked list node used by LeetCode.
type ListNode struct {
	Val  int
	Next *ListNode
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
	const notFound = -1

	first, last := notFound, notFound
	minDist := notFound

	prev := head
	if prev == nil {
		return []int{-1, -1}
	}

	// cur is the candidate node; index is its 1-based position in the list.
	index := 2
	for cur := prev.Next; cur != nil && cur.Next != nil; cur, index = cur.Next, index+1 {
		next := cur.Next

		isMaxima := cur.Val > prev.Val && cur.Val > next.Val
		isMinima := cur.Val < prev.Val && cur.Val < next.Val
		prev = cur

		if !isMaxima && !isMinima {
			continue
		}

		if first == notFound {
			first = index
		} else if gap := index - last; minDist == notFound || gap < minDist {
			minDist = gap
		}
		last = index
	}

	if minDist == notFound {
		return []int{-1, -1}
	}
	return []int{minDist, last - first}
}
