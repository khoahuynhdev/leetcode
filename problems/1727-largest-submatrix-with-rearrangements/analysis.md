# 1727. Largest Submatrix With Rearrangements

[LeetCode Link](https://leetcode.com/problems/largest-submatrix-with-rearrangements/)

Difficulty: Medium
Topics: Array, Greedy, Sorting, Matrix
Acceptance Rate: 76.1%

## Hints

### Hint 1

Think about how you would solve the "largest rectangle in a histogram" problem. Can you transform this matrix into something that looks like histograms? For each cell, consider how many consecutive 1s exist above it (including itself).

### Hint 2

For each row, compute the "height" of consecutive 1s ending at that row for every column. Since you can rearrange columns freely, you can sort these heights within each row. After sorting, the problem reduces to finding the largest rectangle you can form with sorted bar heights.

### Hint 3

After sorting each row's heights in descending order, iterate through the sorted heights. At position `j` (0-indexed), you know that at least `j+1` columns have a height of at least `heights[j]`. So the area of the rectangle you can form is `heights[j] * (j + 1)`. Track the maximum area across all rows and all positions.

## Approach

The algorithm works in two phases:

**Phase 1 — Build height matrix:** For each column, compute the number of consecutive 1s ending at each row. Walk top to bottom: if `matrix[i][j] == 1`, set `matrix[i][j] = matrix[i-1][j] + 1`; otherwise leave it as 0. This transforms the binary matrix into a histogram-height matrix in place.

For example, given:
```
0 0 1
1 1 1
1 0 1
```
The height matrix becomes:
```
0 0 1
1 1 2
2 0 3
```

**Phase 2 — Sort and compute max area per row:** For each row, copy the heights into an array and sort it in descending order. Then iterate through the sorted array: at index `j`, the rectangle width is `j+1` and the height is `heights[j]` (since all previous bars are at least as tall). Compute `area = heights[j] * (j + 1)` and update the global maximum.

For row 2 of our example: heights = [2, 0, 3] → sorted descending = [3, 2, 0].
- j=0: area = 3 * 1 = 3
- j=1: area = 2 * 2 = 4 ← maximum
- j=2: area = 0 * 3 = 0

The answer is 4.

## Complexity Analysis

Time Complexity: O(m * n * log n) — We iterate over every cell once to build heights (O(m*n)), then sort each row (O(n log n)) for m rows.

Space Complexity: O(n) — We use a temporary array of length n for sorting each row's heights. The height computation is done in place.

## Edge Cases

- **Single row:** The answer is simply the count of 1s in that row, since all 1-columns can be rearranged adjacent to each other.
- **Single column:** No rearrangement is possible; the answer is the longest streak of consecutive 1s.
- **All zeros:** The answer is 0.
- **All ones:** The answer is m * n, the entire matrix.
- **Very large m with small n (or vice versa):** The constraint `1 <= m * n <= 10^5` means dimensions can be skewed. The algorithm handles this naturally.
