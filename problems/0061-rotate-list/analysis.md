# 0061. Rotate List

[LeetCode Link](https://leetcode.com/problems/rotate-list/)

Difficulty: Medium
Topics: Linked List, Two Pointers
Acceptance Rate: 41.8%

## Hints

### Hint 1

This is a linked list manipulation problem. Before reaching for clever pointer tricks, ask yourself: what does "rotating right by k" actually do to the structure of the list? Try drawing a small example on paper and pay attention to which node ends up as the new head and which node ends up as the new tail.

### Hint 2

Notice that `k` can be much larger than the list length (up to 2×10^9). Rotating by the list length yields the same list, so the *effective* number of rotations is `k % n`, where `n` is the length. That means a single pass to find the length is essentially free, and after it the work is bounded by `n`.

### Hint 3

The key insight: rotating right by `k'` (where `k' = k % n`) is the same as breaking the list at the `(n - k' - 1)`-th node (0-indexed) and reattaching the back half in front of the front half. A clean trick is to first connect the tail to the head, forming a cycle, then walk `n - k'` steps from the head and break the cycle at that point — the node you stop at becomes the new tail, and its `Next` becomes the new head.

## Approach

We use a two-pass approach over the linked list.

1. **Measure the list.** Walk from `head` to the last node, counting nodes (`n`) and remembering the tail.
2. **Reduce `k`.** Replace `k` with `k % n`. If the result is `0`, the list is unchanged — return `head` as is.
3. **Form a cycle.** Point the tail's `Next` back to `head`. The list is now circular.
4. **Find the new tail.** Walk `n - k - 1` steps from `head`. The node we land on is the new tail of the rotated list.
5. **Break the cycle.** The new head is `newTail.Next`. Set `newTail.Next = nil` to terminate the rotated list, and return the new head.

**Worked example:** `[1,2,3,4,5]`, `k = 2`, `n = 5`, effective `k = 2`. Connect 5 → 1 to form a cycle. Walk `5 - 2 - 1 = 2` steps from head: `1 → 2 → 3`, so the new tail is the node holding `3`. The new head is `3.Next = 4`, and we cut after `3`, yielding `4 → 5 → 1 → 2 → 3`.

This works because in a circular list, every node has a successor, and the rotation is just a choice of where to "break" that circle.

## Complexity Analysis

Time Complexity: O(n) — one pass to measure, one partial pass to find the new tail.
Space Complexity: O(1) — only a constant number of pointers, regardless of input size.

## Edge Cases

- **Empty list (`head == nil`)**: Nothing to rotate; return `nil`. Always check this first to avoid dereferencing a nil pointer when measuring length.
- **Single node**: Any rotation returns the same node. The early-return check handles this naturally.
- **`k == 0`**: No rotation is needed; return `head` immediately to avoid wasted work and an unnecessary cycle/break dance.
- **`k` is a multiple of `n`**: After `k % n`, effective `k` is `0` — the list is unchanged. Don't forget the second early-return *after* taking the modulus.
- **`k > n`**: The modulus reduction keeps the algorithm O(n) even when `k` is up to 2×10^9; without it, a naive rotate-by-one-`k`-times approach would time out.
