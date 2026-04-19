# 1855. Maximum Distance Between a Pair of Values

[LeetCode Link](https://leetcode.com/problems/maximum-distance-between-a-pair-of-values/)

Difficulty: Medium
Topics: Array, Two Pointers, Binary Search
Acceptance Rate: 56.2%

## Hints

### Hint 1

Both arrays are **non-increasing**. Whenever you see "sorted" (even in reverse) paired with a pair/window search, think about whether you can avoid the brute-force `O(n*m)` nested loop by exploiting monotonicity. What classic patterns leverage two sorted sequences to work in linear time?

### Hint 2

Two strong candidates here: **binary search** (for every `i`, find the largest `j >= i` with `nums2[j] >= nums1[i]`) and **two pointers**. The two-pointer approach is the cleanest because both arrays are sorted in the same direction. Think about what it means to advance each pointer: when do you move `i` forward, when do you move `j` forward, and how do you preserve the invariant `i <= j`?

### Hint 3

Use pointers `i` and `j` both starting at `0`.
- If `i > j`, the pair is invalid — push `j` forward (it must never fall behind `i`).
- If `nums1[i] > nums2[j]`, no `j'` at or beyond this `j` can satisfy the condition for this `i` (since `nums2` is non-increasing, `nums2[j']` only gets smaller). So this `i` is finished — advance `i`.
- Otherwise the pair is valid; record `j - i` and try to stretch the distance further by advancing `j`.

Each pointer only moves forward, so the work is linear.

## Approach

Walk `i` over `nums1` and `j` over `nums2`, both starting at `0`:

1. Skip the prefix where `i > j` by nudging `j` forward — we need `j >= i` for a valid pair.
2. If `nums1[i] > nums2[j]`, this `i` can never pair with `j` or anything later (since `nums2` only decreases), so advance `i`.
3. Otherwise `(i, j)` is valid. Update `ans = max(ans, j - i)` and advance `j` to try to widen the distance.

Why advancing `j` on a valid hit is safe: we already recorded the current distance, and a larger `j` could only help the answer, not hurt it — we are free to abandon smaller `j` values because they produce smaller distances for the same `i`.

Why advancing `i` on an invalid hit is safe: `nums2[j], nums2[j+1], ...` are all `<= nums2[j] < nums1[i]`, so no later `j` can rescue this `i`.

**Example trace** — `nums1 = [55,30,5,4,2]`, `nums2 = [100,20,10,10,5]`:

| step | i | j | nums1[i] | nums2[j] | action | ans |
|------|---|---|----------|----------|--------|-----|
| 1 | 0 | 0 | 55 | 100 | valid, j++ | 0 |
| 2 | 0 | 1 | 55 | 20  | invalid, i++ | 0 |
| 3 | 1 | 1 | 30 | 20  | invalid, i++ | 0 |
| 4 | 2 | 1 | —  | —   | i>j, j++ | 0 |
| 5 | 2 | 2 | 5  | 10  | valid, j++ | 0 |
| 6 | 2 | 3 | 5  | 10  | valid, j++ | 1 |
| 7 | 2 | 4 | 5  | 5   | valid, j++ | 2 |

Final answer: `2`.

## Complexity Analysis

Time Complexity: O(n + m), where `n = len(nums1)` and `m = len(nums2)`. Each pointer advances at most `n` or `m` times.
Space Complexity: O(1). Only a handful of integer variables.

## Edge Cases

- **No valid pair exists**: e.g. `nums1 = [5]`, `nums2 = [1]`. Return `0` (the problem defines this explicitly; the two-pointer loop naturally yields `0`).
- **Single valid pair at `(0, 0)`**: distance is `0`, which is still a valid answer to return.
- **Length mismatch** `len(nums1) != len(nums2)`: loop terminates when either pointer runs out; the `i <= j` constraint caps `j` relative to `nums1`'s exhaustion.
- **Plateaus of equal values** (e.g. `nums2 = [25,25,25,25,25]`): equality satisfies `nums1[i] <= nums2[j]`, so don't accidentally use strict `<`.
- **All of `nums1` larger than all of `nums2`**: the `i++` branch fires repeatedly until `i` exits; answer stays `0`.
