# 3546. Equal Sum Grid Partition I

[LeetCode Link](https://leetcode.com/problems/equal-sum-grid-partition-i/)

Difficulty: Medium
Topics: Array, Matrix, Enumeration, Prefix Sum
Acceptance Rate: 46.6%

## Hints

### Hint 1

Think about what it means to split a grid into two equal parts. What must be true about the total sum of the grid for any valid partition to exist?

### Hint 2

If the total sum is even, you need to find a cut (horizontal or vertical) where the running sum of complete rows or complete columns equals exactly half the total. Consider using prefix sums along rows and columns separately.

### Hint 3

For a horizontal cut after row `i`, the sum of all elements in rows `0..i` must equal half the total. Similarly, for a vertical cut after column `j`, the sum of all elements in columns `0..j` must equal half the total. You only need to check `m-1` possible horizontal cuts and `n-1` possible vertical cuts — never the last row/column since both sections must be non-empty.

## Approach

1. **Compute the total sum** of the entire grid. If the total is odd, immediately return `false` — you cannot split an odd number into two equal integers.

2. **Check horizontal cuts.** Iterate through rows `0` to `m-2`. Maintain a running sum of all elements seen so far (full rows). After adding each row, check if the running sum equals `totalSum / 2`. If yes, return `true`.

3. **Check vertical cuts.** Iterate through columns `0` to `n-2`. Maintain a running sum of all elements in columns `0..j` across every row. After processing each column, check if the running sum equals `totalSum / 2`. If yes, return `true`.

4. If no cut works, return `false`.

**Example walkthrough** with `grid = [[1,4],[2,3]]`:
- Total sum = 1 + 4 + 2 + 3 = 10, half = 5.
- Horizontal cut after row 0: sum of row 0 = 1 + 4 = 5 = half. Return `true`.

## Complexity Analysis

Time Complexity: O(m * n) — we iterate over every element at most twice (once for horizontal checks, once for vertical).

Space Complexity: O(1) — only a few running totals are maintained; no extra data structures needed.

## Edge Cases

- **Odd total sum**: No valid partition can exist; return `false` immediately.
- **Single row grid (m=1)**: Only vertical cuts are possible; horizontal cuts are skipped since you cannot split a single row horizontally into two non-empty parts.
- **Single column grid (n=1)**: Only horizontal cuts are possible.
- **Minimum grid (1x2 or 2x1)**: Exactly one possible cut to check.
- **All elements equal**: A cut works if and only if the number of rows (for horizontal) or columns (for vertical) is even.
