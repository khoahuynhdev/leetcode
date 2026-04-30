# 3742. Maximum Path Score in a Grid

[LeetCode Link](https://leetcode.com/problems/maximum-path-score-in-a-grid/)

Difficulty: Medium
Topics: Array, Dynamic Programming, Matrix
Acceptance Rate: 42.3%

## Hints

### Hint 1

You can only move right or down, so this looks like a classic grid path problem. A greedy "always pick the larger neighbor" strategy will not work because of the cost budget — the most valuable cell is not always the cheapest. Think about which family of techniques typically handles "best path under a budget constraint" on a grid.

### Hint 2

Try Dynamic Programming. The usual 2D DP `dp[i][j] = best score to reach (i, j)` is not enough on its own, because two paths to the same cell can differ in score *and* in cost spent. A path with a smaller score but smaller cost may still win later when the cells beyond are expensive. You need to remember the cost as part of the state.

### Hint 3

Add a third dimension: `dp[i][j][c] = maximum score to reach cell (i, j) using exactly cost c`. From any cell, you can come from above or from the left. The cell (i, j) contributes `score = grid[i][j]` and `cost = (1 if grid[i][j] > 0 else 0)`. The answer is the maximum of `dp[m-1][n-1][c]` over all `c <= k`. Since the path length is at most `m + n - 1`, you can cap `c` at `min(k, m + n - 1)`. Use a rolling array over rows to keep memory tight.

## Approach

Define `dp[i][j][c]` = the maximum score on a valid path from `(0, 0)` to `(i, j)` whose total cost equals exactly `c`. Use a sentinel (e.g. a very negative number) to mark unreachable states.

For each cell `(i, j)` with value `v`:

- `costInc = 1` if `v > 0`, else `0`.
- `scoreInc = v`.
- `dp[i][j][c] = max(dp[i-1][j][c - costInc], dp[i][j-1][c - costInc]) + scoreInc`, only when `c >= costInc` and at least one source is reachable.

Base case: `dp[0][0][0] = 0` because `grid[0][0] == 0` always (cost 0, score 0).

Final answer: `max(dp[m-1][n-1][c])` for `c` in `[0, maxCost]`, where `maxCost = min(k, m + n - 1)`. If every entry is the sentinel, return `-1`.

Memory optimization: each row only depends on the previous row plus the current row's already-filled entries to its left. Keep two 2D slices `prev[n][maxCost+1]` and `cur[n][maxCost+1]` and swap them after each row.

Walkthrough on Example 1 — `grid = [[0, 1], [2, 0]]`, `k = 1`, `maxCost = 1`:

- Row 0: `dp[0][0][0] = 0`. At `(0, 1)` (v=1), `dp[0][1][1] = dp[0][0][0] + 1 = 1`.
- Row 1: At `(1, 0)` (v=2), `dp[1][0][1] = dp[0][0][0] + 2 = 2`. At `(1, 1)` (v=0), `dp[1][1][1] = max(dp[0][1][1], dp[1][0][1]) + 0 = 2`.
- Answer = max over c in [0, 1] of `dp[1][1][c]` = 2.

For Example 2 the value-2 cell is at the corner and every path needs at least cost 2, so all entries at `(m-1, n-1)` for `c <= 1` stay at the sentinel and we return `-1`.

## Complexity Analysis

Time Complexity: O(m * n * min(k, m + n))
Space Complexity: O(n * min(k, m + n)) using the rolling-row optimization

## Edge Cases

- 1x1 grid: only `grid[0][0] = 0`, so the answer is `0` regardless of `k`.
- `k = 0` with any non-zero cell on every path: return `-1` because every path costs at least 1.
- Grid of all zeros: every path has cost 0 and score 0.
- `k` larger than `m + n - 1`: the cost budget cannot bind, so the answer is the unconstrained maximum-score path.
- Two paths with equal score but different costs: the DP correctly keeps both because cost is part of the state — important so a future cell can extend whichever fits.
- The destination is unreachable within budget even if many shorter prefixes were reachable: must check the sentinel before returning, and return `-1` in that case.
