# 1345. Jump Game IV

[LeetCode Link](https://leetcode.com/problems/jump-game-iv/)

Difficulty: Hard
Topics: Array, Hash Table, Breadth-First Search
Acceptance Rate: 47.4%

## Hints

### Hint 1

You're asked for the **minimum number of steps** in a graph where each index can move to a small set of neighbors. That phrasing — "fewest moves on an unweighted graph" — is a strong signal toward a specific traversal pattern. Which traversal explores nodes in order of distance from the start?

### Hint 2

BFS is the right framework, but the naive edge set is dangerous. From every index `i` you can move to `i-1`, `i+1`, **and** to every other index sharing the same value. If many indices share a value (e.g. `[7, 7, 7, ..., 7]`), naively enumerating those edges on every visit explodes to O(n²). Think about how to precompute the same-value groups, and — more importantly — how to make sure you only pay for each group **once** across the whole BFS.

### Hint 3

Precompute a map from value → list of indices that hold that value. The key trick: the **first** time BFS encounters any index with value `v`, it should push *every* index in the value-`v` group onto the next frontier, and then **delete that group from the map**. Why does deleting work? Because any later index with value `v` would offer no new neighbors — they were all reached at this distance or earlier. This collapses the worst case back to O(n).

## Approach

Model the problem as an unweighted graph and run BFS from index `0`. Each index `i` has up to three kinds of outgoing edges:

1. `i + 1` (if in bounds)
2. `i - 1` (if in bounds)
3. Every `j` with `arr[j] == arr[i]` and `j != i`

BFS gives the shortest path because every edge has equal cost (one step).

**The performance pitfall** comes from same-value edges. Consider `arr = [7, 7, 7, ..., 7]` of length 50,000 — every index is connected to every other index. A naive BFS would re-iterate the full group from every popped node, costing O(n²). The fix:

- Build `groups: map[int][]int` once at the start: `groups[v]` is every index where `arr[i] == v`.
- When BFS visits some index `i`, expand the group `groups[arr[i]]` exactly once: enqueue every unvisited member, then **delete the entry from `groups`**. Subsequent visits to indices with the same value find the entry gone and skip the group expansion — they only contribute their `±1` neighbors.

This guarantees each index is pushed at most once and each group is expanded at most once, yielding O(n) total work.

**Algorithm in level-order BFS form:**

1. If `len(arr) <= 1`, return `0` (already at the last index).
2. Build `groups` mapping value → indices.
3. Initialize `visited[0] = true`, `frontier = [0]`, `steps = 0`.
4. While `frontier` is non-empty:
   - Build the next frontier by expanding every node in the current one.
   - If any node equals `n-1`, return the current `steps`.
   - For each `i`: enqueue unvisited `i+1`, `i-1`, and every unvisited `j` in `groups[arr[i]]`, then delete `groups[arr[i]]`.
   - Increment `steps`.

**Walkthrough with `arr = [100, -23, -23, 404, 100, 23, 23, 23, 3, 404]`** (target index `9`):

| Step | Frontier (indices) | Expansion notes |
|---|---|---|
| 0 | `[0]` | Start at index 0 (value 100). |
| 1 | `[1, 4]` | From 0: `+1 → 1`, same-value 100 group adds `4`. Delete group{100}. |
| 2 | `[2, 3, 5]` | From 1: `+1 → 2`, same-value -23 adds nothing new. From 4: `+1 → 5`, `-1 → 3`. |
| 3 | contains `9` | From 3 (value 404): same-value 404 adds `9`. **Return 3.** |

The path `0 → 4 → 3 → 9` matches the expected answer.

## Complexity Analysis

Time Complexity: O(n). Each index is enqueued at most once. Each value-group is expanded at most once across the entire BFS (because we delete the map entry after first expansion). Building the groups map is O(n). Total work is linear in the array length.

Space Complexity: O(n). The `groups` map stores all `n` indices across its buckets, the `visited` slice has `n` entries, and the BFS frontier holds at most `n` indices.

## Edge Cases

- **Single element (`n == 1`)**: Start index is already the last index — return `0` immediately, without entering the BFS loop.
- **Two elements**: One step suffices via the `+1` edge, regardless of whether the values match.
- **All identical values (e.g. `[7, 7, 7, 7, 7]`)**: One same-value jump reaches the last index in a single step. The "delete after expansion" rule is essential here — without it, the BFS would do quadratic work.
- **Strictly distinct values (e.g. `[1, 2, 3, 4, 5]`)**: Same-value edges contribute nothing; the BFS walks `±1` only, giving `n - 1` steps.
- **Duplicates that don't help reach the end early (e.g. `[11, 22, 11, 33, 44]`)**: The same-value bucket for `11` is expanded but doesn't shortcut to the goal; BFS still correctly returns the shortest path through `±1` edges.
- **Large input near the constraint limit (`n = 5 * 10^4`) with many duplicates**: The early-termination check inside the level loop, combined with the group-deletion optimization, keeps runtime well within limits.
