# 3310. Remove Methods From Project

[LeetCode Link](https://leetcode.com/problems/remove-methods-from-project/)

Difficulty: Medium
Topics: Depth-First Search, Breadth-First Search, Graph Theory
Acceptance Rate: 56.0%

## Hints

### Hint 1

The phrase "method `a` invokes method `b`" is just a directed edge `a -> b` in disguise. Before you write any code, redraw the input as a directed graph and re-read the two questions the problem is really asking: *which nodes are suspicious?* and *is the suspicious set safe to delete?* They are two separate sub-problems, and each one is simpler than the combined statement sounds.

### Hint 2

"Method `k`, along with any method invoked by it, directly or indirectly" is the definition of **reachability from a single source**. One traversal — DFS or BFS, your choice — starting at `k` gives you the entire suspicious set. Watch out for cycles: the graph is not a DAG (Example 3 is a 3-cycle), so you need a `visited` array, not a naive recursion.

### Hint 3

The removal rule — "no method **outside** the group invokes any method **within** it" — is not a second traversal. Once you have the suspicious set marked in a boolean array, it is a single linear scan over the original edge list: if there exists any edge `a -> b` where `a` is *not* suspicious and `b` *is* suspicious, the removal is blocked and you return all `n` methods unchanged. Otherwise you return every method that isn't suspicious. Note that an edge from suspicious to non-suspicious can never exist — reachability is closed — so you only ever check one direction.

## Approach

The solution is two linear passes over the graph.

**Step 1 — build the adjacency list.**
Convert `invocations` into `adj[a] = [b, ...]`. With `n` up to `10^5` and up to `2 * 10^5` edges, a slice of slices is the right structure; a map would work but costs constant factors you don't need.

**Step 2 — mark the suspicious set.**
Run an iterative DFS (or BFS) from `k`, marking `suspicious[node] = true` on first visit. The `visited` check is what makes cycles safe: in Example 3 (`2 -> 0 -> 1 -> 2`) the traversal walks the cycle once and stops. An explicit stack is preferred over recursion here — a degenerate chain of `10^5` nodes is a deep call stack, and while Go grows goroutine stacks dynamically, the iterative version has no such concern at all.

After this pass, `suspicious` holds exactly the set the problem wants to delete: `k` plus everything transitively reachable from it.

**Step 3 — validate the removal.**
Scan `invocations` once. The removal is blocked precisely when some edge crosses *into* the suspicious set from outside:

```
for each edge (a, b):
    if !suspicious[a] && suspicious[b] -> blocked
```

Why is checking the raw edge list enough? Because "no method outside the group invokes a method within it" is literally a statement about individual edges — there is no need to reason about paths. And the reverse direction (`suspicious[a] && !suspicious[b]`) is impossible by construction: if `a` is reachable from `k` and `a -> b` is an edge, then `b` is reachable from `k` too, so `b` is suspicious.

**Step 4 — produce the answer.**
If blocked, return `[0, 1, ..., n-1]` — the problem states that if you cannot remove *all* suspicious methods, you remove *none*, so it's all-or-nothing rather than a partial removal. If not blocked, return every index `i` with `suspicious[i] == false`. Any order is accepted; emitting them in ascending index order is free and makes tests deterministic.

**Worked example (Example 1):** `n = 4, k = 1, invocations = [[1,2],[0,1],[3,2]]`.
DFS from `1` visits `1`, then `2`, giving `suspicious = {1, 2}`. Scanning the edges: `[1,2]` is inside-to-inside (fine), `[0,1]` has `0` outside and `1` inside — blocked. Return `[0,1,2,3]`.

**Worked example (Example 2):** `n = 5, k = 0, invocations = [[1,2],[0,2],[0,1],[3,4]]`.
DFS from `0` visits `0, 2, 1`. Scanning: `[1,2]`, `[0,2]`, `[0,1]` are all inside-to-inside; `[3,4]` is entirely outside. Nothing blocks, so we return the non-suspicious nodes: `[3,4]`.

This is honestly on the easier side for a Medium — the difficulty is almost entirely in translating the wordy "group removal" condition into the one-line edge check. Once you see that, the code is a textbook traversal.

## Complexity Analysis

Time Complexity: O(n + m), where `m = len(invocations)`. Building the adjacency list is O(n + m), the traversal visits each node and edge at most once, the validation scan is O(m), and assembling the result is O(n).
Space Complexity: O(n + m) for the adjacency list, the `suspicious` array, and the DFS stack. The output itself is up to O(n).

## Edge Cases

- **`invocations` is empty.** The constraints allow `invocations.length == 0`. The suspicious set is then just `{k}`, nothing invokes it, and the answer is every index except `k`. Make sure your adjacency list is allocated (not left `nil`) so indexing `adj[k]` doesn't panic.
- **`n == 1`.** The only method is `k` itself, it is suspicious, and there are no edges to block it. The answer is the **empty** list. Return an allocated empty slice rather than `nil` if your tests compare with `reflect.DeepEqual`.
- **Cycles in the graph.** Example 3 is a 3-cycle; without a `visited` array the traversal loops forever. Cycles also mean you cannot rely on topological order or DAG-only tricks.
- **The whole graph is suspicious.** When every node is reachable from `k`, no edge can originate outside the set, so removal always succeeds and the answer is empty. This is a nice sanity check that your "blocked" logic doesn't produce false positives.
- **A blocking edge that points deep into the set.** The outside invoker doesn't have to point at `k` — in `0 -> 5` where `5` is three hops from `k`, the removal is still blocked. Scanning *all* edges (not just `k`'s in-edges) is what catches this.
- **Disconnected components.** Nodes unreachable from `k` and not connected to it at all (like `3 -> 4` in Example 2) simply survive; they neither block nor get removed.
- **Multiple blocking edges.** You can stop scanning at the first one, but there's no correctness issue with scanning the whole list — just don't accidentally "un-block" after setting the flag.
