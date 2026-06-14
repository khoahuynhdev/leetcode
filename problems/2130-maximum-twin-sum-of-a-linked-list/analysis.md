# 2130. Maximum Twin Sum of a Linked List

[LeetCode Link](https://leetcode.com/problems/maximum-twin-sum-of-a-linked-list/)

Difficulty: Medium
Topics: Linked List, Two Pointers, Stack
Acceptance Rate: 82.1%

## Hints

### Hint 1

A "twin" pairs node `i` with node `n-1-i`. That means you're pairing the front of the list with the back, the second element with the second-to-last, and so on — working inward from both ends. Whenever a problem asks you to pair the start with the end of a sequence, think about how you'd reach both ends at once. With an array that's trivial (`a[i] + a[n-1-i]`), but a singly linked list only lets you walk forward. How can you bridge that gap?

### Hint 2

There are two broad ways to get random/back-to-front access on a singly linked list. One is to dump the values into a structure that *does* support indexing or last-in-first-out access — an array or a **stack**. The other is to physically rearrange the list so the second half points backward. Either way, the goal is to line up node `i` with node `n-1-i` and add them. Try the simplest version first: copy all values into a slice, then use two indices marching toward the middle.

### Hint 3

The optimal trick avoids the extra storage: use the classic **slow/fast pointer** to find the middle in one pass, then **reverse the second half** in place. Now you have two list halves — the original first half (head ... mid) and the reversed second half (tail ... mid). Walk both simultaneously, one node at a time; each step gives you exactly one twin pair. Take the running maximum of the sums. This is O(n) time and O(1) extra space, since reversal just rewires the existing nodes.

## Approach

The relationship "node `i` is the twin of node `n-1-i`" is the same pairing you'd get by folding the list in half. So the whole problem reduces to: pair up the first half with the reversed second half and find the largest pair sum.

**Optimal approach — middle + reverse + two pointers (O(1) space):**

1. **Find the middle** with the slow/fast pointer technique. Start both at `head`; advance `slow` by one and `fast` by two until `fast` runs off the end. Because the length is even, when `fast == nil`, `slow` sits at the start of the second half (node `n/2`).

2. **Reverse the second half** in place starting from `slow`. Standard iterative reversal with three pointers (`prev`, `curr`, `next`). After this, `prev` is the head of the reversed second half — i.e., the original last node.

3. **Walk both halves together.** Let `first = head` and `second = prev`. At each step, `first.Val + second.Val` is one twin sum: on the first step it's node `0` + node `n-1`, then node `1` + node `n-2`, and so on. Keep a running maximum. Advance both pointers; stop when the reversed half is exhausted (`second == nil`).

Walking through Example 1, `head = [5,4,2,1]`:
- Middle search lands `slow` on the node with value `2` (index 2).
- Reversing `[2,1]` gives `[1,2]`, headed by the node with value `1`.
- Pair up: `5 + 1 = 6`, then `4 + 2 = 6`. Maximum is `6`. ✓

**Simpler alternative — stack or array (O(n) space):** Push every value onto a slice in order, then use two indices `lo = 0`, `hi = n-1` moving inward, summing `vals[lo] + vals[hi]`. Easier to write and reason about; the only cost is the extra O(n) storage. Great as a first pass before optimizing.

This problem is a clean, very approachable Medium — the 82% acceptance rate reflects that. If the in-place reversal feels fiddly, start with the array version to get a green check, then come back and tighten the space. Both are completely valid.

## Complexity Analysis

Time Complexity: O(n) — one pass to find the middle, one pass to reverse the second half, one pass to compare pairs. All linear.
Space Complexity: O(1) for the optimal reverse-in-place approach (only a handful of pointers). The array/stack alternative is O(n).

## Edge Cases

- **Minimum size (n = 2):** The smallest legal list, e.g. `[1, 100000]`. There's exactly one twin pair, so the answer is just the sum of the two nodes. Verify the slow/fast loop and reversal don't mishandle a two-node list (they shouldn't — `slow` ends on the second node, and reversing a single-node tail is a no-op).
- **All equal values:** e.g. `[5,5,5,5]`. Every twin sum is identical; make sure the max logic still returns that value rather than, say, only comparing the first pair.
- **Maximum value pairs not adjacent:** The largest twin sum may come from any pair, not necessarily the outermost. Don't short-circuit after the first pair — always scan all `n/2` pairs.
- **Guaranteed even length:** The constraints promise an even number of nodes (≥ 2), so you don't need to handle odd lengths or an empty list. Relying on this keeps the middle-finding logic simple.
