# 1886. Determine Whether Matrix Can Be Obtained By Rotation

[LeetCode Link](https://leetcode.com/problems/determine-whether-matrix-can-be-obtained-by-rotation/)

Difficulty: Easy
Topics: Array, Matrix
Acceptance Rate: 60.6%

## Hints

### Hint 1

Think about how many distinct orientations a matrix can have when rotated in 90-degree increments. You don't need to consider infinite rotations — the rotations eventually cycle back to the original.

### Hint 2

There are only 4 possible rotations: 0°, 90°, 180°, and 270°. If you can figure out the formula for rotating a matrix 90° clockwise, you can apply it repeatedly and compare each result to the target.

### Hint 3

When rotating an n×n matrix 90° clockwise, element at position `(i, j)` moves to position `(j, n-1-i)`. You can either rotate in-place or create a new matrix each time. After each rotation, check if the matrix equals the target.

## Approach

The key observation is that rotating a matrix 90° clockwise four times brings it back to the original orientation. So we only need to check at most 4 rotations (0°, 90°, 180°, 270°).

**Algorithm:**

1. Check if `mat` already equals `target` (0° rotation). If yes, return `true`.
2. Rotate `mat` by 90° clockwise. Check if it equals `target`. If yes, return `true`.
3. Repeat step 2 two more times (for 180° and 270°).
4. If none of the 4 orientations match, return `false`.

**Rotating 90° clockwise:** For an n×n matrix, the element at `(i, j)` moves to `(j, n-1-i)`. We can build a new rotated matrix where `rotated[j][n-1-i] = mat[i][j]`.

**Example walkthrough (Example 1):**
- `mat = [[0,1],[1,0]]`, `target = [[1,0],[0,1]]`
- 0° rotation: `[[0,1],[1,0]]` ≠ target
- 90° rotation: `[[1,0],[0,1]]` = target → return `true`

## Complexity Analysis

Time Complexity: O(n²) — We perform at most 4 rotations, each requiring O(n²) work to rotate and O(n²) to compare. Since 4 is a constant, overall is O(n²).

Space Complexity: O(n²) — We create a new n×n matrix for each rotation.

## Edge Cases

- **Already equal (0° rotation):** `mat` and `target` are identical — should return `true` without any rotation.
- **1×1 matrix:** A single-element matrix is always equal to itself after any rotation. Return `true` if `mat[0][0] == target[0][0]`, otherwise `false`.
- **All same values:** If both matrices are all 0s or all 1s, any rotation matches — return `true`.
- **No rotation matches:** After checking all 4 rotations, none equal the target — return `false`.
