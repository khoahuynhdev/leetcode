# 1878. Get Biggest Three Rhombus Sums in a Grid

[LeetCode Link](https://leetcode.com/problems/get-biggest-three-rhombus-sums-in-a-grid/)

Difficulty: Medium
Topics: Array, Math, Sorting, Heap (Priority Queue), Matrix, Prefix Sum
Acceptance Rate: 55.3%

## Hints

### Hint 1

Think about what defines a rhombus on a grid. A rhombus is a diamond shape — a square rotated 45 degrees. Its border can be described by a center point and a "size" (the distance from center to any corner). How would you enumerate all possible rhombuses in the grid?

### Hint 2

For each cell in the grid, treat it as the center of a rhombus and try all valid sizes (including size 0, which is just the cell itself). For a rhombus of size `s` centered at `(r, c)`, the four corners are at `(r-s, c)`, `(r, c+s)`, `(r+s, c)`, and `(r, c-s)`. You need to sum the elements along the four sides of the border. Think about how to walk along each side using a loop variable `i` from `0` to `s-1`.

### Hint 3

The key insight is that you can decompose the rhombus border into four sides, each traversed by incrementing one coordinate and decrementing the other. For each side, iterate `i` from `0` to `s-1` to avoid double-counting corners. Collect all distinct rhombus sums, then return the top three. Given the small constraints (m, n <= 50), a brute-force enumeration of all centers and sizes is efficient enough — no prefix sum optimization is needed.

## Approach

1. **Enumerate all rhombuses**: For every cell `(r, c)` in the grid, consider it as the center of a rhombus. Try every valid size `s` starting from 0:
   - **Size 0**: The rhombus sum is just `grid[r][c]` (a single cell).
   - **Size s >= 1**: The rhombus has four corners at `(r-s, c)`, `(r, c+s)`, `(r+s, c)`, `(r, c-s)`. All four corners must be within bounds.

2. **Sum the border**: For a rhombus of size `s`, walk along the four sides. Each side is traversed with index `i` from `0` to `s-1`, carefully chosen so that corners are counted exactly once:
   - Top to Right: `(r-s+i, c+i)`
   - Right to Bottom: `(r+i, c+s-i)`
   - Bottom to Left: `(r+s-i, c-i)`
   - Left to Top: `(r-i, c-s+i)`

3. **Track distinct sums**: Use a set (map) to store all distinct rhombus sums.

4. **Return top 3**: Sort the distinct sums in descending order and return the first three (or fewer if there aren't three distinct values).

## Complexity Analysis

Time Complexity: O(m * n * min(m, n) * min(m, n)) — for each cell we try all sizes up to min(m, n)/2, and each rhombus border has O(s) elements. In the worst case with m = n = 50, this is roughly 50 * 50 * 25 * 25 = ~1.5M operations, which is very fast.

Space Complexity: O(m * n * min(m, n)) — for storing distinct sums in the set. In practice, far fewer distinct sums exist.

## Edge Cases

- **Single cell grid (1x1)**: Only one rhombus of size 0 exists. Return that single value.
- **Single row or column**: Only size-0 rhombuses are possible since no larger rhombus can fit.
- **All identical values**: All rhombus sums may be distinct multiples of the value, or many may coincide. Return at most 3 distinct values.
- **Fewer than 3 distinct sums**: Return all distinct sums, even if fewer than 3.
- **Large rhombus spanning the entire grid**: Ensure bounds checking correctly allows the maximum possible rhombus size.
