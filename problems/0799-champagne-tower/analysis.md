# 0799. Champagne Tower

[LeetCode Link](https://leetcode.com/problems/champagne-tower/)

Difficulty: Medium
Topics: Dynamic Programming
Acceptance Rate: 59.0%

## Hints

### Hint 1

Think of the champagne tower as a triangle (2D grid). Instead of tracking how much each glass "receives," consider simulating the pour from the top down, tracking the total liquid that flows through each position.

### Hint 2

Use a simulation approach: maintain an array (or 2D array) where each entry tracks how much liquid has been poured into that glass so far. If a glass has more than 1 cup, the excess spills equally to the two glasses below it. Process row by row from top to bottom.

### Hint 3

The key insight is that you only need to track the *total* liquid at each position, not just what it holds. A glass at position `(i, j)` receives overflow from `(i-1, j-1)` and `(i-1, j)`. The overflow from a glass is `max(0, total - 1) / 2`. The answer for any glass is `min(1, total)` since a glass can hold at most 1 cup. You can optimize space to use only two rows at a time.

## Approach

We simulate the champagne pour using a row-by-row DP approach:

1. Start by placing all `poured` cups into position `(0, 0)`.
2. For each row from `0` to `query_row - 1`:
   - For each glass `j` in the current row, if the glass has more than 1 cup of liquid, compute the overflow: `overflow = (amount[j] - 1.0) / 2.0`.
   - Distribute this overflow equally to positions `j` and `j+1` in the next row.
   - Cap the current glass at 1.0 (it can only hold 1 cup).
3. After processing all rows up to `query_row`, the answer is `min(1.0, amount[query_glass])`.

We optimize space by noting that when computing row `i+1`, we only need the values from row `i`. So we use a single 1D slice of size `query_row + 1` and process it carefully, or use two alternating rows.

A simpler approach uses a single 2D simulation array of size `(query_row+1) x (query_row+2)`, which is at most 100x101 — well within limits.

## Complexity Analysis

Time Complexity: O(query_row^2) — we process each glass in each row up to `query_row`, and each row `i` has `i+1` glasses.

Space Complexity: O(query_row^2) — for the 2D simulation array. This can be optimized to O(query_row) with a rolling array.

## Edge Cases

- **poured = 0**: No champagne is poured; every glass is empty (returns 0.0).
- **query_row = 0, query_glass = 0**: Only the top glass matters; result is `min(1.0, poured)`.
- **Large poured value (e.g., 10^9)**: The glass at a deep enough row will be full (1.0). Floating-point precision is not an issue since we only need 5 decimal places.
- **query_glass at the edge of a row**: Edge glasses receive overflow from only one parent, not two.
