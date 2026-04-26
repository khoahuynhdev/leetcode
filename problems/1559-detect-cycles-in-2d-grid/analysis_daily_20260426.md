# 1559. Detect Cycles in 2D Grid

[LeetCode Link](https://leetcode.com/problems/detect-cycles-in-2d-grid/)

Difficulty: Medium
Topics: Array, Depth-First Search, Breadth-First Search, Union-Find, Matrix

Acceptance Rate: 54.8%

## Hints

### Hint 1

Think of the grid as a graph where each cell is a node and there is an edge between two cells only if they are 4-directionally adjacent and share the same character. The question "does a cycle exist?" is then a classic graph question. Which traversal techniques are best at detecting cycles in undirected graphs?

### Hint 2

A standard depth-first search can detect a cycle in an undirected graph if you carefully ignore the edge you just traversed. As you walk through cells of the same value, keep track of which cell you arrived from (the "parent"). If you ever reach a cell that is already visited and is *not* your parent, you have found a cycle. Notice how this matches the problem's "you cannot return to the cell you just visited" constraint.

### Hint 3

The constraint "cycle length ≥ 4" is a free consequence of the parent rule on a 4-directional grid: any cycle in such a grid must enclose at least one cell, so its length is automatically at least 4. This means you do *not* need to track depth or path length — you just need DFS that ignores the parent and returns true the first time it lands on an already-visited same-value cell. Use a single global `visited` matrix; once a connected component is fully explored without a cycle, none of its cells need to be revisited from a different starting point.

## Approach

Treat the grid as an undirected graph. Two cells are connected if they are 4-adjacent and contain the same character. The question becomes: does any connected component contain a cycle?

For each cell that has not yet been visited, run a DFS that walks only through cells with the same character. While exploring, pass along the coordinates of the cell you arrived from (the parent). For each neighbor:

1. Skip it if it is out of bounds or has a different character.
2. Skip it if it is the parent (the immediate previous cell).
3. If it is already marked visited, you have just discovered a back-edge — return `true`, a cycle exists.
4. Otherwise, mark it visited and recurse with the current cell as the new parent.

Why "length ≥ 4" is automatic: in a 4-directional grid graph, the smallest possible undirected cycle is a 2×2 square, which has length 4. There are no triangles. So any back-edge found by the DFS corresponds to a cycle of length at least 4 — no extra bookkeeping needed.

Why a single shared `visited` matrix is enough: once a DFS finishes a component without finding a cycle, every cell in that component has been visited and proven cycle-free. Starting a new DFS from any of those cells would only re-explore the same component. The outer loop only kicks off DFS for genuinely unvisited cells, so each cell is processed exactly once.

Worked example. Take grid 3:

```
a b b
b z b
b b a
```

Starting at `(0,1)='b'`, DFS reaches `(0,2)` then `(1,2)`, but every onward neighbor is either out of bounds, a different character, or the parent. No back-edge. Similarly for the lower-left `b` cluster from `(1,0)`. The `a`s are isolated, `z` is alone. No DFS ever lands on a previously visited cell that is not its parent, so the answer is `false`.

For grid 1, the outer ring of `a`s forms a 12-step loop. DFS will eventually walk around the ring and arrive back at `(0,0)` — already visited, not the current parent — and immediately report a cycle.

## Complexity Analysis

Time Complexity: O(m * n). Each cell is marked visited at most once, and each cell expands to at most 4 neighbors, so the total work is linear in the number of cells.

Space Complexity: O(m * n). The `visited` matrix is m × n. The DFS recursion stack can grow up to O(m * n) in the worst case (e.g., a long snake of identical characters covering the whole grid). For grids up to 500 × 500 this is acceptable in Go because goroutine stacks grow dynamically; if you wanted a hard guarantee, the same logic can be rewritten with an explicit stack.

## Edge Cases

- **1 × 1 grid**: A single cell can never form a cycle. The DFS finds no neighbors of the same value and returns `false`.
- **Single-row or single-column grid (1 × n or m × 1)**: No cycle is possible because the graph is acyclic (a path graph). The parent-skip rule correctly prevents any false positives.
- **All cells share the same character**: Even a 2 × 2 of the same character contains a length-4 cycle, so the algorithm should return `true`. Verifying this small case is a good correctness check.
- **All cells distinct**: Every DFS visits exactly one cell. No back-edge can fire. Returns `false`.
- **Multiple disconnected same-value components, only one with a cycle**: The outer loop must keep scanning every cell so that components are not skipped. The shared `visited` array handles this correctly because each component is entered exactly once.
- **Long snake-like component without a cycle**: This is the worst case for recursion depth. Make sure the DFS does not over-mark or miss the parent check, which would otherwise produce a spurious cycle report.
