# 1391. Check if There is a Valid Path in a Grid

[LeetCode Link](https://leetcode.com/problems/check-if-there-is-a-valid-path-in-a-grid/)

Difficulty: Medium
Topics: Array, Depth-First Search, Breadth-First Search, Union-Find, Matrix
Acceptance Rate: 53.6%

## Hints

### Hint 1

This is fundamentally a graph reachability problem hiding inside a grid. Forget the picture for a moment — ask yourself: "from cell `(r, c)`, which neighboring cells can I legally step into?" Once you can answer that, you can reach the destination with any standard graph traversal (BFS, DFS, or Union-Find).

### Hint 2

The tricky part is that connectivity is **mutual**. If cell A's street points to its right neighbor B, that's only half the contract — B's street must also accept an opening from its left side. So model each street type as a set of two directions it exposes (e.g. street `4` exposes `{right, down}`), and only traverse the edge between A and B when A exposes the direction toward B *and* B exposes the opposite direction back toward A.

### Hint 3

Encode each of the six street types as the pair of directions it connects (`left`, `right`, `up`, `down`). For each direction define its delta (`dr`, `dc`) and its opposite. Then a BFS or DFS from `(0, 0)` becomes mechanical: for the current cell, look up its two exposed directions; for each, compute the neighbor and verify the neighbor's two exposed directions include the opposite of the one you used. Mark visited and continue. The grid is reachable iff `(m-1, n-1)` ever gets enqueued.

## Approach

Treat the grid as a graph. Each cell is a node, and an edge exists between two adjacent cells iff both of their street types expose the connecting direction.

1. **Direction encoding.** Number the four cardinal directions: `0 = left`, `1 = right`, `2 = up`, `3 = down`. Their deltas are `dr = [0, 0, -1, 1]` and `dc = [-1, 1, 0, 0]`. The opposite-direction lookup is `opp = [1, 0, 3, 2]`.
2. **Street → directions table.** From the problem definition:
   - `1 → {left, right}`
   - `2 → {up, down}`
   - `3 → {left, down}`
   - `4 → {right, down}`
   - `5 → {left, up}`
   - `6 → {right, up}`
3. **Traversal.** Run BFS from `(0, 0)`. Pop a cell `(r, c)`; if it is `(m-1, n-1)` return `true`. Otherwise, for each of its two exposed directions `d`, compute neighbor `(nr, nc) = (r + dr[d], c + dc[d])`. Skip if out of bounds or already visited. Then look up the neighbor's exposed directions and require that they contain `opp[d]` — otherwise the streets do not actually connect even though they sit next to each other. If they do, mark visited and enqueue. If BFS drains without reaching the destination, return `false`.

**Worked example (Example 1, `[[2,4,3],[6,5,2]]`).** Start at `(0,0) = 2` (exposes up, down). Down leads to `(1,0) = 6` (exposes right, up); `6` accepts `up`, so enqueue. From `(1,0)` go right to `(1,1) = 5` (exposes left, up); `5` accepts `left`, enqueue. From `(1,1)` go up to `(0,1) = 4` (exposes right, down); `4` accepts `down`, enqueue. From `(0,1)` go right to `(0,2) = 3` (exposes left, down); `3` accepts `left`. Destination reached → `true`.

A Union-Find variant with cell subdivision exists and is elegant, but for the given constraints (`m, n ≤ 300`, so at most 90,000 cells) the direct BFS is simpler, has the same asymptotic complexity, and is easier to get right under interview pressure.

## Complexity Analysis

Time Complexity: O(m * n) — each of the `m * n` cells is enqueued at most once and processes a constant number (≤ 2) of neighbors.
Space Complexity: O(m * n) — the visited matrix and the BFS queue can each hold up to `m * n` entries in the worst case.

## Edge Cases

- **1×1 grid (`m == n == 1`).** Start equals destination, so the answer is always `true` regardless of the street value. The BFS handles this on its very first pop.
- **Start cell has no compatible neighbor.** E.g. `[[1, 2, 1], [1, 2, 1]]` — `(0,0) = 1` exposes left/right but `(0,1) = 2` does not accept a connection from its left. BFS exits without progress → `false`.
- **Path exists but is not the obvious greedy direction.** A BFS/DFS that explores both exposed directions covers this; do not be tempted to greedily prefer "right then down."
- **Disconnected dead-ends mid-grid.** E.g. `[[1, 1, 2]]` reaches `(0,1)` but `(0,2) = 2` only accepts up/down, so the chain breaks. The mutual-connectivity check is what catches this — without it you would falsely accept the step.
- **Maximum-size grid (300 × 300).** Worst case is 90,000 nodes; iterative BFS stays well within limits, while a recursive DFS would risk stack overflow on a serpentine path. Prefer the iterative form.
