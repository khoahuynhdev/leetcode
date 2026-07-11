# 2685. Count the Number of Complete Components

[LeetCode Link](https://leetcode.com/problems/count-the-number-of-complete-components/)

Difficulty: Medium
Topics: Depth-First Search, Breadth-First Search, Union-Find, Graph Theory
Acceptance Rate: 78.2%

## Hints

### Hint 1

This problem is about **connected components** in an undirected graph. Whenever you need to group vertices that are reachable from one another, think about the classic tools for the job: a graph traversal (DFS/BFS) or a **Union-Find (Disjoint Set Union)** structure. Start by asking yourself how you would separate the graph into its independent pieces.

### Hint 2

Once you can identify each component, you need a way to decide whether a component is *complete*. A complete graph on `k` vertices has a very specific, well-known number of edges. So for each component, the two quantities you should track are: **how many vertices** it contains and **how many edges** belong entirely inside it.

### Hint 3

The critical insight is the formula for a complete graph: a component with `k` vertices is complete **iff** it contains exactly `k * (k - 1) / 2` edges (every pair of vertices connected exactly once). Because the input has no repeated edges, you can simply count vertices and edges per component and compare against this formula. Equivalently, every vertex in a complete component of size `k` must have degree `k - 1`.

## Approach

We use **Union-Find** to group the vertices into connected components, then verify completeness with a counting argument.

1. **Build the DSU.** Initialize `parent[i] = i` and `size[i] = 1` for every vertex. For each edge `[a, b]`, union the two vertices. Use path compression and union-by-size to keep operations near-constant time. After processing all edges, every component is represented by a single root, and `size[root]` tells us how many vertices are in that component.

2. **Count edges per component.** Iterate over the edges again. Both endpoints of any edge are guaranteed to be in the same component, so take the root of one endpoint and increment an `edgeCount[root]` tally. This gives the exact number of edges inside each component.

3. **Check completeness.** For every vertex that is a root of its component, let `k = size[root]` and `e = edgeCount[root]`. The component is complete exactly when `e == k * (k - 1) / 2`. Count how many roots satisfy this.

Why it works: a connected component with `k` vertices needs at least `k - 1` edges (a spanning tree). It becomes *complete* only when every pair of vertices shares an edge, which is precisely `C(k, 2) = k * (k - 1) / 2` edges. Since the problem guarantees no duplicate edges, an edge count equal to this maximum can only happen when all pairs are present. Isolated vertices form complete components of size 1 (0 edges, and `1 * 0 / 2 = 0`), so they are correctly counted.

**Example walkthrough** (`n = 6, edges = [[0,1],[0,2],[1,2],[3,4],[3,5]]`):
- Component `{0, 1, 2}`: `k = 3`, edges inside = 3. `3 * 2 / 2 = 3` → complete. ✅
- Component `{3, 4, 5}`: `k = 3`, edges inside = 2 (missing 4–5). `3 * 2 / 2 = 3 ≠ 2` → not complete. ❌

Answer: `1`.

## Complexity Analysis

Time Complexity: O(V + E · α(V)), where `V = n` and `E = edges.length`. Each union/find is effectively constant time with path compression and union-by-size (α is the inverse Ackermann function). We make one pass to union edges, one pass to count edges, and one pass over vertices to check roots.

Space Complexity: O(V) for the `parent`, `size`, and `edgeCount` arrays.

## Edge Cases

- **No edges (`edges = []`):** Every vertex is its own component of size 1 with 0 internal edges. Each is trivially complete, so the answer is `n`. The formula `1 * 0 / 2 = 0` handles this naturally.
- **Single vertex (`n = 1`):** One isolated, complete component → answer `1`.
- **Fully connected graph:** All `n` vertices form one component with `n * (n - 1) / 2` edges → answer `1`. Good for confirming the completeness formula.
- **A component that is connected but not complete** (e.g. a path or a tree of ≥ 3 vertices): connected but missing edges, so it must be excluded. This is the case that distinguishes "connected" from "complete."
- **Mix of complete and incomplete components:** Ensures per-component counting is isolated correctly and one incomplete component doesn't affect the others.
