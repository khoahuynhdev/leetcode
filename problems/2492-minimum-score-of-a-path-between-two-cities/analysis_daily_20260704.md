# 2492. Minimum Score of a Path Between Two Cities

[LeetCode Link](https://leetcode.com/problems/minimum-score-of-a-path-between-two-cities/)

Difficulty: Medium
Topics: Depth-First Search, Breadth-First Search, Union-Find, Graph Theory
Acceptance Rate: 60.2%

## Hints

### Hint 1

The problem lets a path revisit cities and reuse roads as many times as you like. That freedom is a big clue: think carefully about what "path" really constrains here, and which category of graph problem this belongs to (connectivity, shortest path, minimum spanning tree, ...). It is not as complicated as the "path" wording first suggests.

### Hint 2

Because you can walk back and forth freely, any road that is reachable from city `1` within the same connected chunk of the graph can be included in your path. So the score is just the minimum-distance road among *all* roads you can possibly touch on a walk from `1` to `n`. Now ask: which roads are those?

### Hint 3

Since a path can wander anywhere within a connected component, the set of usable roads is exactly the set of roads belonging to the connected component that contains both city `1` and city `n` (the problem guarantees they are connected). The answer is the smallest `distance` among every road inside that component. You do not need any real "pathfinding" — you only need to explore the component reachable from city `1` and track the minimum edge weight seen.

## Approach

The key realization is that the word "path" is misleading. A path may repeat roads and revisit cities without limit, so once you can reach any road, you can incorporate its distance into your path's minimum. Therefore the score you can achieve is the minimum distance over the entire set of roads that lie in the connected component containing city `1`. The guarantee that `1` and `n` are connected means `n` is in that same component, so this minimum is valid for a `1`-to-`n` path.

This reduces to a simple graph traversal:

1. Build an adjacency list. For each road `[a, b, d]`, add `(b, d)` to `a`'s list and `(a, d)` to `b`'s list.
2. Starting from city `1`, run a BFS (or DFS) over the component. Whenever you traverse an edge, update a running `answer = min(answer, d)`.
3. Mark cities as visited so each is processed once; you still relax every incident edge's weight into the minimum.
4. Return `answer`.

Why it works: every road in the component is incident to some city reachable from `1`, so BFS will visit that city and see the road's weight. No road outside the component can ever be part of a walk from `1`, so ignoring those is correct.

Walkthrough of Example 1 (`n = 4`, roads `[[1,2,9],[2,3,6],[2,4,5],[1,4,7]]`): BFS from `1` reaches `2`, `3`, `4` — the whole graph is one component. The edge weights encountered are `9, 7, 6, 5`, and the minimum is `5`, which matches the expected output. Even though the shortest topological path `1 -> 2 -> 4` gives `min(9, 5) = 5`, we did not need to reason about specific paths — the component minimum gives the same answer directly.

A Union-Find alternative works too: union all roads, then take the minimum distance among roads whose endpoints share the root of city `1`. BFS/DFS is the most direct here.

## Complexity Analysis

Time Complexity: O(n + m), where `n` is the number of cities and `m = roads.length`. We build the adjacency list in O(m) and visit each vertex and edge a constant number of times during traversal.
Space Complexity: O(n + m) for the adjacency list, the visited array, and the BFS queue.

## Edge Cases

- **Disconnected graph with unrelated cheap roads:** A very cheap road in a component that does *not* contain city `1` must be ignored. Restricting the traversal to city `1`'s component handles this — a global minimum over all roads would be wrong.
- **Roads reusable / cities revisited:** The problem's allowance to repeat roads is exactly why component-minimum works; do not try to compute an actual simple path.
- **Minimum comes from an edge not on any direct `1 -> n` path:** As in Example 2, the answer edge (`1-2` with weight `2`) is a detour, not on a straightforward path to `n`. Tracking the component minimum captures it naturally.
- **Multiple edges between the same reachable region:** All incident edges are relaxed into the minimum, so the smallest is always found.
- **1-indexed cities:** Cities are numbered `1..n`; size arrays accordingly (or offset) to avoid off-by-one errors.
