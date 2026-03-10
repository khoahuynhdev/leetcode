# 3130. Find All Possible Stable Binary Arrays II

[LeetCode Link](https://leetcode.com/problems/find-all-possible-stable-binary-arrays-ii/)

Difficulty: Hard
Topics: Dynamic Programming, Prefix Sum
Acceptance Rate: 42.6%

## Hints

### Hint 1

Think about building the binary array one element at a time from left to right. What state do you need to track to know whether the array is still "stable" so far? Consider how DP states can encode the counts of 0s and 1s used, along with what value was placed last.

### Hint 2

Define `dp[i][j][k]` as the number of valid arrays using exactly `i` zeros and `j` ones where the last element is `k` (0 or 1). The key challenge is enforcing the constraint that no run of consecutive identical values exceeds `limit`. Instead of tracking the full run length, think about how to subtract the overcounted invalid configurations using an inclusion-exclusion style correction.

### Hint 3

When computing `dp[i][j][0]` (appending a 0), the naive count is `dp[i-1][j][0] + dp[i-1][j][1]`. The overcounted cases are exactly those where the last `limit+1` positions are all 0s. Such arrays must have had a 1 at position `i-limit-1` (or the run started from the very beginning of the array). This gives a clean subtraction term. Use a "virtual" base state at `(0, 0)` to uniformly handle both mid-array and start-of-array boundary cases.

## Approach

We use a 2D DP table indexed by `(i, j)` where `i` is the number of 0s used and `j` is the number of 1s used. Each cell stores two values: the count of valid arrays ending with 0, and the count ending with 1.

**Recurrence:**

For `dp[i][j][0]` (place a 0 at the end):
- Add `dp[i-1][j][0]` (extend a run of 0s) and `dp[i-1][j][1]` (start a new run of 0s after a 1).
- If `i > limit`, subtract the invalid cases where the last `limit+1` entries are all 0s. These correspond to arrays where position `(i-limit-1, j)` ended with a 1, giving the subtraction term `dp[i-limit-1][j][1]`.

The trick is handling the boundary: when `i-limit-1 = 0` and `j = 0`, the entire array would be `limit+1` consecutive 0s. There is no preceding 1 in the DP table, but this case is exactly 1 invalid array. We handle this by defining a virtual function that returns 1 for `(0, 0)` and the normal DP value otherwise. Symmetrically for `dp[i][j][1]`.

**Virtual state:** Define `getV(i, j, k)` = 1 if `(i, j) = (0, 0)`, else `dp[i][j][k]`. This elegantly handles start-of-array boundaries in both the addition and subtraction terms.

The answer is `dp[zero][one][0] + dp[zero][one][1]` modulo `10^9 + 7`.

## Complexity Analysis

Time Complexity: O(zero * one) — we fill a 2D table of this size with O(1) work per cell.

Space Complexity: O(zero * one) — for the DP table.

## Edge Cases

- **limit >= zero + one**: The constraint never activates, so the answer is simply C(zero + one, zero) (the number of ways to interleave zeros and ones). The DP still handles this correctly.
- **limit = 1**: Maximum constraint. Only strictly alternating arrays are valid (e.g., [0,1,0,1,...] or [1,0,1,0,...]). The answer is 0, 1, or 2 depending on whether |zero - one| <= 1.
- **zero = 0 or one = 0**: The only valid array is all-zeros or all-ones, which is valid only if the count does not exceed `limit`. The DP correctly returns 1 or 0.
- **Large inputs (zero, one, limit = 1000)**: The O(zero * one) solution handles this efficiently within constraints.
