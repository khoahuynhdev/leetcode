# 3693. Climbing Stairs II

[LeetCode Link](https://leetcode.com/problems/climbing-stairs-ii/)

Difficulty: Medium
Topics: Array, Dynamic Programming
Acceptance Rate: 70.2%

## Hints

### Hint 1

This looks a lot like the classic "Climbing Stairs" / "Min Cost Climbing Stairs" problems. Ask yourself: if you already knew the cheapest way to reach every earlier step, could you work out the cheapest way to reach the current one?

### Hint 2

Define `dp[j]` as the minimum total cost to land on step `j`. Which earlier steps can lead directly to step `j`? There are only three of them, and each jump has a known, fixed cost.

### Hint 3

The jump cost `costs[j] + (j - i)^2` only depends on the destination `j` and the jump length (1, 2, or 3). So:

`dp[j] = costs[j] + min(dp[j-1] + 1, dp[j-2] + 4, dp[j-3] + 9)`

with `dp[0] = 0` and any negative index treated as unreachable. You only ever look back three steps, so you don't need the whole array.

## Approach

1. **State**: `dp[j]` = minimum cost to arrive at step `j`. Start with `dp[0] = 0`.
2. **Transition**: To land on step `j`, the last jump came from `i ∈ {j-1, j-2, j-3}` (only those with `i >= 0`). That jump costs `costs[j] + (j-i)^2`, i.e. `costs[j] + 1`, `+4`, or `+9`. Since `costs[j]` is paid no matter which predecessor you pick, factor it out:
   `dp[j] = costs[j] + min(dp[j-1] + 1, dp[j-2] + 4, dp[j-3] + 9)`.
3. **Answer**: `dp[n]`.
4. **Space optimization**: Each state depends only on the previous three, so keep three rolling variables `a = dp[j-3]`, `b = dp[j-2]`, `c = dp[j-1]`. Initialize `a = b = ∞` (steps -2 and -1 don't exist) and `c = 0` (step 0). After computing the new value, shift: `a, b, c = b, c, new`.

Why it works: any path to step `j` must end with exactly one of the three possible jumps, and the cost of everything before that jump is independent of the jump itself. So the optimal path to `j` is an optimal path to some predecessor plus one jump — classic optimal substructure.

**Walkthrough** (Example 2, `costs = [5,1,6,2]`):

| j | from j-1 | from j-2 | from j-3 | costs[j] | dp[j] |
|---|----------|----------|----------|----------|-------|
| 0 | – | – | – | – | 0 |
| 1 | 0+1=1 | ∞ | ∞ | 5 | 6 |
| 2 | 6+1=7 | 0+4=4 | ∞ | 1 | 5 |
| 3 | 5+1=6 | 6+4=10 | 0+9=9 | 6 | 12 |
| 4 | 12+1=13 | 5+4=9 | 6+9=15 | 2 | 11 |

Answer: `11`.

## Complexity Analysis

Time Complexity: O(n) — one pass, constant work per step.
Space Complexity: O(1) — only three rolling DP values.

## Edge Cases

- **n = 1**: Only one possible jump (0 → 1), cost `costs[1] + 1`. Make sure the "unreachable" predecessors don't cause errors or get chosen.
- **n = 2 or 3**: Jumping directly from step 0 may be optimal (Example 3); the base cases must allow jumps of length 2 and 3 from step 0.
- **Unreachable sentinels**: Use a large-but-safe infinity (e.g. `1 << 60`) so adding 4 or 9 doesn't overflow.
- **Uniform low costs**: The squared jump penalty means small steps are often best (1+1 < 4), so don't assume "jump as far as possible" is greedy-optimal.
- **Expensive intermediate steps**: Large `costs` values should be skipped via longer jumps — the DP handles this naturally.
- **Magnitude**: Max total is roughly `10^5 * (10^4 + 1)` ≈ `10^9`, which fits comfortably in Go's 64-bit `int`.
