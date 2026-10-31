# 3665. Twisted Mirror Path Count

[LeetCode Link](https://leetcode.com/problems/twisted-mirror-path-count/)

Difficulty: Medium
Topics: Array, Dynamic Programming, Matrix
Acceptance Rate: 57.5%

## Hints

### Hint 1

This is a twist on the classic "Unique Paths" grid problem: the robot still only moves right or down, so there are no cycles and you can count paths cell by cell. Ask yourself what about the classic DP breaks once mirrors are added.

### Hint 2

When the robot hits a mirror, where it goes next depends on **how it arrived**, not just where it is. A single count per cell isn't enough. Add one more dimension to your DP state.

### Hint 3

Let `dp[i][j][d]` be the number of ways to *arrive* at `(i, j)` moving in direction `d` (right or down). An empty cell sends all of its arrivals both right and down. A mirror sends right-arrivals down and down-arrivals right. A chain of mirrors needs no special handling: a deflected arrival is just another arrival, and the next cell processes it. Since every move goes right or down, row-major order processes each cell after everything that feeds into it.

## Approach

**State.** For each cell keep two counts:
- `dp[i][j][R]`: ways to enter `(i, j)` while moving right
- `dp[i][j][D]`: ways to enter `(i, j)` while moving down

Seed the start with `dp[0][0][R] = 1`. The start cell is empty, so the direction doesn't matter.

**Transitions.** Process cells in row-major order:
- **Empty cell (`0`)**: the robot stops here and picks its next move. Let `total = R + D`.
  - push `total` to `dp[i][j+1][R]` (move right)
  - push `total` to `dp[i+1][j][D]` (move down)
- **Mirror cell (`1`)**: the robot never rests here, it is deflected right away.
  - `R` arrivals turn down: push to `dp[i+1][j][D]`
  - `D` arrivals turn right: push to `dp[i][j+1][R]`

Drop any push that would leave the grid, because those paths are invalid. The answer is `R + D` at `(m-1, n-1)`, taken mod `1e9+7`.

**Why it works.** The problem's "keep reflecting until you land on a non-mirror" rule is exactly what you get by passing arrivals through each mirror one at a time. A robot that enters a mirror moving right leaves it moving down, enters the next cell moving down, and so on. All edges point right or down, so the graph of (cell, direction) states is a DAG and row-major order is a valid topological order. Each distinct path is counted exactly once.

**Example 1** (`[[0,1,0],[0,0,1],[1,0,0]]`):
- `(0,0)` is empty, so it sends 1 right into the mirror at `(0,1)` and 1 down to `(1,0)`.
- Mirror `(0,1)` got 1 right-arrival and turns it down into `(1,1)`.
- `(1,0)` is empty with 1 arrival. It sends 1 right to `(1,1)` and 1 down into the mirror at `(2,0)`.
- `(1,1)` now has 2 arrivals. It sends 2 right into mirror `(1,2)` and 2 down to `(2,1)`.
- Mirror `(1,2)` turns its 2 right-arrivals down into `(2,2)`.
- Mirror `(2,0)` turns its 1 down-arrival right into `(2,1)`, which now holds 3. All 3 go right to `(2,2)`.
- `(2,2)` gets 2 + 3 = **5**.

**Space optimization.** Contributions only go to the same row (to the right) or the next row, so two rolling rows of `[2]int` are enough.

## Complexity Analysis

Time Complexity: O(m · n). Each cell is processed once with O(1) work.
Space Complexity: O(n) with two rolling rows (O(m · n) with a full 3D table).

## Edge Cases

- **Reflection out of bounds**: a mirror in the last row hit while moving down sends the robot right, and one in the last column hit while moving right sends it down. Either can leave the grid, so those contributions must be dropped instead of wrapping around or crashing.
- **Chained mirrors**: several mirrors in a row or a zig-zag, as in Example 3. The direction-aware state handles these without an explicit loop, but a simulation-based approach has to follow the whole chain.
- **Mirror right next to the target**: a deflection can land the robot exactly on `(m-1, n-1)`, and that counts as reaching it. Example 1 does this through `(1,2)`.
- **No mirrors**: the answer drops back to the binomial `C(m+n-2, m-1)`, which is a useful sanity check.
- **Large grids (500×500)**: counts overflow quickly, so take the modulo on every addition.
- **Zero valid paths**: return 0 instead of something left over from earlier cells.
