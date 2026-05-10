# 2770. Maximum Number of Jumps to Reach the Last Index

[LeetCode Link](https://leetcode.com/problems/maximum-number-of-jumps-to-reach-the-last-index/)

Difficulty: Medium
Topics: Array, Dynamic Programming
Acceptance Rate: 37.2%

## Hints

### Hint 1

You want the *maximum* number of jumps, not the minimum. That word "maximum" is a strong signal: from any index, several next indices may be reachable, and you must explore all of them rather than greedily picking the farthest. Think about what category of technique lets you consider every choice and combine sub-results — this smells like dynamic programming over the indices.

### Hint 2

Define a value for each index `i`: the best (largest) number of jumps with which you can *arrive* at index `i` starting from index `0`. Index `0` arrives "for free" with 0 jumps. Every other index gets its value from some earlier index that can legally jump to it. If no earlier reachable index can jump to `i`, then `i` is unreachable.

### Hint 3

Iterate `i` from `1` to `n-1`. For each `i`, scan all earlier indices `j < i`. The jump `j -> i` is allowed when `|nums[i] - nums[j]| <= target`. If `j` is reachable (its value isn't the "unreachable" marker) and the jump is legal, then `i` can be reached with `dp[j] + 1` jumps — keep the maximum over all such `j`. The answer is `dp[n-1]`, or `-1` if it stayed unreachable. With `n <= 1000`, an `O(n^2)` double loop is comfortably fast enough; you don't need anything cleverer.

## Approach

This is a classic "longest path in a DAG" expressed as 1‑D dynamic programming over array positions.

Let `dp[i]` = the maximum number of jumps needed to reach index `i` from index `0`, using `-1` to mean "index `i` is not reachable".

1. Initialize `dp[0] = 0` (we start there with zero jumps) and `dp[i] = -1` for all `i >= 1`.
2. For each `i` from `1` to `n-1`:
   - For each `j` from `0` to `i-1`:
     - If `dp[j] == -1`, skip — we can't even get to `j`, so we can't jump from it.
     - Compute `diff = |nums[i] - nums[j]|`. If `diff <= target`, the jump `j -> i` is legal.
     - If legal, then arriving at `i` via `j` costs `dp[j] + 1` jumps. Update `dp[i] = max(dp[i], dp[j] + 1)`.
3. Return `dp[n-1]`.

Why it works: any valid jumping sequence ending at index `i` has a second-to-last position `j < i` with a legal jump `j -> i`. The number of jumps in that sequence is (jumps to reach `j`) + 1. To maximize the total, we maximize the jumps to reach `j` — exactly `dp[j]`. Because we process indices left to right, `dp[j]` is already finalized when we use it. Indices with no valid predecessor keep `-1`, and that `-1` correctly prevents them from being used as a stepping stone.

Walkthrough on Example 1, `nums = [1,3,6,4,1,2]`, `target = 2`:

- `dp = [0, -1, -1, -1, -1, -1]`
- `i = 1`: `j = 0`, `|3-1| = 2 <= 2` → `dp[1] = dp[0] + 1 = 1`
- `i = 2`: `j = 0` diff 5 (no); `j = 1` diff 3 (no) → `dp[2] = -1`
- `i = 3`: `j = 1` diff 1 → `dp[3] = dp[1] + 1 = 2`
- `i = 4`: `j = 0` diff 0 → `1`; `j = 1` diff 2 → `dp[1]+1 = 2` → `dp[4] = 2`
- `i = 5`: `j = 3` diff 2 → `dp[3]+1 = 3`; `j = 4` diff 1 → `dp[4]+1 = 3` → `dp[5] = 3`

Answer: `dp[5] = 3`. ✔

Honest note: the idea is not hard once you see it as "longest path / max-jumps DP", but it's easy to trip on (a) using `-1` both as the answer sentinel and the unreachable marker, and (b) accidentally writing the *minimum*-jumps recurrence out of habit. Be deliberate about the `max`.

## Complexity Analysis

Time Complexity: O(n^2) — for each of the `n` indices we scan all earlier indices.
Space Complexity: O(n) — the `dp` array.

## Edge Cases

- **`target = 0`**: only jumps between indices with equal values are allowed; often the last index is unreachable → return `-1` (Example 3).
- **Last index unreachable in general**: make sure `-1` propagates — never build on a `dp[j] == -1`.
- **Minimum-size array (`n == 2`)**: the only question is whether `|nums[1] - nums[0]| <= target`; answer is `1` or `-1`.
- **All elements equal**: every jump is legal, so the maximum is `n - 1` (jump one step at a time).
- **Large value range / negative numbers**: `nums[i]` can be down to `-1e9` and `target` up to `2e9`; the difference fits in a 64‑bit int (Go's `int` is 64‑bit on common platforms), so plain integer subtraction and `abs` are safe — no overflow concerns here.
