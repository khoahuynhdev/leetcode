# 1582. Special Positions in a Binary Matrix

[LeetCode Link](https://leetcode.com/problems/special-positions-in-a-binary-matrix/)

Difficulty: Easy
Topics: Array, Matrix
Acceptance Rate: 69.3%

## Hints

### Hint 1

Think about what makes a position "special." You need to verify a property about an entire row and an entire column. Can you precompute something about each row and column to avoid repeatedly scanning them?

### Hint 2

If you count how many 1s are in each row and each column ahead of time, you can determine whether a cell is the *only* 1 in its row and column in constant time per cell.

### Hint 3

A cell `(i, j)` is special if and only if `mat[i][j] == 1`, the total number of 1s in row `i` is exactly 1, and the total number of 1s in column `j` is exactly 1. Precompute row sums and column sums in a single pass, then check each 1-valued cell against those sums.

## Approach

1. **Precompute row and column sums.** Iterate through the entire matrix once. For each cell `(i, j)`, add `mat[i][j]` to `rowSum[i]` and `colSum[j]`.

2. **Count special positions.** Iterate through the matrix a second time. For each cell where `mat[i][j] == 1`, check whether `rowSum[i] == 1` and `colSum[j] == 1`. If both conditions hold, this cell is the only 1 in its row and the only 1 in its column, so it is a special position. Increment the count.

**Why this works:** A special position requires `mat[i][j] == 1` with all other elements in row `i` and column `j` equal to 0. If the sum of row `i` is 1 and cell `(i, j)` is 1, then every other cell in that row must be 0. The same logic applies to the column.

**Example walkthrough with `[[1,0,0],[0,0,1],[1,0,0]]`:**
- Row sums: [1, 1, 1], Column sums: [2, 0, 1]
- Cell (0,0): mat=1, rowSum=1, colSum=2 → not special (column has another 1)
- Cell (1,2): mat=1, rowSum=1, colSum=1 → special
- Cell (2,0): mat=1, rowSum=1, colSum=2 → not special
- Answer: 1

## Complexity Analysis

Time Complexity: O(m × n) — two passes over the matrix.
Space Complexity: O(m + n) — for the row sum and column sum arrays.

## Edge Cases

- **Single cell matrix `[[1]]`**: The only element is 1 with no other elements in its row or column, so it is special. Result: 1.
- **Single cell matrix `[[0]]`**: The only element is 0, so no special position exists. Result: 0.
- **All zeros**: No cell has value 1, so the answer is 0.
- **All ones**: Every row sum and column sum exceeds 1, so no position is special. Result: 0.
- **Single row or single column**: Reduces to checking if there is exactly one 1 in the entire vector.
