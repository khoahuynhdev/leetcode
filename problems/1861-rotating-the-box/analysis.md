# 1861. Rotating the Box

[LeetCode Link](https://leetcode.com/problems/rotating-the-box/)

Difficulty: Medium
Topics: Array, Two Pointers, Matrix
Acceptance Rate: 79.8%

## Hints

### Hint 1

The problem combines two operations: gravity simulation and a 90° clockwise rotation. Try to think about whether you really need to rotate first and *then* apply gravity, or whether you can decouple the two by reasoning about how gravity will act in the rotated frame.

### Hint 2

After a 90° clockwise rotation, the original right edge of each row becomes the bottom edge of a column. So gravity in the rotated box pulls stones toward the original right of each row. This means you can simulate gravity *before* rotating, working row by row in the original matrix — no extra dimension juggling required.

### Hint 3

For each row, scan from right to left with a two-pointer trick: keep an `empty` index pointing at the rightmost slot a stone could currently fall into. When you see a stone, move it to `empty` and decrement; when you see an obstacle `*`, reset `empty` to the cell just to the left of the obstacle, since stones cannot pass through it. After all rows are settled, do a single 90° clockwise rotation: `result[i][j] = box[m-1-j][i]`.

## Approach

The clever observation is that we can avoid actually simulating physics in the rotated frame. Because a 90° clockwise rotation maps "right" to "down", a stone that ends up at the bottom of a column in the rotated box corresponds to a stone that ended up at the right of its row in the original box. So we can:

1. **Apply gravity row-by-row in the original matrix.** For each row, walk from right to left with a pointer `empty` that tracks where the next stone should land. Initialize `empty = n - 1`.
   - If the current cell is `'.'`, do nothing.
   - If the current cell is `'#'` (a stone), move it to position `empty` (clear the current cell, set `box[i][empty] = '#'`), then decrement `empty`.
   - If the current cell is `'*'` (an obstacle), it blocks all stones to its left. Reset `empty = j - 1` so that the next stones pile up on the left side of this obstacle.

2. **Rotate 90° clockwise.** Build a new `n × m` matrix where `result[i][j] = box[m-1-j][i]`. This single index expression captures the rotation: the original bottom-left corner becomes the new bottom-right corner, and so on.

Worked example with row `# . *  .`:
- Start `empty = 3`. Cell 3 = `.`, skip.
- Cell 2 = `*`, set `empty = 1`.
- Cell 1 = `.`, skip.
- Cell 0 = `#`, move to position 1: row becomes `. # * .`, `empty = 0`.

After rotation, the column `[# .]` (top-to-bottom of original right column became left of new) lands correctly with the stone at the bottom.

## Complexity Analysis

Time Complexity: O(m * n) — each cell is visited once during the gravity pass and once during the rotation.
Space Complexity: O(m * n) for the rotated output matrix. The gravity step itself is O(1) extra space (in-place on the input).

## Edge Cases

- **Single row (`m = 1`)**: still works; gravity collapses the row toward the right, and rotation produces an `n × 1` column.
- **Single column (`n = 1`)**: each row has only one cell, so gravity is a no-op; rotation produces a `1 × m` row.
- **No stones**: gravity does nothing; rotation just transposes/reflects the empties and obstacles.
- **No obstacles**: every stone in a row falls all the way to the right (becomes the bottom of a column after rotation).
- **Obstacles at the right edge**: the `empty` pointer resets to `j - 1`, correctly preventing stones from passing through.
- **Adjacent stones**: stones stack on top of each other naturally because each one decrements `empty` after landing.
