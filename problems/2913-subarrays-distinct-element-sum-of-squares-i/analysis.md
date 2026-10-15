# 2913. Subarrays Distinct Element Sum of Squares I

[LeetCode Link](https://leetcode.com/problems/subarrays-distinct-element-sum-of-squares-i/)

Difficulty: Easy
Topics: Array, Hash Table, Binary Indexed Tree, Segment Tree
Acceptance Rate: 83.6%

## Hints

### Hint 1

Look at the constraints: `nums.length <= 100` and `nums[i] <= 100`. Small limits like these mean you can visit every subarray directly. Don't reach for a segment tree yet. That's for part II of this problem.

### Hint 2

Fix the left endpoint `i`, then extend the right endpoint `j` one step at a time. As the window grows you only add elements. If you keep track of what you've already seen, you can update the distinct count without recounting the whole subarray.

### Hint 3

Use a "seen" set (a `[101]bool` array works since values are at most 100) that you reset for each new `i`. When you extend to `j`, the distinct count goes up by one only if `nums[j]` hasn't been seen yet. Add `count * count` to the answer after each extension. That gives O(n²) total.

## Approach

There are `n(n+1)/2` subarrays. Recounting each one from scratch costs O(n), which is O(n³) overall. That is still fine for n = 100, but we can get O(n²) easily:

1. Initialize `ans = 0`.
2. For each start index `i`:
   - Reset a boolean array `seen` of size 101 and set `distinct = 0`.
   - For each end index `j` from `i` to `n-1`:
     - If `nums[j]` is not in `seen`, mark it and increment `distinct`.
     - `distinct` is now the distinct count of `nums[i..j]`, so add `distinct * distinct` to `ans`.
3. Return `ans`.

This works because `nums[i..j]` is just `nums[i..j-1]` plus one element. The distinct count can only stay the same or grow by exactly one, depending on whether the new element appeared earlier in the window.

**Example:** `nums = [1,2,1]`

- `i = 0`: add 1 → distinct 1 (+1); add 2 → distinct 2 (+4); add 1 → already seen, distinct 2 (+4). Subtotal 9.
- `i = 1`: add 2 → 1 (+1); add 1 → 2 (+4). Subtotal 5.
- `i = 2`: add 1 → 1 (+1). Subtotal 1.

Total = 9 + 5 + 1 = **15**. ✓

**Going further:** Part II (problem 2916) has n up to 10⁵. There you count how each new element changes the sum of squares for all subarrays ending at `j`, using a segment tree with lazy range-add. If `(d+1)² - d² = 2d + 1`, then extending all windows that don't already contain `nums[j]` adds `2·(sum of their counts) + (number of such windows)`. That's where the BIT/Segment Tree topic tags come from.

## Complexity Analysis

Time Complexity: O(n²) — two nested loops with O(1) work inside.
Space Complexity: O(1) — the `seen` array has a fixed size of 101, no matter what n is.

## Edge Cases

- **Single element** (`[5]`): only one subarray, distinct count 1, answer 1.
- **All elements equal** (`[1,1]`): every subarray has distinct count 1, so the answer is just the number of subarrays, `n(n+1)/2`.
- **All elements distinct** (`[1,2,3]`): the distinct count equals the subarray length. This is the largest possible answer for a given n.
- **Repeated element later in the window** (`[1,2,1]`): make sure a repeat doesn't increment the count. That's exactly what `seen` is for.
- **Resetting state**: forgetting to clear `seen` between start indices carries over stale values and undercounts.
