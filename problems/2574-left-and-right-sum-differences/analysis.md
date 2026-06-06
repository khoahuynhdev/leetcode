# 2574. Left and Right Sum Differences

[LeetCode Link](https://leetcode.com/problems/left-and-right-sum-differences/)

Difficulty: Easy
Topics: Array, Prefix Sum
Acceptance Rate: 88.5%

## Hints

### Hint 1

For each index `i` you need to know the sum of everything *before* it and the sum of everything *after* it. Before reaching for nested loops, ask yourself: is there a way to reuse work you've already done as you scan across the array? This is the classic territory of the **prefix sum** pattern.

### Hint 2

Notice a useful relationship: if you know the **total sum** of the array and you keep a running **left sum** as you move from left to right, then the right sum at index `i` falls out for free. There's no need to compute `rightSum` with a separate pass over the elements to the right.

### Hint 3

At index `i`, let `left` be the sum of `nums[0..i-1]`. Then `rightSum[i] = total - left - nums[i]`, where `total` is the sum of the whole array. So `answer[i] = |left - (total - left - nums[i])|`. Compute `total` once, then sweep through once while updating `left` after you record each answer.

## Approach

The brute-force idea is to, for every index `i`, loop over all elements left of `i` and all elements right of `i`. That is `O(n^2)` and wasteful because it recomputes overlapping sums again and again.

We can do much better by recognizing that the left sums form a running total. Walk the array once and maintain a single accumulator `left` that holds the sum of all elements strictly to the left of the current index.

The right sum never needs its own loop. If `total` is the sum of every element, then the sum strictly to the right of index `i` is simply everything except the left part and the current element:

```
rightSum[i] = total - left - nums[i]
```

Algorithm:

1. Compute `total`, the sum of all elements in `nums`.
2. Initialize `left = 0` and an `answer` slice of length `n`.
3. For each index `i` from `0` to `n-1`:
   - Compute `right = total - left - nums[i]`.
   - Set `answer[i] = abs(left - right)`.
   - Add `nums[i]` to `left` so it is ready for the next index.
4. Return `answer`.

Walking through Example 1, `nums = [10, 4, 8, 3]`, `total = 25`:

| i | nums[i] | left (before) | right = 25 - left - nums[i] | answer = \|left - right\| |
|---|---------|---------------|-----------------------------|---------------------------|
| 0 | 10      | 0             | 25 - 0 - 10 = 15            | \|0 - 15\| = 15           |
| 1 | 4       | 10            | 25 - 10 - 4 = 11            | \|10 - 11\| = 1           |
| 2 | 8       | 14            | 25 - 14 - 8 = 3             | \|14 - 3\| = 11           |
| 3 | 3       | 22            | 25 - 22 - 3 = 0             | \|22 - 0\| = 22           |

This yields `[15, 1, 11, 22]`, matching the expected output.

This works because `left` is exactly `leftSum[i]` by construction (we only add `nums[i]` *after* recording the answer), and the right sum is derived from the invariant `total = leftSum[i] + nums[i] + rightSum[i]`.

## Complexity Analysis

Time Complexity: O(n) — one pass to compute the total and one pass to fill the answer.
Space Complexity: O(1) extra — ignoring the output array, we only keep a couple of integer accumulators. (The returned `answer` array is O(n), which is unavoidable.)

## Edge Cases

- **Single element** (`nums = [1]`): There is nothing to the left or right, so `leftSum[0] = rightSum[0] = 0` and the answer is `[0]`. The formula handles this automatically: `right = total - 0 - nums[0] = 0`.
- **All elements equal**: The running-sum logic still holds; this just confirms there are no off-by-one mistakes in updating `left`.
- **Largest array (n = 1000) with max values**: The total can be up to `1000 * 10^5 = 10^8`, which fits comfortably in Go's `int` (64-bit on common platforms, at least 32-bit guaranteed), so there is no overflow concern.
- **Absolute value**: Remember the answer uses `|leftSum - rightSum|`; forgetting the absolute value would produce negative numbers for early indices where the right side dominates.
