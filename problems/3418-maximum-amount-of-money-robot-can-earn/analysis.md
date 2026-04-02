# 3418. Maximum Amount of Money Robot Can Earn

[LeetCode Link](https://leetcode.com/problems/maximum-amount-of-money-robot-can-earn/)

Difficulty: Medium
Topics: Array, Dynamic Programming, Matrix
Acceptance Rate: 33.7%

## Hints

### Hint 1

This is a grid path problem where the robot can only move right or down. Think about how you would solve a basic "maximum path sum" problem first — then consider what additional state you need to track for the neutralization ability.

### Hint 2

Standard grid DP uses `dp[i][j]` to store the best result reaching cell `(i, j)`. Here, the robot has up to 2 neutralizations. What if you added a third dimension to track how many neutralizations have been used so far? Consider `dp[i][j][k]` where `k` is the number of neutralizations used.

### Hint 3

At each cell, if `coins[i][j] < 0`, you have a choice: either take the loss (add the negative value) or spend a neutralization (treat the cell as 0, incrementing `k`). If `coins[i][j] >= 0`, you always collect the coins — there's no benefit to neutralizing a non-negative cell. The transition comes from the cell above or to the left, picking whichever gives a higher total.

## Approach

We use 3D dynamic programming with state `dp[i][j][k]` representing the maximum coins the robot can have when it reaches cell `(i, j)` having used exactly `k` neutralizations (where `k` ranges from 0 to 2).

**Transitions:**

For each cell `(i, j)`, let `prev` be the maximum of `dp[i-1][j][k]` (from above) and `dp[i][j-1][k]` (from the left), handling boundary conditions.

1. **If `coins[i][j] >= 0`:** The robot always collects the coins.
   - `dp[i][j][k] = prev + coins[i][j]`

2. **If `coins[i][j] < 0`:** The robot can either take the loss or neutralize:
   - Without neutralizing: `dp[i][j][k] = prev + coins[i][j]`
   - With neutralizing (only if `k >= 1`): `dp[i][j][k] = max(dp[i-1][j][k-1], dp[i][j-1][k-1]) + 0`
   - Take the better of these two options.

**Base case:** `dp[0][0][0] = coins[0][0]`. If `coins[0][0] < 0`, then `dp[0][0][1] = 0` (neutralize the starting cell).

**Answer:** `max(dp[m-1][n-1][0], dp[m-1][n-1][1], dp[m-1][n-1][2])`

**Example walkthrough with `[[0,1,-1],[1,-2,3],[2,-3,4]]`:**

The optimal path is (0,0) → (0,1) → (1,1) → (1,2) → (2,2), collecting 0 + 1 + 0(neutralize -2) + 3 + 4 = 8.

## Complexity Analysis

Time Complexity: O(m × n × 3) = O(m × n), where m and n are the grid dimensions. We visit each cell once for each of the 3 neutralization states.

Space Complexity: O(m × n × 3) = O(m × n) for the DP table. This can be optimized to O(n × 3) = O(n) using row-by-row processing, but the full table is clearer for understanding.

## Edge Cases

- **Single cell grid (1×1):** The robot starts and ends at the same cell. If negative, it can neutralize it (answer is 0). If non-negative, that's the answer.
- **All positive values:** No neutralizations are needed; this reduces to a standard maximum path sum problem.
- **All negative values:** The robot should neutralize the two most costly robbers on its path. The answer may still be negative.
- **Grid dimensions 1×n or m×1:** Only one possible path exists (all right or all down), so the robot just picks the 2 worst robbers to neutralize along that single path.
- **More than 2 robbers on every path:** The robot must accept losses from some robbers — it should strategically neutralize the ones that cost the most along the optimal route.
