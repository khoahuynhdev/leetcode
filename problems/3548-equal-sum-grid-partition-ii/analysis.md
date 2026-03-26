# 3548. Equal Sum Grid Partition II

[LeetCode Link](https://leetcode.com/problems/equal-sum-grid-partition-ii/)

Difficulty: Hard
Topics: Array, Hash Table, Matrix, Enumeration, Prefix Sum
Acceptance Rate: 29.7%

## Hints

### Hint 1

Think about what happens when you make a single horizontal or vertical cut through the grid. Each cut produces two rectangular sections. How can you efficiently compute the sum of each section as you sweep the cut position across the grid?

### Hint 2

Use prefix sums to compute section sums in O(1) per cut. When sums aren't equal, the difference tells you the exact value of the cell you'd need to discount. Maintain a frequency map (hash set) of values in each section so you can check in O(1) whether such a cell exists. As you sweep the cut position, incrementally update both the prefix sum and the frequency map.

### Hint 3

The critical insight is the connectivity constraint. When you remove a cell from a rectangular section, the rest must stay connected. For a section with at least 2 rows AND at least 2 columns, removing any single cell always leaves a connected region (you can route around the gap). But for a single-row section, only the two endpoint cells can be safely removed (removing an interior cell splits the row). Similarly for a single-column section. This dramatically limits which cells you need to check in thin sections.

## Approach

1. **Compute the total sum** of the grid.

2. **Sweep horizontal cuts** (after each row r, from r=0 to m-2):
   - Maintain a running `topSum` and frequency maps for both sections.
   - After adding row r to the top section, check:
     - If `topSum == bottomSum`, return true immediately.
     - Otherwise, compute `diff = largerSum - smallerSum`. We need a cell with value `diff` in the larger section that can be removed while keeping that section connected.
   - **Connectivity check for the larger section** (which is a rectangle of height h and width n):
     - If h >= 2 and n >= 2: any cell is removable, so just check if `diff` exists in the frequency map.
     - If h == 1 (single row): only the leftmost and rightmost cells are removable.
     - If n == 1 (single column): only the topmost and bottommost cells are removable.
     - If h == 1 and n == 1 (single cell): cannot remove anything.

3. **Sweep vertical cuts** (after each column c, from c=0 to n-2) with the same logic, but now thin sections are single-column (width=1) or single-row (height=1).

4. If no valid partition is found, return false.

The key optimization is maintaining frequency maps incrementally: as we sweep the cut position, we add values from one row/column to one map and remove them from the other. This keeps the total work proportional to the grid size.

## Complexity Analysis

Time Complexity: O(m * n) — We sweep through all rows once for horizontal cuts and all columns once for vertical cuts, doing O(1) work per cell (hash map operations).

Space Complexity: O(m * n) — For the frequency maps that store counts of all cell values.

## Edge Cases

- **Grid with 2 cells** (1x2 or 2x1): Only one possible cut. Can only succeed if both values are equal, since each section is a single cell and you cannot discount from a single-cell section.
- **Single row grid** (1xn): Only vertical cuts possible. Each cut produces two single-row sections where only endpoint cells can be removed.
- **Single column grid** (mx1): Only horizontal cuts possible. Each cut produces two single-column sections where only endpoint cells can be removed.
- **All cells have the same value**: If the total count is even, a balanced cut exists. If odd, you need to discount one cell (the diff equals that common value).
- **Sums already equal without discounting**: Return true immediately — no connectivity check needed.
- **Diff larger than any cell value**: No valid discount exists for that cut, move on.
