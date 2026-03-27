# 2946. Matrix Similarity After Cyclic Shifts

[LeetCode Link](https://leetcode.com/problems/matrix-similarity-after-cyclic-shifts/)

Difficulty: Easy
Topics: Array, Math, Matrix, Simulation
Acceptance Rate: 63.9%

## Hints

### Hint 1

Think about what happens when you cyclically shift a row multiple times. How does the number of columns relate to when a row returns to its original state?

### Hint 2

You don't need to actually simulate all k shifts. A cyclic left shift by k on a row of length n is equivalent to a single shift by `k % n`. The same applies to right shifts. Can you check each row directly?

### Hint 3

The key insight is that for a row to look the same after being shifted by k positions (whether left or right), every element must equal the element k positions ahead (cyclically): `mat[i][j] == mat[i][(j+k) % n]` for all j. This condition is identical for both left and right shifts, so you don't even need to distinguish between even and odd rows.

## Approach

Since cyclic shifting is periodic with period n (the number of columns), shifting by k is the same as shifting by `k % n`. For the matrix to be identical after the shifts, every row must be unchanged.

A row is unchanged after a left cyclic shift of k positions if and only if `row[j] == row[(j + k) % n]` for every column j. The same condition holds for right shifts (the equivalence follows from substitution).

This means we don't need to differentiate between even and odd rows at all — we just check the same condition for every row. If any element violates this, return false.

Algorithm:
1. Let n = number of columns.
2. Compute `shift = k % n`. If shift is 0, return true immediately.
3. For each row, for each column j, check if `mat[i][j] == mat[i][(j + shift) % n]`.
4. If all checks pass, return true; otherwise return false.

## Complexity Analysis

Time Complexity: O(m * n) where m is the number of rows and n is the number of columns. We visit each element once.

Space Complexity: O(1) — no extra space beyond loop variables.

## Edge Cases

- **k is a multiple of n**: The shift amount is effectively 0, so the matrix is always identical. The `k % n == 0` check handles this.
- **All elements in a row are the same**: Any shift leaves the row unchanged, so it always returns true for that row.
- **Single column matrix**: Any shift on a single-element row is a no-op, always true.
- **k larger than n**: We must use `k % n` rather than k directly to avoid unnecessary work or index errors.
