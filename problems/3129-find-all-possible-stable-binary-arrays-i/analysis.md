# 3129. Find All Possible Stable Binary Arrays I

[LeetCode Link](https://leetcode.com/problems/find-all-possible-stable-binary-arrays-i/)

Difficulty: Medium
Topics: Dynamic Programming, Prefix Sum
Acceptance Rate: 34.4%

## Hints

### Hint 1

Think about building the array one element at a time. At each step, what information do you need to decide whether placing a 0 or a 1 is valid? Consider tracking how many of each value you've used so far and what the last element was.

### Hint 2

Define a DP state based on the number of 0s used, the number of 1s used, and which value was placed last (0 or 1). The "limit" constraint means you cannot have more than `limit` consecutive identical elements. Think about how to enforce this: rather than tracking the exact run length, can you use subtraction (inclusion-exclusion) to remove the overcounted invalid arrangements?

### Hint 3

The key insight is the subtraction trick. Let `dp[i][j][0]` be the number of valid arrays using `i` zeros and `j` ones that end with a 0. The recurrence is:

`dp[i][j][0] = dp[i-1][j][0] + dp[i-1][j][1]`

This counts all ways to append a 0. But some of these create a run of `limit+1` consecutive zeros. That happens exactly when the previous `limit` elements were also all 0, meaning the element before those was a 1. So subtract `dp[i-limit-1][j][1]` when `i > limit`. The same logic applies symmetrically for ending with 1.

## Approach

We use a 3D DP table where `dp[i][j][last]` represents the number of valid (stable) binary arrays that use exactly `i` zeros and `j` ones, with the last element being `last` (0 or 1).

**Base cases:**
- `dp[i][0][0] = 1` for `1 <= i <= limit` (an array of all zeros is valid only if its length doesn't exceed `limit`)
- `dp[0][j][1] = 1` for `1 <= j <= limit` (an array of all ones is valid only if its length doesn't exceed `limit`)

**Transitions:**
- `dp[i][j][0] = dp[i-1][j][0] + dp[i-1][j][1]`: we can append a 0 whether the previous element was 0 or 1
  - If `i > limit`: subtract `dp[i-limit-1][j][1]` to exclude arrangements that would create a run of `limit+1` consecutive zeros
- `dp[i][j][1] = dp[i][j-1][0] + dp[i][j-1][1]`: we can append a 1 whether the previous element was 0 or 1
  - If `j > limit`: subtract `dp[i][j-limit-1][0]` to exclude arrangements that would create a run of `limit+1` consecutive ones

The subtraction works because a run of exactly `limit+1` zeros ending at position `(i, j)` means the element at position `(i-limit-1, j)` must have been a 1 (to start the invalid run). This is a clean inclusion-exclusion that avoids tracking the run length explicitly.

**Answer:** `(dp[zero][one][0] + dp[zero][one][1]) % MOD`

## Complexity Analysis

Time Complexity: O(zero * one) — we fill a table of size `(zero+1) x (one+1) x 2` with O(1) work per cell.

Space Complexity: O(zero * one) — for the DP table.

## Edge Cases

- **limit >= zero + one**: The limit is large enough that no constraint is ever violated. The answer becomes C(zero+one, zero) — all interleavings are valid.
- **limit = 1**: Maximum constraint. Only strictly alternating arrays are valid. If `|zero - one| > 1`, the answer is 0; otherwise it's 1 or 2.
- **zero = 0 or one = 0**: Only one possible array (all ones or all zeros), valid only if the count doesn't exceed `limit`.
- **Modular arithmetic**: All additions and subtractions must be done modulo 10^9+7, taking care to keep values non-negative after subtraction.
