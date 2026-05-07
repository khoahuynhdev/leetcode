# 3660. Jump Game IX

[LeetCode Link](https://leetcode.com/problems/jump-game-ix/)

Difficulty: Medium
Topics: Array, Dynamic Programming
Acceptance Rate: 29.7%

## Hints

### Hint 1

A naive graph traversal (BFS/DFS) starting from every index would be too slow with `n` up to `10^5`. Step back and look at the structure of the two jump rules. Even though they are described as "forward" and "backward" jumps with different directions, they encode the *same* relationship between two indices. Try to characterize, for two indices `i < j`, exactly when an edge exists, and notice that the resulting reachability is symmetric.

### Hint 2

Once you see the underlying graph is undirected, the answer for index `i` is just the maximum value in its connected component. So the task becomes: identify the connected components efficiently. Try a few examples and convince yourself that each component is a *contiguous range of indices*, and that the per-component maxima are non-decreasing as you move left to right. The non-trivial part is figuring out where one range ends and the next begins.

### Hint 3

Sweep left to right with a stack of segments, each tagged with its maximum value, kept non-decreasing from the bottom of the stack to the top. When a new value `v` arrives, *every* segment on top whose max strictly exceeds `v` must merge with `v` into a single segment — that segment contains an index whose value is greater than `v`, which creates a backward-jump edge to `v`. After all such merges, push the combined segment back. The popping is amortized O(1) per element, so the whole sweep is linear.

## Approach

Reframe the problem as a graph problem. For two indices `i < j`, the forward-jump rule allows `i → j` iff `nums[j] < nums[i]`, and the backward-jump rule allows `j → i` iff `nums[i] > nums[j]` — the *same* condition. So the directed jump graph is actually undirected, and the answer for index `i` is the maximum value in its connected component.

Two more observations make the solution clean:

1. Each connected component is a *contiguous range* of indices.
2. The maxima of these ranges are *non-decreasing* from left to right (otherwise an earlier max would have a backward-jump edge into the later component, merging them).

We compute the partition with a single left-to-right pass using a monotonic stack of `(segmentMax, segmentStart)` entries, with `segmentMax` non-decreasing from the bottom of the stack to the top. For each `nums[i]`:

1. While the top segment's max is **strictly greater** than `nums[i]`, pop it. The popped segment must merge with `i` because its max is some earlier index whose value exceeds `nums[i]`, creating a backward-jump edge. Track the largest popped max and the earliest popped start.
2. Push a single segment whose max is the larger of the popped maxima and `nums[i]`, with the earliest popped start (or `i` itself if nothing was popped).

When the sweep finishes, each stack segment covers indices `[start, nextStart)` and the answer for every index in that range is the segment's max.

Walk through Example 2, `nums = [2, 3, 1]`:

- `i = 0`: stack empty, push `(2, 0)` → stack `[(2,0)]`.
- `i = 1`: `nums[1] = 3` is not less than the top max `2`, push `(3, 1)` → stack `[(2,0), (3,1)]`.
- `i = 2`: `nums[2] = 1`. Top max `3 > 1`, pop. Next top max `2 > 1`, pop. Push `(3, 0)` → stack `[(3, 0)]`.

The single final segment covers `[0, 3)` with max `3`, so `ans = [3, 3, 3]`.

## Complexity Analysis

Time Complexity: O(n). Each index is pushed onto the stack exactly once and popped at most once, so the total work across the inner pop-loop is amortized O(n). Filling the answer array is also O(n).

Space Complexity: O(n) for the stack in the worst case (a strictly increasing array) and for the answer array.

## Edge Cases

- **Single element** (`n == 1`): the answer is `[nums[0]]`. The algorithm handles this naturally — one segment, never popped.
- **Strictly increasing array** (e.g., `[1, 2, 3]`): no element is smaller than any earlier max, so the stack only grows. Every index is its own component, and `ans == nums`.
- **Strictly decreasing array** (e.g., `[5, 4, 3, 2, 1]`): every new element pops everything before it, leaving one segment whose max is the first element. The answer is the array filled with that max.
- **Repeated values** (e.g., `[5, 5, 5]`): equal values do not create edges (the rules require strict inequality), so each repeated value starts its own segment. Be careful to use `>` (not `>=`) as the popping condition.
- **Plateau then dip** (e.g., `[3, 1, 4, 2]`): the dip after a new running maximum bridges everything back into one component — the trickiest case. Popping while the top max strictly exceeds the new value is exactly what handles this correctly.
