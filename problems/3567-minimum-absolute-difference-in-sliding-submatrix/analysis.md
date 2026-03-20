# 3567. Minimum Absolute Difference in Sliding Submatrix

[LeetCode Link](https://leetcode.com/problems/minimum-absolute-difference-in-sliding-submatrix/)

Difficulty: Medium
Topics: Array, Sorting, Matrix
Acceptance Rate: 71.8%

## Hints

### Hint 1

Think about what happens when you sort a collection of numbers. Where do the closest pairs end up? This observation lets you avoid checking every pair.

### Hint 2

For each k x k submatrix, extract the elements into a flat list, sort them, and then look at consecutive differences. The minimum absolute difference between any two distinct values in a sorted array is always between adjacent elements.

### Hint 3

The grid dimensions are small (up to 30x30), so a brute-force approach of iterating every submatrix, collecting its k*k elements, sorting, and scanning consecutive differences is efficient enough. The key insight is that sorting reduces the pairwise comparison from O(k^4) to O(k^2 log k).

## Approach

1. **Iterate over all valid top-left corners** `(i, j)` where `0 <= i <= m-k` and `0 <= j <= n-k`.
2. **Collect elements**: For each submatrix starting at `(i, j)`, gather all `k * k` values into a flat slice.
3. **Sort** the slice in ascending order.
4. **Find minimum adjacent difference**: Walk through consecutive pairs in the sorted slice. Track the minimum difference, skipping pairs where both values are equal (difference = 0 — if all values are the same, the result is 0).
5. **Store** the result in `ans[i][j]`.

Since sorting places the closest values next to each other, the minimum absolute difference between any two distinct values must occur between some pair of adjacent elements in the sorted order.

## Complexity Analysis

Time Complexity: O((m - k + 1) * (n - k + 1) * k^2 * log(k^2)) — for each submatrix we collect k^2 elements and sort them.

Space Complexity: O(k^2) — for the temporary slice used to sort each submatrix's elements.

## Edge Cases

- **k = 1**: Every submatrix is a single element, so there are no two distinct values. The answer is 0 for every position.
- **All elements identical in a submatrix**: The minimum absolute difference is 0 since no two distinct values exist.
- **Negative values**: The grid can contain negative numbers; sorting and absolute difference handle this naturally.
- **Minimum grid size (1x1)**: Only one submatrix with one element, result is `[[0]]`.
