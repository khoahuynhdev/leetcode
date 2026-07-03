# 3620. Network Recovery Pathways

[LeetCode Link](https://leetcode.com/problems/network-recovery-pathways/)

Difficulty: Hard
Topics: Array, Binary Search, Dynamic Programming, Graph Theory, Topological Sort, Heap (Priority Queue), Shortest Path
Acceptance Rate: 36.6%

This is a genuinely hard problem, but it becomes very approachable once you spot that it combines two classic ideas you already know. Take your time — the payoff is a clean, fast solution.

## Hints

### Hint 1

You are asked for the **maximum** of a **minimum** ("the largest minimum-edge cost"). Whenever you see "maximize the minimum" (or "minimize the maximum"), that is a strong signal to think about **binary searching on the answer** rather than trying to compute it directly. Ask yourself: *if I fix a target score `x`, can I check whether some valid path achieves a score of at least `x`?*

### Hint 2

Suppose you fix a candidate score `x`. A path has score at least `x` if and only if **every edge on it costs at least `x`**. So throw away every edge with cost `< x`, and now you only have to answer a simpler yes/no question: using only the remaining edges (and only online intermediate nodes), does there exist a path from `0` to `n − 1` whose total cost is at most `k`? Notice the feasibility is **monotonic**: if a score `x` is achievable, any smaller score is too. That monotonicity is exactly what binary search needs.

### Hint 3

The yes/no check is "is there a `0 → n−1` path with total cost `≤ k`?". The cheapest such path is what matters, so compute the **minimum total cost** to reach `n − 1` using only edges with cost `≥ x`. Because the graph is a **DAG**, you don't need Dijkstra — a single pass in **topological order** gives you the minimum cost with a simple DP relaxation. Only relax *out of* a node that is online (endpoints `0` and `n − 1` are always online), so offline nodes are automatically excluded as intermediates. Feasible when `dist[n−1] ≤ k`. The candidate values of `x` only ever need to be actual edge costs, so binary search over the sorted distinct costs.

## Approach

We want the largest value `x` such that there is a valid path from `0` to `n − 1` on which every edge costs at least `x`. Equivalently, the answer is the best achievable "bottleneck" (minimum edge) subject to the total-cost budget `k`.

**Binary search on the score.** Let `feasible(x)` be true when there exists a path from `0` to `n − 1` such that:

1. every edge used has cost `≥ x`,
2. all intermediate nodes are online,
3. the total cost of the path is `≤ k`.

If `feasible(x)` is true then `feasible(x')` is true for every `x' ≤ x` (fewer restrictions on which edges we may use), so `feasible` is monotonically decreasing in `x`. We binary search for the largest `x` for which it holds. The score of any valid path equals the minimum edge cost along it, which is one of the edge costs, so it suffices to search over the **sorted distinct edge costs**.

**Evaluating `feasible(x)` on a DAG.** Given `x`, keep only edges with cost `≥ x`. We now need the minimum total cost of a path `0 → n−1`. Since the graph is a DAG, process nodes in a fixed **topological order** (computed once with Kahn's algorithm). Maintain `dist[u]` = minimum cost to reach `u` from `0`; initialize `dist[0] = 0` and everything else to infinity. When we visit `u` in topological order:

- Skip it if `dist[u]` is still infinity (unreachable) or if `u` is offline — we may never route *through* an offline node. (`0` and `n − 1` are always online.)
- Otherwise, for each kept edge `u → v` with cost `c ≥ x`, relax `dist[v] = min(dist[v], dist[u] + c)`.

Because we visit nodes in topological order, `dist[u]` is final by the time we process `u`, so one pass suffices. After the pass, `feasible(x)` is `dist[n−1] ≤ k`.

**Putting it together.** Sort the distinct costs, binary search for the largest feasible one. If no cost is feasible (or there are no edges), return `-1`.

*Walkthrough (Example 1):* `edges = [[0,1,5],[1,3,10],[0,2,3],[2,3,4]]`, all online, `k = 10`. Distinct costs `[3,4,5,10]`. For `x = 3` we keep all edges; the cheapest path `0→2→3` costs `7 ≤ 10`, so feasible. For `x = 4` we drop the cost-3 edge, node `2` becomes unreachable, and the only path `0→1→3` costs `15 > 10`, infeasible. The largest feasible candidate is `3`, the answer.

**Watch the magnitudes.** With up to `5·10⁴` edges each costing up to `10⁹`, a path total can reach `5·10¹³`, and `k` itself can be `5·10¹³`. These exceed 32-bit range, so use 64-bit integers throughout (`int64` in Go) for `k`, the distances, and the return value.

## Complexity Analysis

Let `n` be the number of nodes and `m` the number of edges.

Time Complexity: O((n + m) · log m). The topological sort is O(n + m) once. Each `feasible` check is a single O(n + m) DAG pass, and binary search performs O(log m) of them (there are at most `m` distinct costs).

Space Complexity: O(n + m) for the adjacency list, the topological order, the distance array, and the sorted distinct costs.

## Edge Cases

- **No edges at all (`m = 0`).** There is no path from `0` to `n − 1`, so the answer is `-1`. The binary search naturally returns `-1` when the candidate list is empty.
- **No path within budget.** Even the cheapest `0 → n−1` path may exceed `k`, or `0` and `n − 1` may be disconnected. Every candidate is infeasible and the answer is `-1`.
- **Offline intermediate nodes.** A node being offline must forbid routing through it while still allowing it to be reached-but-unused; guarding relaxation on `online[u]` handles this. Endpoints `0` and `n − 1` are guaranteed online.
- **Zero-cost edges and `k = 0`.** Costs can be `0`, so a valid path (and score `0`) can exist even when `k = 0`. Use `≤` comparisons so a total cost exactly equal to `k` counts as valid.
- **Large totals / overflow.** Path totals and `k` can be up to `5·10¹³`; use 64-bit integers to avoid overflow.
- **Direct edge `0 → n−1`.** A single edge is a valid path; its score is its own cost when the cost is within `k`.
```
