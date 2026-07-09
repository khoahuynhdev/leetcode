# 3532. Path Existence Queries in a Graph I

[LeetCode Link](https://leetcode.com/problems/path-existence-queries-in-a-graph-i/)

Difficulty: Medium
Topics: Array, Hash Table, Binary Search, Union-Find, Graph Theory
Acceptance Rate: 61.9%

## Hints

### Hint 1

Each query asks whether two nodes lie in the same connected component. That is the classic signal to reach for **connectivity** tooling — think Union-Find (DSU) or a component-labeling pass. Before jumping to code, ask yourself what special structure the input has that might make things simpler.

### Hint 2

The key given fact is that `nums` is **sorted in non-decreasing order**. Consider what an edge means in a sorted array: an edge exists between `i` and `j` when `|nums[i] - nums[j]| <= maxDiff`. If two *adjacent* elements are within `maxDiff`, they are connected. What does that tell you about the *shape* of every connected component?

### Hint 3

Because the array is sorted, if `nums[i+1] - nums[i] <= maxDiff` then nodes `i` and `i+1` are directly linked, and connectivity chains along consecutive indices. Crucially, **no edge can ever "jump over" a large gap**: if `nums[k+1] - nums[k] > maxDiff`, then any `i <= k < k+1 <= j` has `nums[j] - nums[i] >= nums[k+1] - nums[k] > maxDiff`, so no edge crosses that boundary. Therefore every connected component is a **contiguous block of indices**. Label the blocks in one pass, then each query is a simple equality check on labels.

## Approach

The naive reading suggests building a full graph and running Union-Find over all pairs — but `n` can be up to 10^5, so we cannot afford to examine all pairs. The sorted-array structure gives us something much cheaper.

**Step 1 — Label contiguous components.**
Walk through the array from left to right, maintaining a component id (starting at 0). For each `i` from `1` to `n-1`:

- If `nums[i] - nums[i-1] <= maxDiff`, node `i` belongs to the same component as node `i-1` (there is a direct edge).
- Otherwise, the gap is too large; start a new component by incrementing the id.

Store this id in a `comp` array, so `comp[i]` is the component label of node `i`.

Why is this correct? In a sorted array, differences between adjacent elements are the smallest possible steps. If every adjacent gap inside a range is `<= maxDiff`, the whole range is connected through a chain of edges. And if some adjacent gap exceeds `maxDiff`, nothing on the left can connect to anything on the right (any cross-boundary difference is at least that gap). So components are exactly the maximal contiguous runs where adjacent gaps stay within `maxDiff`.

**Step 2 — Answer each query in O(1).**
Two nodes `u` and `v` are connected if and only if they carry the same component label: `comp[u] == comp[v]`. A self-query `[u, u]` is trivially `true`, and the equality check already covers it.

**Example walk-through** (`nums = [2,5,6,8]`, `maxDiff = 2`):

- Index 0: comp = 0.
- Index 1: `5 - 2 = 3 > 2` → new component → comp = 1.
- Index 2: `6 - 5 = 1 <= 2` → same component → comp = 1.
- Index 3: `8 - 6 = 2 <= 2` → same component → comp = 1.

So `comp = [0, 1, 1, 1]`. Queries `[0,1] -> 0==1? false`, `[0,2] -> false`, `[1,3] -> 1==1? true`, `[2,3] -> true`, giving `[false, false, true, true]`.

## Complexity Analysis

Time Complexity: O(n + q), where `n` is the number of nodes and `q` is the number of queries — one linear pass to label components and one constant-time lookup per query.
Space Complexity: O(n) for the component-label array (excluding the O(q) output).

## Edge Cases

- **Self-queries `[u, u]`**: Always `true`. The `comp[u] == comp[u]` comparison handles this automatically.
- **Single node (`n == 1`)**: There are no adjacent pairs to check; the lone node is its own component. Any query must be `[0, 0]`, which returns `true`.
- **`maxDiff == 0`**: Only equal adjacent values are connected. Nodes with distinct neighboring values become separate components. (Note duplicates in `nums` still connect since their difference is 0.)
- **All nodes connected**: When every adjacent gap is within `maxDiff`, there is a single component and every query is `true`.
- **No edges at all**: When every adjacent gap exceeds `maxDiff`, each node is isolated; only self-queries return `true`.
- **Large inputs**: With `n` and `q` up to 10^5, an O(n·q) or all-pairs approach would time out — the contiguous-component trick keeps it linear.
