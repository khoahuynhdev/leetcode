# 3212. Count Submatrices With Equal Frequency of X and Y

[LeetCode Link](https://leetcode.com/problems/count-submatrices-with-equal-frequency-of-x-and-y/)

Difficulty: Medium
Topics: Array, Matrix, Prefix Sum
Acceptance Rate: 56.0%

## Hints

### Hint 1

The submatrices must always include `grid[0][0]` as the top-left corner. This means every valid submatrix is defined by choosing a bottom-right corner `(i, j)`. Think about how you can efficiently count the number of `'X'` and `'Y'` characters in each such submatrix.

### Hint 2

Since every submatrix shares the same top-left corner `(0, 0)`, you can use a 2D prefix sum to compute the count of `'X'` and `'Y'` in the rectangle from `(0, 0)` to `(i, j)` in O(1) time after an O(m*n) preprocessing step. You need two separate prefix sum arrays — one for `'X'` and one for `'Y'`.

### Hint 3

Build prefix sum arrays `prefX[i][j]` and `prefY[i][j]` where each stores the count of the respective character in the submatrix from `(0,0)` to `(i-1, j-1)`. For each cell `(i, j)`, check if `prefX[i+1][j+1] == prefY[i+1][j+1]` and `prefX[i+1][j+1] > 0`. If both conditions hold, that bottom-right corner produces a valid submatrix.

## Approach

1. **Build two 2D prefix sum arrays** — one for counting `'X'` and one for counting `'Y'`. Use 1-indexed prefix sums for cleaner boundary handling. The recurrence is:

   ```
   prefX[i][j] = prefX[i-1][j] + prefX[i][j-1] - prefX[i-1][j-1] + (grid[i-1][j-1] == 'X' ? 1 : 0)
   ```

   Same formula applies for `prefY` with `'Y'`.

2. **Iterate over every cell** `(i, j)` in the grid. For each cell, query the prefix sums to get `countX = prefX[i+1][j+1]` and `countY = prefY[i+1][j+1]`.

3. **Check the two conditions**: `countX == countY` and `countX > 0`. If both hold, increment the answer.

4. Return the total count.

**Example walkthrough** with `grid = [["X","Y","."],["Y",".","."]`]:

- Submatrix ending at (0,1): contains 1 X and 1 Y → valid.
- Submatrix ending at (0,2): contains 1 X and 1 Y → valid.
- Submatrix ending at (1,1): contains 1 X and 2 Y → invalid.
- Submatrix ending at (1,2): contains 1 X and 2 Y → invalid.
- Submatrix ending at (1,0): contains 1 X and 1 Y → valid.

Total = 3.

## Complexity Analysis

Time Complexity: O(m * n) — one pass to build prefix sums, one pass to count valid submatrices.
Space Complexity: O(m * n) — for the two prefix sum arrays.

## Edge Cases

- **Grid with no `'X'` at all**: Every submatrix has `countX == 0`, so the "at least one X" condition is never met. Answer is 0.
- **Grid with no `'Y'`**: `countX > countY` for every non-trivial submatrix containing an `'X'`, so the equal-frequency condition is never met. Answer is 0.
- **Grid entirely of `'.'`**: Same as no `'X'` — answer is 0.
- **1x1 grid with `'X'`**: The only submatrix has 1 X and 0 Y — not equal. Answer is 0.
- **1x1 grid with `'.'`**: No X present. Answer is 0.
- **Large grid (1000x1000)**: The O(m*n) approach handles this efficiently within constraints.
