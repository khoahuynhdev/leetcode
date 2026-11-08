# 1267. Count Servers that Communicate

[LeetCode Link](https://leetcode.com/problems/count-servers-that-communicate/)

Difficulty: Medium
Topics: Array, Depth-First Search, Breadth-First Search, Union-Find, Matrix, Counting
Acceptance Rate: 74.6%

## Hints

### Hint 1

The topics list mentions DFS, BFS and Union-Find, but you may not need graph traversal at all. Think about what "communicate" actually requires: a server only needs **one** other server sharing its row or its column.

### Hint 2

For any single server, the question "is there another server in my row or column?" can be answered if you know how many servers are in each row and each column. Can you precompute that information in one pass?

### Hint 3

Count servers per row (`rowCount[i]`) and per column (`colCount[j]`). A server at `(i, j)` communicates if and only if `rowCount[i] > 1` or `colCount[j] > 1`. A second pass over the grid counting the servers that satisfy this condition gives the answer.

## Approach

**Two-pass counting.**

1. **First pass:** Walk the grid. For every cell with `grid[i][j] == 1`, increment `rowCount[i]` and `colCount[j]`.
2. **Second pass:** Walk the grid again. For each server at `(i, j)`, if `rowCount[i] > 1` (another server in its row) or `colCount[j] > 1` (another server in its column), add it to the answer.

Why it works: communication is defined directly (same row or same column), not transitively. So whether a server counts depends only on whether its row or column holds at least one other server — exactly what the counts tell us. Note that the counts include the server itself, which is why the threshold is `> 1` rather than `> 0`.

**Example:** `[[1,1,0,0],[0,0,1,0],[0,0,1,0],[0,0,0,1]]`

- `rowCount = [2, 1, 1, 1]`, `colCount = [1, 1, 2, 1]`
- `(0,0)`: row 0 has 2 → counts. `(0,1)`: row 0 has 2 → counts.
- `(1,2)`: column 2 has 2 → counts. `(2,2)`: column 2 has 2 → counts.
- `(3,3)`: row 3 has 1, column 3 has 1 → isolated.
- Answer: 4.

A graph approach (Union-Find or DFS joining servers in the same row/column, then summing sizes of components with more than one server) also works, but it's more code for the same result. The counting approach is the one to reach for in an interview.

## Complexity Analysis

Time Complexity: O(m · n) — two passes over the grid.
Space Complexity: O(m + n) — the row and column count arrays.

## Edge Cases

- **No servers at all:** every count is zero, so the answer is 0.
- **Single server in the grid:** it has nobody to talk to; the `> 1` threshold correctly excludes it.
- **Diagonal placement:** each server is alone in its row and column (Example 1), answer 0.
- **All cells are servers:** with at least 2 cells, every server communicates; with a 1×1 grid, the answer is 0.
- **Single row or single column grid:** reduces to "if there are 2+ servers, all count; otherwise 0".
- **Server counted via both row and column:** use `||` so it's counted only once, not twice.
