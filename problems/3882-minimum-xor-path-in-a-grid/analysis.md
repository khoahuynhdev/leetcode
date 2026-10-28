# 3882. Minimum XOR Path in a Grid

[LeetCode Link](https://leetcode.com/problems/minimum-xor-path-in-a-grid/)

Difficulty: Medium
Topics: Array, Dynamic Programming, Bit Manipulation, Matrix
Acceptance Rate: 48.2%

## Hints

### Hint 1

This looks like a classic "min path sum" grid DP, but try it: does keeping only the *minimum* XOR at each cell work? Build a tiny counterexample. Then look carefully at the constraints — especially the range of the cell values.

### Hint 2

Greedy-per-cell fails because XOR is not monotonic: a large prefix can become small after XORing with the next value. Instead of storing one value per cell, store *every* XOR value that can be reached at that cell. How many distinct values can there be?

### Hint 3

Every value is at most 1023 (10 bits), so any XOR of them is also in `[0, 1023]`. That means each cell has at most 1024 possible prefix XORs. Let `reach[i][j][x]` be true if some path to `(i, j)` has XOR `x`. Then `reach[i][j][x ^ grid[i][j]]` is true if `reach[i-1][j][x]` or `reach[i][j-1][x]` is true. The answer is the smallest `x` with `reach[m-1][n-1][x]` true.

## Approach

**Why the normal DP fails.** In min-path-sum, the best prefix at a cell always leads to the best total, because adding is monotonic. XOR isn't. For example, prefixes `1` and `2` at a cell, followed by a value `3`: `1 ^ 3 = 2` but `2 ^ 3 = 1`. Keeping only the smaller prefix (`1`) misses the better answer (`1`). So we have to remember *all* reachable prefix XORs.

**State compression via the value bound.** All values are in `[0, 1023]`, and XOR never sets bits beyond the highest bit of its operands, so every prefix XOR is also in `[0, 1023]`. Each cell's reachable set is therefore a boolean array of size 1024.

**Transition.**

1. At `(0, 0)` the only reachable value is `grid[0][0]`.
2. For any other cell `(i, j)`, take the union of the reachable sets of the cell above and the cell to the left, then XOR every value in it with `grid[i][j]`.
3. At `(m-1, n-1)`, scan `x` from 0 upward and return the first reachable value.

**Space trick.** Row `i` only depends on row `i-1` and on the cell to its left in the same row, so we keep one reachable-set per column and update it in place: before the update, `dp[j]` holds the set for the cell above; `dp[j-1]` already holds the set for the cell to the left in the current row.

**Example** — `grid = [[1,2],[3,4]]`:

- `(0,0)`: {1}
- `(0,1)`: {1^2} = {3}
- `(1,0)`: {1^3} = {2}
- `(1,1)`: ({3} ∪ {2}) ^ 4 = {7, 6} → min is **6**.

## Complexity Analysis

Time Complexity: O(m · n · V) where V = 1024, i.e. at most about 10^6 operations since m · n ≤ 1000.
Space Complexity: O(n · V) for one reachable-set per column.

## Edge Cases

- **Single cell (1x1)**: the answer is just `grid[0][0]`; make sure the start cell is initialized and no transition is required.
- **Single row or single column**: there is exactly one path; only one neighbor (left or above) exists for each cell, so don't read out of bounds.
- **XOR cancels to 0**: e.g. `[[2,7,5]]` — the answer can be 0, so the final scan must start at 0.
- **Zeros in the grid**: XOR with 0 leaves the set unchanged; the logic must still merge sets correctly.
- **Max values (1023)**: confirms the 1024-size array is large enough — XOR of 10-bit values stays 10-bit.
