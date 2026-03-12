# 3600. Maximize Spanning Tree Stability with Upgrades

[LeetCode Link](https://leetcode.com/problems/maximize-spanning-tree-stability-with-upgrades/)

Difficulty: Hard
Topics: Binary Search, Greedy, Union-Find, Graph Theory, Minimum Spanning Tree
Acceptance Rate: 49.0%

## Hints

### Hint 1

Instead of trying to build the best spanning tree directly, think about what happens if you fix the answer — the minimum edge strength in the tree. If you knew the target stability value, could you check whether it's achievable? What search technique works well when you can frame a problem as "is this target feasible?"

### Hint 2

Binary search on the answer. For a candidate stability value `mid`, every edge in your spanning tree must have effective strength >= `mid`. Must-include edges have fixed strength. Optional edges can either use their original strength or their doubled strength (if you spend an upgrade). Use Union-Find to greedily check whether you can connect all `n` nodes: first add must edges, then free optional edges (strength >= mid), and finally upgraded optional edges (2 * strength >= mid) — minimizing upgrade usage.

### Hint 3

The key insight is a two-pass greedy within each binary search check. After including all mandatory edges (checking for cycles — if any must edges form a cycle, it's always -1), add optional edges that already meet the threshold for free, then only spend upgrades on edges that need doubling. This minimizes upgrade count. If the graph becomes connected with at most `k` upgrades, the threshold `mid` is achievable. Binary search finds the maximum achievable threshold.

## Approach

1. **Early termination**: First check if the mandatory edges (must == 1) form a cycle using Union-Find. If any two must-edges connect already-connected nodes, no valid spanning tree exists — return -1.

2. **Binary search on the stability value**: The answer must be an integer in the range [1, 200000] (since max strength is 10^5 and doubling gives at most 2*10^5). Binary search for the maximum feasible stability.

3. **Feasibility check for a candidate value `mid`**:
   - Create a fresh Union-Find for `n` nodes.
   - Add all must-edges. If any must-edge has strength < `mid`, this threshold is infeasible (we can't upgrade must-edges).
   - Add all optional edges with strength >= `mid` (no upgrade needed — free).
   - Add optional edges where strength < `mid` but 2 * strength >= `mid` (needs one upgrade). Only count an upgrade if the edge actually connects two different components.
   - If all nodes are connected and upgrades used <= `k`, the threshold is feasible.

4. **Result**: The maximum feasible threshold from binary search, or -1 if even threshold = 1 is infeasible.

## Complexity Analysis

Time Complexity: O(E * log(S_max) * α(N)) where E is the number of edges, S_max is 200000 (the maximum possible doubled strength), and α is the inverse Ackermann function from Union-Find. Effectively O(E * log(S_max)).

Space Complexity: O(N) for the Union-Find data structure.

## Edge Cases

- **Must-edges form a cycle**: Any set of mandatory edges that creates a cycle makes it impossible to form a spanning tree. Return -1 immediately.
- **Graph is disconnected**: If there aren't enough edges to connect all components even at the lowest threshold, return -1.
- **No optional edges needed**: If must-edges already form a spanning tree, the answer is simply the minimum strength among must-edges.
- **All edges are optional**: Standard maximum-bottleneck spanning tree problem with the added upgrade mechanic.
- **k = 0**: No upgrades available — pure maximum-bottleneck spanning tree with must-edge constraints.
- **Single must-edge with low strength**: This caps the stability regardless of how many upgrades you have for other edges.
