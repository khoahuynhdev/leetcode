# 1254. Number of Closed Islands

[LeetCode Link](https://leetcode.com/problems/number-of-closed-islands/)

Difficulty: Medium
Topics: Array, Depth-First Search, Breadth-First Search, Union-Find, Matrix
Acceptance Rate: 68.0%

## Hints

### Hint 1

This is a grid connectivity problem. Think about how you would count *all* islands first (the classic "Number of Islands" problem), then ask: what makes some islands different from others?

### Hint 2

An island is **not** closed exactly when some cell of it touches the border of the grid, because past the border there is no water to surround it. Use a flood fill (DFS or BFS) to explore each island and track whether it reaches the edge.

### Hint 3

Get rid of the "open" islands first: flood-fill every land cell on the border, turning those whole islands into water. After that, every island still in the grid must be closed, so you can just count them with a normal flood fill.

## Approach

**Two-pass flood fill:**

1. **Remove border-connected land.** Go over the four edges of the grid. Whenever you find a `0` on the edge, run a DFS from it and mark every reachable land cell as water (`1`). These islands touch the boundary, so none of them can be closed.
2. **Count the remaining islands.** Scan the inner cells. Each time you find a `0`, it starts a new island that has no path to the border, which means it is fully surrounded by water. Add one to the count and flood-fill it to water so you don't count it again.

**Why it works:** An island is closed if and only if none of its cells sit on the border. If a cell were on the border, the side facing out of the grid would have no water next to it. The first pass removes exactly the islands that have a border cell. Everything left is closed.

**Alternative (single pass):** Run a DFS from every unvisited `0` and return a boolean meaning "this island never reached the boundary". Be careful to keep exploring the *whole* island even after you find it touches the edge. If you short-circuit with `&&`, you leave cells unvisited, and they get miscounted later.

**Example 2:** `[[0,0,1,0,0],[0,1,0,1,0],[0,1,1,1,0]]`. The left and right columns of zeros touch the border, so pass 1 removes them. The only land left is the `0` at `(1,2)`. Its neighbors are all `1`, so pass 2 counts it, and the answer is `1`.

## Complexity Analysis

Time Complexity: O(m · n). Each cell is visited a constant number of times.
Space Complexity: O(m · n) in the worst case, for the DFS recursion stack (for example, a grid that is all land in a snake shape). The grid is changed in place, so no visited array is needed.

## Edge Cases

- **Single row or single column:** every cell is on the border, so the answer is always `0`.
- **All water:** there are no islands, so the answer is `0`.
- **All land:** there is one island and it touches the border, so the answer is `0`.
- **Island that touches the border at a single cell:** the whole island is still open. Your flood fill must remove or reject the entire island, not just that one cell.
- **Nested islands (Example 3):** a closed island inside a ring of water, which itself sits inside a closed ring of land. Both count, and 4-directional connectivity keeps them separate.
- **Diagonal adjacency:** land that only touches diagonally is not connected. Use only the 4 directions.
