# 3420. Count Non-Decreasing Subarrays After K Operations

[LeetCode Link](https://leetcode.com/problems/count-non-decreasing-subarrays-after-k-operations/)

Difficulty: Hard
Topics: Array, Stack, Segment Tree, Queue, Sliding Window, Monotonic Stack, Monotonic Queue
Acceptance Rate: 49.9%

## Hints

### Hint 1

First work out the cost of one fixed subarray: what is the fewest increments that make it non-decreasing? Since you can only increase values, each element has to go up to some value that depends on the elements before it. After that, ask whether the cost is **monotonic**. If a subarray is valid, are all of its sub-subarrays valid too? If so, a sliding window will work.

### Hint 2

The cost of `[l..r]` is `Σ (max(nums[l..i]) - nums[i])`: every element is raised to the running prefix maximum. Shrinking a valid window never makes it invalid, so two pointers work. The hard part is keeping the cost up to date. When the window's **left** end changes, the prefix maxima of many elements can change at once. Try moving the left pointer from **right to left**, so the end that controls the prefix maxima is the end where elements arrive.

### Hint 3

Group the window into blocks of elements that share the same prefix maximum. Adding a new element `x` on the left takes over every block whose max is `<= x`. Each element of such a block now costs `x - oldMax` more, so you can pop those blocks off a **monotonic stack** and charge `(x - oldMax) * blockSize`. When the cost goes over `k`, remove elements from the **right** end. The element `nums[r]` lies in the rightmost block, so it saves `blockMax - nums[r]`. The structure is a stack on one side and a queue on the other, which makes it a monotonic deque.

## Approach

1. **Cost of a subarray.** To make `[l..r]` non-decreasing with increments only, the best choice is to raise each `nums[i]` to `M(i) = max(nums[l..i])`. That gives `cost(l, r) = Σ_{i=l..r} (M(i) - nums[i])`.

2. **Monotonicity.** Every term is non-negative, so for a fixed `l` the cost does not decrease as `r` grows. The prefix maxima can only go up when `l` moves left, so for a fixed `r` the cost does not decrease as `l` decreases either. This means the largest valid `r` for each `l` (call it `R(l)`) never increases as `l` moves left. Two pointers apply.

3. **Iterate `l` from `n-1` down to `0`**, keeping a window `[l..r]`. For each element, `M(i)` depends on what lies to its left, which is the end where elements arrive. Store the window as blocks `(max, count)`, ordered left to right with strictly increasing maxima. Each block is a run of elements whose prefix maximum is that block's leading element.

4. **Add `nums[l]`.** Start a new block `(nums[l], 1)`. While the leftmost existing block has `max <= nums[l]`, pop it, add `(nums[l] - max) * count` to the cost, and merge its count into the new block. Then push the new block. Each element joins and leaves the stack side at most once.

5. **Shrink from the right.** While `cost > k`, remove `nums[r]`. It belongs to the rightmost block, so subtract `block.max - nums[r]`, decrement the block's count (remove the block when it reaches 0), and decrement `r`.

6. **Count.** Every subarray `[l..j]` with `l <= j <= r` is valid, so add `r - l + 1`.

In Go, one slice holds both ends: append and pop at the tail (the left side of the window) and advance a `head` index (the right side of the window).

**Example:** `nums = [6,3,1,2,4,4], k = 7`
- l=5..1: the window `[3,1,2,4,4]` has blocks `(3,3),(4,2)` and cost `0+2+1+0+0 = 3`. The answer so far is 1+2+3+4+5 = 15.
- l=0 (6): blocks `(3,3)` and `(4,2)` are both popped. The cost rises by `3*3 + 2*2 = 13`, giving 16, so we shrink. Removing 4 saves 2 (14), removing the next 4 saves 2 (12), 2 saves 4 (8), and 1 saves 5 (3). The window is now `[6,3]`, so we add 2. The total is **17**.

## Complexity Analysis

Time Complexity: O(n). Each index is pushed once, merged at most once, and removed from the right at most once.
Space Complexity: O(n) for the block deque.

## Edge Cases

- **Single element:** the cost is always 0, so the answer is 1.
- **Already non-decreasing array:** the cost never goes above 0, so all `n(n+1)/2` subarrays count.
- **Strictly decreasing array with small `k`:** the cost grows fast and the window keeps shrinking. This checks that `r` never moves behind `l` (a window of one element always costs 0, so it does not).
- **Large values (`1e9`) with large `k`:** the costs can reach about `1e14`, so use 64-bit integers. Go's `int` is 64-bit on LeetCode, and the answer is returned as `int64`.
- **Equal elements:** a block with `max == nums[l]` must be merged (`<=`, not `<`). Either way the cost added is 0, and merging keeps the blocks' maxima strictly increasing.
- **A block emptied by right-side removal:** remove it from the queue so the next removal uses the right block's max.
