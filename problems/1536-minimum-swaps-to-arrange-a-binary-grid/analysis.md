# 1536. Minimum Swaps to Arrange a Binary Grid

[LeetCode Link](https://leetcode.com/problems/minimum-swaps-to-arrange-a-binary-grid/)

Difficulty: Medium
Topics: Array, Greedy, Matrix
Acceptance Rate: 53.7%

## Hints

### Hint 1

Think about what property each row must have for the grid to be valid (all zeros above the main diagonal). What does "above the main diagonal" mean for a specific row? Focus on **trailing zeros** — the consecutive zeros at the end of each row.

### Hint 2

For row `i` (0-indexed), figure out the minimum number of trailing zeros it needs. Then think about this as a **greedy selection** problem: for each position from top to bottom, find the nearest row below that satisfies the requirement and bubble it up using adjacent swaps.

### Hint 3

Row `i` needs at least `n - 1 - i` trailing zeros. Precompute the trailing zero count for every row. Then iterate from row 0 to row n-1: find the first row at or below position `i` whose trailing zero count meets the requirement, and bubble it up by swapping with its neighbors. The total number of bubble-up swaps is your answer. If no qualifying row can be found for any position, return -1.

## Approach

1. **Precompute trailing zeros**: For each row, count how many consecutive zeros appear from the right end. Store these counts in an array.

2. **Determine requirements**: Row `i` (0-indexed) in a valid grid must have all zeros in columns `i+1` through `n-1`. That means row `i` needs at least `n - 1 - i` trailing zeros.

3. **Greedy placement**: Iterate through each row position `i` from top to bottom:
   - Search downward from position `i` for the first row whose trailing zero count is >= `n - 1 - i`.
   - If found at position `j`, bubble it up to position `i` by performing `j - i` adjacent swaps. Update the trailing zeros array accordingly.
   - If no such row exists, the grid cannot be made valid — return -1.

4. **Return total swaps**.

**Why greedy works**: We process rows top-to-bottom. Row 0 has the strictest requirement (most trailing zeros needed). By always picking the closest qualifying row, we minimize the number of swaps for each position without disturbing already-placed rows.

**Walkthrough with Example 1** (`[[0,0,1],[1,1,0],[1,0,0]]`):
- Trailing zeros: `[0, 1, 2]`
- Row 0 needs >= 2 trailing zeros → found at index 2, bubble up: 2 swaps → `[2, 0, 1]`
- Row 1 needs >= 1 trailing zero → found at index 2, bubble up: 1 swap → `[2, 1, 0]`
- Row 2 needs >= 0 → already satisfied.
- Total: **3 swaps**.

## Complexity Analysis

Time Complexity: O(n^2) — for each of the n rows, we may search and bubble up across at most n positions.

Space Complexity: O(n) — for the trailing zeros array.

## Edge Cases

- **1x1 grid**: Always valid (no cells above the diagonal), return 0.
- **Already valid grid**: No swaps needed, return 0 immediately.
- **Impossible grid**: When multiple rows have the same insufficient trailing zero count (e.g., all rows identical with not enough trailing zeros), return -1.
- **All zeros grid**: Always valid regardless of row order, return 0.
- **Diagonal already has the right structure**: Verify that the greedy approach correctly handles grids that are almost valid and need only a few swaps.
