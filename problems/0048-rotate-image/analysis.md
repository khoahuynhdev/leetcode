# 0048. Rotate Image

[LeetCode Link](https://leetcode.com/problems/rotate-image/)

Difficulty: Medium
Topics: Array, Math, Matrix
Acceptance Rate: 79.7%

## Hints

### Hint 1

This is a matrix transformation problem with a hard constraint: you must mutate the matrix in place, so you cannot allocate a second `n x n` grid. Think about whether the rotation can be expressed as a composition of two simpler, in-place operations you already know how to do on a 2D array.

### Hint 2

Two classic in-place primitives on a square matrix are: **transposing** (swapping element `(i, j)` with `(j, i)`) and **reversing a row** (or a column). Try drawing a small `3x3` example and see what happens to the cells when you apply one of these and then the other. Does some combination of these two yield a 90° clockwise rotation?

### Hint 3

A 90° clockwise rotation equals **transpose then reverse each row**. Equivalently, it is **reverse each row then transpose** done in the other order, or **reverse each column then transpose**. Pick the transpose-then-reverse-each-row form and be careful that when you transpose you only iterate the upper triangle (`j > i`), otherwise you swap each pair twice and end up where you started.

## Approach

We rotate 90° clockwise in place using two passes over the `n x n` matrix.

**Step 1 — Transpose.** Walk the upper triangle of the matrix and swap `matrix[i][j]` with `matrix[j][i]` for every `j > i`. The diagonal stays put. After this pass, what used to be column `k` (read top-to-bottom) is now row `k` (read left-to-right).

**Step 2 — Reverse each row.** For each row `i`, swap pairs `(left, right)` working inward until the pointers meet. This flips every row horizontally, which is exactly the second half of the transformation.

Why does transpose + row-reverse equal a 90° clockwise rotation? In the rotated image, the cell at `(i, j)` should hold the original `(n - 1 - j, i)` (the bottom-up reading of the original column `i` becomes the new row `i`). Transposing maps `(n - 1 - j, i)` to `(i, n - 1 - j)`, and reversing row `i` then maps `(i, n - 1 - j)` to `(i, j)`. The two operations compose into the rotation.

Worked example for the `3x3` case `[[1,2,3],[4,5,6],[7,8,9]]`:

```
input          transpose       reverse rows
1 2 3          1 4 7           7 4 1
4 5 6    -->   2 5 8     -->   8 5 2
7 8 9          3 6 9           9 6 3
```

That matches the expected output.

An alternative is the **layer-by-layer 4-way swap**: for each concentric ring, rotate four cells at a time using a temporary scalar. It has the same complexity and uses no extra space, but it is more error-prone with index arithmetic. Transpose-plus-reverse is shorter, easier to verify, and just as fast.

## Complexity Analysis

Time Complexity: O(n^2) — each cell is touched a constant number of times across the transpose and the row reversal.
Space Complexity: O(1) — only a handful of scalar swap temporaries; no auxiliary matrix is allocated.

## Edge Cases

- **`n == 1`**: a single-cell matrix is its own rotation. The transpose loop has nothing to do (`j > i` is never satisfied) and the reverse loop terminates immediately, leaving the matrix untouched, which is correct.
- **`n == 2`**: the smallest non-trivial case. Worth a unit test because off-by-one errors in the transpose bound (`j > i` vs `j >= i`) and in the reverse pointer comparison are easy to make and only show up here or on small even sizes.
- **Even vs odd `n`**: with odd `n` the center cell of each row stays put during reversal (`left == right`), with even `n` every cell gets swapped. The `left < right` loop condition handles both correctly without special casing.
- **Negative values**: values can be negative per the constraints. The algorithm only swaps cells, so the sign of the data is irrelevant — but it is worth a test to confirm we never accidentally treat `0` or negatives as a sentinel.
- **Duplicate values**: nothing in the algorithm depends on uniqueness, but a test with repeated values guards against any future "skip if equal" optimization that would silently break correctness.
