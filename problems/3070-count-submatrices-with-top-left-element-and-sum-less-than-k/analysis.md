# 3070. Count Submatrices with Top-Left Element and Sum Less Than k

[LeetCode Link](https://leetcode.com/problems/count-submatrices-with-top-left-element-and-sum-less-than-k/)

Difficulty: Medium
Topics: Array, Matrix, Prefix Sum
Acceptance Rate: 62.8%

## Hints

### Hint 1

Every submatrix that contains the top-left element of the grid is uniquely determined by one thing — its bottom-right corner. Think about what that means for how many submatrices you actually need to consider.

### Hint 2

If you can quickly compute the sum of any submatrix from (0, 0) to (i, j), you can check each candidate in O(1) time. A classic technique lets you build such a lookup over the entire grid in a single pass. Look into 2D prefix sums.

### Hint 3

Because all grid values are non-negative, the prefix sums are monotonically non-decreasing along each row (left to right) and each column (top to bottom). This means once a submatrix sum exceeds k for some column j in row i, every wider submatrix in that row will also exceed k — so you can break out of the inner loop early.

## Approach

Since every valid submatrix must include the top-left corner (0, 0), each submatrix is fully defined by its bottom-right corner (i, j). The sum of the submatrix from (0, 0) to (i, j) can be computed using the 2D prefix sum formula:

```
prefix[i][j] = grid[i][j]
             + prefix[i-1][j]   (sum of everything above)
             + prefix[i][j-1]   (sum of everything to the left)
             - prefix[i-1][j-1] (subtract double-counted region)
```

We compute the prefix sums in-place (or in a separate array) row by row, left to right. For each cell (i, j), if `prefix[i][j] <= k`, we increment our count. Because all values are non-negative, prefix sums only grow as j increases within a row, so we can break early when the sum exceeds k.

**Walkthrough with Example 1** (`grid = [[7,6,3],[6,6,1]], k = 18`):

| (i,j) | prefix | <= 18? |
|--------|--------|--------|
| (0,0)  | 7      | yes    |
| (0,1)  | 13     | yes    |
| (0,2)  | 16     | yes    |
| (1,0)  | 13     | yes    |
| (1,1)  | 25     | no — break row |

Count = 4, which matches the expected output.

## Complexity Analysis

Time Complexity: O(m × n) — we visit each cell once to compute the prefix sum and check the condition.

Space Complexity: O(1) if we modify the grid in-place, or O(m × n) if we use a separate prefix sum array. The solution below modifies the grid in-place.

## Edge Cases

- **Single cell grid**: The answer is 1 if `grid[0][0] <= k`, otherwise 0.
- **First cell exceeds k**: If `grid[0][0] > k`, no submatrix can satisfy the condition, return 0 immediately.
- **All zeros**: Every submatrix has sum 0 which is always <= k, so the answer is m × n.
- **Large k**: If k is large enough that the total sum of the grid fits, the answer is m × n.
- **Single row or single column**: Degenerates to a 1D prefix sum problem; the early-break optimization still applies.
