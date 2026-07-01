# 2812. Find the Safest Path in a Grid

[LeetCode Link](https://leetcode.com/problems/find-the-safest-path-in-a-grid/)

Difficulty: Medium
Topics: Array, Binary Search, Breadth-First Search, Union-Find, Heap (Priority Queue), Matrix
Acceptance Rate: 50.6%

## Hints

### Hint 1

This is really two problems stitched together. First, notice that the "safeness factor" of a path depends only on how close each cell gets to the *nearest* thief. So before you worry about paths at all, ask: for every cell in the grid, how far is it from the closest thief? Think about which traversal computes shortest distances from *multiple* starting points at once.

### Hint 2

Once you have, for each cell, its Manhattan distance to the nearest thief, you can precompute it with a **multi-source BFS**: seed a queue with *all* thief cells at distance 0 and let the wave expand outward. In a 4-connected grid, the number of BFS steps to reach a cell equals its Manhattan distance to the nearest source — so a single BFS fills in the whole distance grid in O(n²).

### Hint 3

Now the real question: among all paths from `(0,0)` to `(n-1,n-1)`, you want the one whose *minimum* cell-distance is as *large* as possible. This is a classic **maximin / widest-path** problem. Two clean ways to solve it: (a) a Dijkstra-style walk with a **max-heap**, always expanding the cell reachable with the largest bottleneck value; or (b) **binary search** on the answer `k` combined with a BFS/DSU that only walks through cells with distance `≥ k`. Both give the same result.

## Approach

**Step 1 — Distance-to-nearest-thief grid (multi-source BFS).**
Create a `dist` grid initialized to `-1`. Push every thief cell into a queue with `dist = 0`. Run BFS: when you pop a cell, relax its four neighbors, setting their distance to `current + 1` the first time they're visited. Because every source starts at distance 0 simultaneously, the value that lands on each cell is exactly the Manhattan distance to the *closest* thief. This costs O(n²) time.

**Step 2 — Widest path via a max-heap (Dijkstra variant).**
We want to maximize the minimum `dist` value encountered along a path. Define `safe[r][c]` = the best achievable path-safeness when arriving at `(r, c)`. Start at `(0,0)` with `safe[0][0] = dist[0][0]` and push it onto a max-heap keyed by safeness.

Repeatedly pop the cell with the **largest** current safeness (this greedy choice is what makes Dijkstra correct here). For each neighbor, the safeness of extending the path is `min(current safeness, dist[neighbor])`. If that value improves the neighbor's recorded `safe`, update it and push it. The first time we pop `(n-1, n-1)`, its safeness is the answer — greedy pop order guarantees we've found its optimal bottleneck.

**Why it works.** A path's safeness is a *bottleneck* (minimum along the path), and the max-heap always finalizes the cell with the current best bottleneck first, exactly like Dijkstra finalizes the shortest-distance node first. Substituting `min` for `+` and "max-heap" for "min-heap" turns shortest-path into widest-path.

**Worked example.** For `grid = [[0,0,1],[0,0,0],[0,0,0]]` the only thief is at `(0,2)`. The `dist` grid becomes:

```
2 1 0
3 2 1
4 3 2
```

Starting at `(0,0)` with safeness 2, we can route down and along cells whose distance never drops below 2, reaching `(2,2)` with safeness 2 — the answer.

## Complexity Analysis

Time Complexity: O(n² log n) — the BFS is O(n²); the heap walk visits each of the n² cells and performs O(log n²) = O(log n) heap operations per push.
Space Complexity: O(n²) — the `dist` and `safe` grids plus the heap, all bounded by the number of cells.

## Edge Cases

- **Start or destination is a thief** (`grid[0][0] == 1` or `grid[n-1][n-1] == 1`): `dist` at that cell is 0, so every path through it has safeness 0 — the answer is 0. Example 1 is exactly this case.
- **Every cell is a thief**: all distances are 0, answer is 0.
- **`n == 1`** (single cell): the start *is* the destination; the answer is `dist[0][0]`, which is 0 since the lone cell must contain the sole thief (the constraint guarantees at least one thief).
- **Sparse thieves in a large grid**: distances grow large; make sure the heap comparison maximizes safeness and that `safe` values are only overwritten when strictly improved, to avoid redundant work and infinite loops.
- **Don't forget** cells containing thieves are still walkable — you may pass through a thief cell, it just forces the path safeness to 0.
