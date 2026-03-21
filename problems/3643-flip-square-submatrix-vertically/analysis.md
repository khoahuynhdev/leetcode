# 3643. Flip Square Submatrix Vertically

[LeetCode Link](https://leetcode.com/problems/flip-square-submatrix-vertically/)

Difficulty: Easy
Topics: Array, Two Pointers, Matrix
Acceptance Rate: 74.0%

## Hints

### Hint 1

Think about what "flipping vertically" means for a submatrix. You are not rotating or transposing — you are only reversing the order of rows within a specific region. Consider how you would reverse elements in an array using two pointers.

### Hint 2

You can use a two-pointer technique on the rows of the submatrix. Start with one pointer at the top row of the submatrix and another at the bottom row. Swap the relevant columns of these two rows and move the pointers inward. Only the columns within the submatrix range (from `y` to `y+k-1`) need to be swapped.

### Hint 3

The swap is done in-place. For each pair of rows `(x + i, x + k - 1 - i)` where `i < k/2`, swap the values in columns `y` through `y + k - 1`. That is all you need — no extra data structures required.

## Approach

1. Identify the square submatrix defined by the top-left corner `(x, y)` and side length `k`. The rows involved are `x` to `x+k-1`, and the columns are `y` to `y+k-1`.
2. Use a two-pointer approach: set `top = x` and `bottom = x + k - 1`.
3. While `top < bottom`:
   - For each column `j` from `y` to `y+k-1`, swap `grid[top][j]` and `grid[bottom][j]`.
   - Increment `top`, decrement `bottom`.
4. Return the modified grid.

**Example walkthrough (Example 1):**
- Grid is 4×4, `x=1, y=0, k=3`. The submatrix covers rows 1–3, columns 0–2.
- Swap row 1 cols 0–2 with row 3 cols 0–2: `[5,6,7]` ↔ `[13,14,15]`.
- Row 2 is the middle row and stays in place.
- Result: `[[1,2,3,4],[13,14,15,8],[9,10,11,12],[5,6,7,16]]`.

## Complexity Analysis

Time Complexity: O(k²) — we swap at most k/2 rows, each with k columns.
Space Complexity: O(1) — all swaps are done in-place.

## Edge Cases

- **k = 1**: The submatrix is a single cell. No swaps are needed; the grid is returned unchanged.
- **k = 2**: Only one pair of rows to swap — the simplest non-trivial case.
- **Submatrix at the edge of the grid**: When `x + k - 1` equals the last row or `y + k - 1` equals the last column, ensure you don't go out of bounds (the constraints guarantee validity).
- **Entire grid is the submatrix**: When `x=0, y=0, k=min(m,n)` and the grid is square, the entire grid gets flipped vertically.
