# 3286. Find a Safe Walk Through a Grid

[LeetCode Link](https://leetcode.com/problems/find-a-safe-walk-through-a-grid/)

Difficulty: Medium
Topics: Array, Breadth-First Search, Graph Theory, Heap (Priority Queue), Matrix, Shortest Path
Acceptance Rate: 39.5%

## Hints

### Hint 1

You are moving on a grid and every step either costs you something or it doesn't. The question "can I reach the end with at least 1 health left?" is really the question "what is the *cheapest* way to reach the end?" That reframing should point you toward a **shortest-path** mindset rather than plain reachability. Think of the grid as a graph where each cell is a node and each move is an edge.

### Hint 2

What is the weight of an edge? Walking *into* a cell with `grid[i][j] = 1` costs 1 health; walking into a safe cell (`0`) costs nothing. So every edge weight is either **0 or 1**. Shortest path with only 0/1 edge weights has a specialized, linear-time algorithm — you don't need a full priority-queue Dijkstra (though that works too). Which classic technique handles 0/1 weights efficiently?

### Hint 3

Use **0-1 BFS** with a double-ended queue (deque). Maintain `dist[i][j]` = the minimum number of unsafe cells encountered on any path reaching `(i, j)` (count the starting cell's cost too). When you relax a neighbor: if the move costs 0, push it to the **front** of the deque; if it costs 1, push it to the **back**. This keeps the deque in non-decreasing order of distance, so the first time you finalize a cell you have its true minimum cost. At the end, you can survive iff `health - dist[m-1][n-1] >= 1`.

## Approach

We want the path from `(0, 0)` to `(m-1, n-1)` that passes through the fewest unsafe cells, because each unsafe cell drains exactly 1 health. If that minimum cost is `c`, then starting with `health` and requiring at least 1 health at the destination means we need `health - c >= 1`.

Model the grid as a weighted graph:

- Each cell `(i, j)` is a node.
- Moving from one cell to an adjacent cell has weight equal to the `grid` value of the cell you enter (0 for safe, 1 for unsafe).
- The starting cell `(0, 0)` also contributes its own cost, so initialize `dist[0][0] = grid[0][0]`.

Because all edge weights are 0 or 1, **0-1 BFS** finds the shortest path in `O(V + E)` time:

1. Initialize a `dist` matrix to "infinity" and set `dist[0][0] = grid[0][0]`.
2. Push the start cell into a deque.
3. Pop from the front. For each of the 4 neighbors, compute the candidate distance `dist[cur] + grid[neighbor]`.
   - If it improves the neighbor's recorded distance, update it.
   - Push the neighbor to the **front** of the deque if the edge cost was 0 (same distance layer), or the **back** if it was 1 (next distance layer).
4. Continue until the deque is empty.
5. Return `health - dist[m-1][n-1] >= 1`.

The deque discipline (front for 0-cost, back for 1-cost) guarantees cells are processed in non-decreasing distance order, exactly like Dijkstra but without a heap.

**Walkthrough of Example 3:** `grid = [[1,1,1],[1,0,1],[1,1,1]]`, `health = 5`. Every border cell is unsafe. The cheapest route from top-left to bottom-right that touches the single safe center cell `(1,1)` still passes through several `1`s. The minimum cost turns out to be 4, so `5 - 4 = 1 >= 1` → `true`. Any route avoiding the center costs 5, leaving `0` health → not allowed.

## Complexity Analysis

Time Complexity: O(m * n) — each cell is relaxed a constant number of times (4 neighbors), and 0-1 BFS visits each node/edge a bounded number of times.
Space Complexity: O(m * n) — for the `dist` matrix and the deque.

## Edge Cases

- **Starting cell is unsafe (`grid[0][0] = 1`):** its cost must be counted immediately; initializing `dist[0][0] = grid[0][0]` handles this.
- **Destination cell is unsafe:** counted naturally as the entry cost into `(m-1, n-1)`.
- **Health exactly equal to the minimum cost:** you'd arrive with 0 health, which is *not* allowed — the check must be `>= 1`, not `>= 0`.
- **Smallest grids (`m*n = 2`):** a 1x2 or 2x1 grid; the algorithm still works since both cells' costs are accounted for.
- **All-safe grid:** minimum cost is 0, so any positive health succeeds.
- **No 1-cost detour needed vs. forced detours:** 0-1 BFS correctly prefers 0-cost moves by pushing them to the front, ensuring the truly cheapest path is found even when a longer (in steps) but safer route exists.
