# 2095. Delete the Middle Node of a Linked List

[LeetCode Link](https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/)

Difficulty: Medium
Topics: Linked List, Two Pointers
Acceptance Rate: 60.0%

## Hints

### Hint 1

To delete the middle node you first need to *find* it. The middle is the `⌊n / 2⌋th` node (0-based). The naive idea is to walk the whole list once to learn `n`, then walk again to reach the middle. That works — but think about whether there is a way to locate the middle in a **single pass** without counting first. The "Two Pointers" tag is a strong nudge here.

### Hint 2

Consider two pointers that move through the list at different speeds: a `slow` pointer advancing one node per step and a `fast` pointer advancing two nodes per step. When `fast` reaches the end, where does `slow` end up? This "tortoise and hare" technique lets the two pointers naturally meet at the middle.

### Hint 3

Deleting a node in a singly linked list means rewiring its **predecessor** to skip over it (`prev.Next = node.Next`). So you don't just want a pointer *at* the middle — you want a pointer *just before* it. The classic trick is to keep `slow` one step behind by starting `fast` slightly ahead (e.g. advance `fast` by two before the loop, or track a separate `prev`). Don't forget the single-node case: deleting the only node should return an empty list (`nil`).

## Approach

We use the **fast/slow two-pointer** pattern to find the node *just before* the middle in one pass, then splice the middle node out.

Setup:
- If the list has 0 or 1 node, the middle is the head itself; deleting it yields an empty list, so return `nil` immediately.
- Otherwise initialize `slow = head`, `fast = head`, and a `prev` pointer that trails `slow`.

Traversal:
- Move `fast` forward by two nodes and `slow` forward by one node per iteration, updating `prev` to follow `slow`.
- Loop while `fast != nil && fast.Next != nil`.

When the loop ends, `slow` sits exactly on the middle node (`⌊n / 2⌋`) and `prev` is the node immediately before it. Why? The fast pointer covers twice the ground, so by the time it falls off the end, the slow pointer has covered half the list — landing on the floor-division midpoint for both odd and even lengths.

Splice:
- Set `prev.Next = slow.Next`, which unlinks the middle node.
- Return `head` (unchanged, since we never delete the head in this branch).

Example (`[1,3,4,7,1,2,6]`, n = 7, target index 3 = value 7):
- After the traversal `slow` lands on the node with value `7` and `prev` on `4`.
- `prev.Next = slow.Next` connects `4 -> 1`, producing `[1,3,4,1,2,6]`. ✓

Example (`[1,2,3,4]`, n = 4, target index 2 = value 3):
- `slow` lands on `3`, `prev` on `2`, giving `2 -> 4` and result `[1,2,4]`. ✓

This is a clean, single-pass solution and a great drill for mastering pointer rewiring.

## Complexity Analysis

Time Complexity: O(n) — one traversal where `fast` covers the list at double speed.
Space Complexity: O(1) — only a constant number of pointers are used.

## Edge Cases

- **Single node `[x]`**: The only node is the middle. Deleting it must return `nil` (empty list). Handled by the early return.
- **Two nodes `[a,b]`**: n = 2, middle index = 1, so the second node is removed, leaving `[a]`. Verifies that `prev` correctly trails and the head is preserved.
- **Odd vs. even length**: Floor division means the pointer math must land on `⌊n/2⌋` for both parities; the fast/slow setup handles both without special casing.
- **Empty list `nil`**: Although the constraints guarantee at least one node, defensively returning `nil` for a `nil` head keeps the function robust.
