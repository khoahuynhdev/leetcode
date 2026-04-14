# 2463. Minimum Total Distance Traveled

[LeetCode Link](https://leetcode.com/problems/minimum-total-distance-traveled/)

Difficulty: Hard
Topics: Array, Dynamic Programming, Sorting
Acceptance Rate: 59.8%

## Hints

### Hint 1

This is an assignment problem on a 1D line. Before you think about algorithms, ask: does the ordering of robots and factories along the X-axis matter? There is a classical exchange argument for this class of problem — if you swap the endpoints of two "crossed" assignments, the total distance never gets worse. What does that imply about the structure of an optimal assignment?

### Hint 2

Once you convince yourself that, after sorting both arrays by position, each factory ends up serving a **contiguous block** of robots, the problem reduces to partitioning the sorted robot array into contiguous groups and assigning each group to one factory (in order, possibly empty). That is a very DP-friendly setup: decide, factory by factory, how many of the next robots to absorb.

### Hint 3

Define `dp[i][j]` as the minimum total distance to repair robots `i..n-1` using factories `j..m-1` (after sorting both by position). Two choices at state `(i, j)`:
1. Skip factory `j`: transition to `dp[i][j+1]`.
2. Let factory `j` repair `k` robots for `k = 1..min(limit_j, n-i)`: the cost is `sum_{t=0..k-1} |factory[j].pos - robot[i+t]|` plus `dp[i+k][j+1]`.

Accumulate the inner sum as you increase `k` so each transition is O(1) amortized.

## Approach

1. **Sort**: sort `robot` ascending; sort `factory` ascending by position. This makes the "contiguous block" structure explicit.
2. **State**: `dp[i][j]` = minimum total travel to repair robots `robot[i..n-1]` using factories `factory[j..m-1]`.
3. **Base cases**:
   - `dp[n][j] = 0` for all `j` (no robots left to fix).
   - `dp[i][m] = +∞` for `i < n` (robots left but no factories).
4. **Transitions**: for each `(i, j)` with `i < n` and `j < m`:
   - Option A — skip factory `j`: `dp[i][j+1]`.
   - Option B — assign the next `k` robots to factory `j` (`1 ≤ k ≤ min(limit_j, n-i)`): running cost `cost += |factory[j].pos - robot[i+k-1]|`, then `cost + dp[i+k][j+1]`.
   - `dp[i][j]` is the minimum over all options.
5. **Answer**: `dp[0][0]`.

Walk-through for example 1 (`robot=[0,4,6]`, `factory=[[2,2],[6,2]]`, already sorted):
- Factory `[6,2]` best serves robots `{4, 6}` for cost `|6-4|+|6-6| = 2`.
- Factory `[2,2]` then serves only robot `0` at cost `|2-0| = 2`.
- Total `= 4`, matching the expected answer.

Using `int64` for the cost is important: up to 100 robots, each factory can be `~2·10⁹` away, so individual terms already exceed 32-bit `int` safety margins on some platforms.

## Complexity Analysis

Let `n = len(robot)` and `m = len(factory)`.

Time Complexity: `O(n·m·L)` where `L` is the maximum factory limit (at most `n`). In the worst case `O(n²·m)`, which is `≤ 100·100·100 = 10⁶` for the given constraints.
Space Complexity: `O(n·m)` for the DP table. Can be reduced to `O(n)` by iterating factories outer-to-inner and keeping only the next column, but the 2D table is clearer and well within limits.

## Edge Cases

- **A factory with `limit = 0`**: it cannot repair any robot, so only the "skip" transition contributes. The DP naturally handles this because the `k` loop runs zero times.
- **Robot coincides with a factory position**: the contribution is `0`, but the factory still consumes one unit of its limit. The sorted DP treats this correctly.
- **Negative coordinates / large magnitudes up to 10⁹**: absolute differences can reach ~`2·10⁹`, and summed totals can exceed 32-bit range, so use `int64` throughout.
- **More total capacity than robots**: some factories will be left unused; the "skip factory" transition covers this.
- **Only one factory with exactly enough capacity**: the DP collapses to a single block assignment — still correct.
- **All robots at the same side of all factories**: sorting ensures the optimal contiguous block property still applies.
