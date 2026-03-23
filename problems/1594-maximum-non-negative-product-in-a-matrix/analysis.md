# 1594. Maximum Non Negative Product in a Matrix

[LeetCode Link](https://leetcode.com/problems/maximum-non-negative-product-in-a-matrix/)

Difficulty: Medium
Topics: Array, Dynamic Programming, Matrix
Acceptance Rate: 38.7%

## Hints

### Hint 1

This is a grid path problem where you can only move right or down. Think about how dynamic programming applies to grid traversal — what state do you need to track at each cell?

### Hint 2

Since the grid contains negative numbers, the product along a path can flip sign. A standard DP that only tracks the maximum product won't work. Consider tracking both the maximum and minimum product ending at each cell — why might the minimum be useful?

### Hint 3

When you multiply by a negative number, the largest product becomes the smallest and vice versa. So at each cell, the new maximum product is either `max_from_neighbor * grid[i][j]` or `min_from_neighbor * grid[i][j]`, depending on the sign of `grid[i][j]`. Track both extremes at every cell, and the answer is the maximum product at the bottom-right corner (if non-negative).

## Approach

Use dynamic programming where each cell `(i, j)` stores two values:

- `maxProd[i][j]`: the maximum product of any path from `(0,0)` to `(i,j)`
- `minProd[i][j]`: the minimum product of any path from `(0,0)` to `(i,j)`

**Base case:** `maxProd[0][0] = minProd[0][0] = grid[0][0]`.

**First row and first column:** There's only one path (all right or all down), so the products are cumulative.

**Transition:** For each cell `(i, j)`, consider the candidates coming from the cell above `(i-1, j)` and the cell to the left `(i, j-1)`. For each neighbor, multiplying by `grid[i][j]` gives two candidate values. If `grid[i][j] >= 0`, the max stays max and min stays min. If `grid[i][j] < 0`, the max and min swap. Take the overall max of all candidates for `maxProd[i][j]` and the overall min for `minProd[i][j]`.

**Result:** If `maxProd[m-1][n-1] < 0`, return `-1`. Otherwise return `maxProd[m-1][n-1] % (10^9 + 7)`.

**Why track both?** Example 2 illustrates this: the path `1 → 1 → -2 → -4 → 1` has two negatives that cancel out. The intermediate minimum product (which is negative) becomes the maximum after multiplying by the second negative value.

Since grid values range from -4 to 4 and paths are at most 29 cells long (15+15-1), the product can be as large as 4^29, which fits comfortably in a 64-bit integer — no need for big integers during DP.

## Complexity Analysis

Time Complexity: O(m * n) — we visit each cell once and do constant work per cell.

Space Complexity: O(m * n) for the two DP tables. This can be optimized to O(n) by processing row by row, but O(m * n) is clean and sufficient given the small constraints (m, n ≤ 15).

## Edge Cases

- **All negative paths:** Every path from top-left to bottom-right yields a negative product. Return -1. (Example 1)
- **Zero in the path:** A path containing 0 gives product 0, which is non-negative. This can be the answer when all other paths are negative. (Example 3)
- **Single cell grid (1×1):** The product is just that cell's value. If negative, return -1.
- **Large positive products:** With values up to 4 and paths up to 29 cells, the product can exceed 32-bit range. Use `int64` (Go's `int` is 64-bit on most platforms, but being explicit is safer).
- **Grid with all zeros:** The answer is 0.
