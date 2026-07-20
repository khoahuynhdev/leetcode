# 1260. Shift 2D Grid

[LeetCode Link](https://leetcode.com/problems/shift-2d-grid/)

Difficulty: Easy
Topics: Array, Matrix, Simulation
Acceptance Rate: 69.3%

## Hints

### Hint 1

Shifting each element one cell to the right, wrapping from the end of a row to the start of the next, and from the last cell of the grid back to the very first, is just a rotation. Ask yourself: what one-dimensional structure does a 2D grid become if you "unroll" it row by row?

### Hint 2

Instead of literally performing `k` separate shift passes (which is wasteful), think of the grid as a flat list of `m * n` numbers laid out in reading order. A single shift moves every value one position forward in that flat list, and the last value wraps around to the front. That means the whole operation is a **cyclic shift of a 1D array by `k`**.

### Hint 3

The key insight is index arithmetic. An element at row `i`, column `j` sits at flat index `i * n + j`. After shifting `k` times, its new flat index is `(i * n + j + k) mod (m * n)`. Convert that new flat index back to 2D coordinates with `row = idx / n` and `col = idx % n`. Because shifting by `m * n` returns the grid to its original state, reduce `k` with `k mod (m * n)` first — this also explains Example 3, where `k = 9` on a 9-cell grid changes nothing.

## Approach

Treat the grid as a one-dimensional sequence read left-to-right, top-to-bottom. There are `total = m * n` cells. One shift advances every element by exactly one position in this flattened order, with wrap-around. Therefore `k` shifts advance every element by `k` positions modulo `total`.

Algorithm:

1. Read the grid dimensions `m` (rows) and `n` (columns) and compute `total = m * n`.
2. Allocate a fresh result grid of the same shape so we never overwrite values we still need to read.
3. For each cell `(i, j)`:
   - Its flat index is `flat = i*n + j`.
   - Its destination flat index is `dest = (flat + k) % total`.
   - Write the value into `result[dest/n][dest%n]`.
4. Return `result`.

Reducing `k` by `k % total` is optional for correctness here because we already take the modulo when computing `dest`, but it keeps the arithmetic small and makes the "full rotation is a no-op" property explicit.

Walking through Example 1 with `grid = [[1,2,3],[4,5,6],[7,8,9]]`, `k = 1`, `n = 3`, `total = 9`: value `9` is at `(2,2)`, flat index `8`; its destination is `(8+1) % 9 = 0`, i.e. `(0,0)`. Value `1` at flat `0` goes to flat `1`, i.e. `(0,1)`. Doing this for every cell produces `[[9,1,2],[3,4,5],[6,7,8]]`, matching the expected output.

Using a separate output grid avoids the classic in-place hazard of clobbering a cell before it has been moved, and computing destinations directly avoids simulating `k` full passes.

## Complexity Analysis

Time Complexity: O(m * n) — every cell is read once and written once, independent of `k`.
Space Complexity: O(m * n) — for the newly built result grid (O(1) extra beyond the required output).

## Edge Cases

- **k is 0:** No shift should occur; the modulo arithmetic yields `dest = flat`, returning an identical grid.
- **k is a multiple of m * n (e.g. Example 3):** A full number of rotations returns the grid to its original state. `(flat + k) % total == flat`, so the output equals the input.
- **k larger than m * n:** Naively looping `k` times wastes work; the modulo reduces `k` into the meaningful range `[0, total)`.
- **Single row (m == 1):** Degrades to a plain cyclic shift of one array; the same index math handles it.
- **Single column (n == 1):** Every shift moves a value down one row and wraps the bottom to the top; `dest/n` and `dest%n` still resolve correctly since `n == 1`.
- **1x1 grid:** `total == 1`, so every destination is `(flat + k) % 1 == 0`; the single element stays put for any `k`.
