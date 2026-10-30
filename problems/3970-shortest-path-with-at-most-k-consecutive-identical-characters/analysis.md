# 3970. Shortest Path With At Most K Consecutive Identical Characters

[LeetCode Link](https://leetcode.com/problems/shortest-path-with-at-most-k-consecutive-identical-characters/)

Difficulty: Medium
Topics: String, Graph Theory, Heap (Priority Queue), Shortest Path
Acceptance Rate: 53.5%

## Hints

### Hint 1

The weights are positive and you want the minimum total weight from a source to a target, so think of the classic shortest-path algorithm. The twist is the label rule: whether a path is valid depends on how you got to a node, not only on which node you are at.

### Hint 2

Plain Dijkstra keeps one distance per node, and that is not enough here. Two paths can reach the same node with different "history", and the cheaper one is not always better. Ask yourself what you need to remember about the path so far to decide whether the next step is allowed.

### Hint 3

The only history that matters is **how long the current streak of identical labels is** at the node you are on. Make each state a pair `(node, run)` where `1 <= run <= k`, then run Dijkstra on those states. Following an edge `u -> v` sets `run` to `run + 1` when `labels[u] == labels[v]` (you can't take the edge if that goes over `k`), and resets `run` to `1` otherwise.

## Approach

1. **Build an adjacency list** from the directed edges.
2. **State expansion**: define `dist[node][run]` as the minimum weight to reach `node` with the current identical-label streak ending at `node` exactly `run` long. Only `run` values from 1 to `k` are valid, so there are `n * k` states, at most 2.5 million.
3. **Start**: the path starts at node 0, and its label gives a streak of 1, so `dist[0][1] = 0`. Since `k >= 1`, this state is always valid.
4. **Dijkstra**: pop the state with the smallest distance from a min-heap. Skip it if it's stale. For each outgoing edge `(v, w)`:
   - if `labels[v] == labels[node]`, then `newRun = run + 1`, and you skip the edge if `newRun > k`;
   - otherwise `newRun = 1`.
   - Relax `dist[v][newRun]` with `dist[node][run] + w`.
5. **Early exit**: the first time any state of node `n - 1` is popped, its distance is the answer. Dijkstra pops states in non-decreasing order of distance, and every state at `n - 1` is a valid ending. If the heap runs empty first, return `-1`.

**Why one distance per node fails**: in the test `labels = "aabaa"`, `k = 2`, take the edges `0->3 (1)`, `0->2 (10)`, `2->3 (1)`, `3->4 (1)`. Node 3 can be reached for cost 1 with streak `"aa"` (run 2), but then `3->4` would make `"aaa"`. The costlier route `0->2->3` reaches node 3 with run 1 (`"aba"`), and from there `3->4` is allowed, for a total of 12. Keeping `run` in the state lets both versions of node 3 exist side by side.

Also note that a valid path may visit the same node more than once, for example to go through a node with a different label and break a streak. The state graph handles this on its own, because revisiting a node with a different `run` is a different state.

## Complexity Analysis

Time Complexity: O((n + m) · k · log(m · k)), where m is the number of edges. Each of the `n·k` states is settled once and scans its out-edges, giving `m·k` relaxations, and each heap push or pop costs a log factor.
Space Complexity: O(n · k + m · k) for the distance table and, in the worst case, the heap.

## Edge Cases

- **Single node (`n == 1`)**: the start is the target, so the answer is 0 with no edges used.
- **No edges / target unreachable**: return -1.
- **Directed edges**: an edge `[1, 0, w]` does not let you go from 0 to 1.
- **Every path breaks the streak limit**: e.g. `"aaa"` with `k = 2`, so return -1 even though the target is reachable.
- **A cheaper arrival with a longer streak**: it must not discard a costlier arrival with a shorter streak. This is the reason for the `(node, run)` state.
- **Detours that reset the streak**: the best valid path may be longer in edges, or revisit nodes, just to put a different label in between.
- **Large `k`**: if `k` is at least the length of every path, the label rule never applies and this becomes plain Dijkstra.
