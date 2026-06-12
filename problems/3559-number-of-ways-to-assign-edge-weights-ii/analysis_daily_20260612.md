# 3559. Number of Ways to Assign Edge Weights II

[LeetCode Link](https://leetcode.com/problems/number-of-ways-to-assign-edge-weights-ii/)

Difficulty: Hard
Topics: Array, Math, Dynamic Programming, Bit Manipulation, Tree, Depth-First Search
Acceptance Rate: 67.6%

This problem looks intimidating because it combines trees, counting, and modular
arithmetic — but the counting half collapses to a tiny formula once you spot it.
The real work is answering many path-distance queries efficiently. Take it one
layer at a time and it becomes very approachable.

## Hints

### Hint 1

Forget the tree for a moment and look at a single path with exactly `k` edges.
Each edge is assigned either `1` or `2`. Only the *parity* of the total cost
matters. Notice that weight `1` flips the parity and weight `2` leaves it
unchanged. So really each edge is a binary choice: "flip" or "don't flip". How
many of the `2^k` total assignments end up with an *odd* number of flips?

### Hint 2

The number of ways to choose an odd-sized subset out of `k` items is exactly
`2^(k-1)` whenever `k >= 1` (and `0` when `k = 0`, since an empty path has cost
`0`, which is even). That means every query's answer depends on **one number
only**: `k`, the number of edges on the path between `u` and `v`. The entire
problem reduces to: for each query, find the path length, then return
`2^(k-1) mod (1e9+7)`.

### Hint 3

The number of edges on the path between `u` and `v` in a rooted tree is the
classic distance formula:

```
dist(u, v) = depth[u] + depth[v] - 2 * depth[lca(u, v)]
```

So the whole problem is really "answer up to `10^5` LCA / distance queries on a
tree of up to `10^5` nodes." Precompute depths and a **binary-lifting** table
(`up[k][node]` = the `2^k`-th ancestor) so each LCA is `O(log n)`, then plug the
distance into the `2^(k-1)` formula. Don't forget the `k = 0` case (a node paired
with itself, or two equal endpoints).

## Approach

There are two independent parts.

**1. The counting formula.**
For a fixed path of `k` edges, an assignment makes the cost odd exactly when an
odd number of edges receive weight `1`. The count of odd-sized subsets of a
`k`-element set equals `2^(k-1)` for `k >= 1`. (Half of all `2^k` subsets are
odd-sized; pairing each subset with its complement-on-the-last-element gives a
bijection.) For `k = 0` the path is empty, cost is `0` (even), so the answer is
`0`.

So `answer = (k == 0) ? 0 : 2^(k-1) mod M`, with `M = 1e9 + 7`.

**2. Computing `k = dist(u, v)` fast.**
We build the tree from `edges`, root it at node `1`, and run one DFS (done
iteratively to avoid stack overflow at `n = 10^5`) to record `depth[node]` and
each node's immediate parent. We then build a binary-lifting table where
`up[j][v]` is the ancestor of `v` that is `2^j` steps above it. With it, the LCA
of `u` and `v` is found in `O(log n)`:

1. Lift the deeper node up until both are at the same depth.
2. If they coincide, that node is the LCA.
3. Otherwise lift both in lockstep by the largest jumps that keep them distinct;
   their common parent is the LCA.

Finally `dist(u, v) = depth[u] + depth[v] - 2*depth[lca]`, and the answer is the
precomputed power of two.

**Worked example** (`edges = [[1,2],[1,3],[3,4],[3,5]]`):
depths are `depth[1]=0, depth[2]=1, depth[3]=1, depth[4]=2, depth[5]=2`.
- Query `[1,4]`: `lca=1`, `dist = 0 + 2 - 0 = 2`, answer `2^1 = 2`.
- Query `[3,4]`: `lca=3`, `dist = 1 + 2 - 2 = 1`, answer `2^0 = 1`.
- Query `[2,5]`: `lca=1`, `dist = 1 + 2 - 0 = 3`, answer `2^2 = 4`.

Output `[2, 1, 4]`, matching the expected result.

We precompute all powers of two up to `n` once, so each query is just an
`O(log n)` LCA plus an array lookup.

## Complexity Analysis

Let `n` be the number of nodes and `q` the number of queries.

Time Complexity: `O(n log n + q log n)` — building the binary-lifting table costs
`O(n log n)`, and each of the `q` queries costs `O(log n)` for the LCA.

Space Complexity: `O(n log n)` for the binary-lifting table (plus `O(n)` for the
adjacency list, depths, and the table of powers of two).

## Edge Cases

- **Self query (`u == v`)**: distance is `0`, so the answer must be `0`. The
  `2^(k-1)` formula is undefined at `k = 0`; guard it explicitly.
- **Single-edge tree / adjacent nodes**: `k = 1` gives `2^0 = 1`, the smallest
  non-zero answer. Make sure exponent `k-1 = 0` is handled.
- **Deep / skewed (path-like) trees**: with `n = 10^5` a recursive DFS can blow
  the call stack; use an iterative DFS or BFS to compute depths and parents.
- **Large distances and the modulus**: `k` can be up to `n-1 ≈ 10^5`, so
  precompute `2^i mod (1e9+7)` rather than recomputing per query.
- **Node labels are 1-indexed**: keep arrays sized `n+1` and be consistent, or
  off-by-one bugs creep into depth/parent lookups.
