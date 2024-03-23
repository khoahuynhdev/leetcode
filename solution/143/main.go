package main

type ListNode struct {
     Val int
     Next *ListNode
}
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    for head != nil {
        memoNext := head.Next
        head.Next = prev
        prev = head
        head = memoNext
    }
    return prev
}

func reorderList(head *ListNode)  {
    res := &ListNode{Val: 0, Next: head}

    slow, fast := res, res
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    fast = slow.Next
    slow.Next = nil

    left := head
    right := reverseList(fast)
    prev := res

    for left != nil || right != nil {
        prev.Next = left
        prev = prev.Next
        left = left.Next

        if right != nil {
            prev.Next = right
            prev = prev.Next
            right = right.Next
        }
    }

    head = res.Next
}
