# 3558. Number of Ways to Assign Edge Weights I

[LeetCode Link](https://leetcode.com/problems/number-of-ways-to-assign-edge-weights-i/)

Difficulty: Medium
Topics: Math, Tree, Depth-First Search
Acceptance Rate: 59.4%

## Hints

### Hint 1

The problem mentions a tree and "maximum depth," but read the question carefully: you only ever care about the *path from the root to a deepest node*. That path is just a chain of edges. So the first job is simply to find how many edges are on that longest root-to-leaf path. Think about how you normally find the depth of a tree.

### Hint 2

Now focus on the counting. Each edge gets weight `1` or `2`. The cost of the path is the sum of those weights, and you want the sum to be **odd**. Notice that `2` is even and `1` is odd. So adding a weight-`2` edge never changes the parity of the running sum — only weight-`1` edges flip it. The parity of the total depends entirely on *how many edges you set to `1`*.

### Hint 3

So for a path with `k` edges, an assignment makes the cost odd exactly when an **odd number** of those edges are set to `1`. The number of subsets of a `k`-element set that have odd size is always `2^(k-1)` (it's exactly half of all `2^k` subsets, as long as `k >= 1`). Therefore the answer is `2^(k-1) mod (1e9 + 7)`, where `k` is the maximum depth (in edges). No DFS over weights is needed — just find `k` and do one fast exponentiation.

## Approach

**Step 1 — Find the maximum depth (in edges).**
Build an adjacency list from `edges`. The tree is rooted at node `1`. Run a BFS (or DFS) from node `1`, tracking the distance (number of edges) from the root. The largest distance reached is `k`, the length of the longest root-to-leaf path. BFS is convenient here because each level of the queue corresponds to one more edge of depth, so the number of levels traversed minus one is `k`.

**Step 2 — Count the odd-cost assignments.**
Along a path of `k` edges, every edge is independently `1` or `2`. The total cost is odd iff an odd number of edges carry weight `1`. The count of length-`k` binary choices with an odd number of "ones" is:

```
C(k,1) + C(k,3) + C(k,5) + ... = 2^(k-1)
```

This identity holds for every `k >= 1` (half of all `2^k` combinations have odd parity). Since the constraints guarantee `n >= 2`, there is always at least one edge, so `k >= 1` and the formula is well defined.

Compute `2^(k-1) mod (1e9 + 7)` using fast (binary) exponentiation to keep it efficient and avoid overflow.

**Why it works — a small example.**
Take `edges = [[1,2],[1,3],[3,4],[3,5]]`. BFS from `1`: depth 0 = {1}, depth 1 = {2,3}, depth 2 = {4,5}. The maximum depth is `k = 2`. The valid assignments on a 2-edge path are `(1,2)` and `(2,1)` — both have exactly one weight-`1` edge (odd count). That's `2^(2-1) = 2`, matching the expected output.

## Complexity Analysis

Time Complexity: O(n) — building the adjacency list and the BFS each visit every node and edge once; the fast exponentiation adds O(log k).
Space Complexity: O(n) — for the adjacency list and the BFS queue / visited bookkeeping.

## Edge Cases

- **Single edge (`n = 2`, `edges = [[1,2]]`):** Maximum depth `k = 1`, so the answer is `2^0 = 1`. Exactly one of the two weight choices (`1`) is odd. This is the smallest legal input given `n >= 2`.
- **All leaves at the same depth vs. an unbalanced tree:** Only the *maximum* depth matters; siblings or shorter branches are irrelevant. The BFS must track the deepest level, not the first leaf encountered.
- **Large depth (long chain):** With up to `10^5` nodes the path could have ~`10^5` edges, so `2^(k-1)` is astronomically large — you must reduce modulo `1e9 + 7` during the exponentiation, never compute the raw power.
- **Wide, shallow tree (star graph):** Many children of the root but depth `1`; the answer is still `1`. Confirms the count depends on depth, not node count.
