# 3534. Path Existence Queries in a Graph II

[LeetCode Link](https://leetcode.com/problems/path-existence-queries-in-a-graph-ii/)

Difficulty: Hard
Topics: Array, Two Pointers, Binary Search, Dynamic Programming, Greedy, Bit Manipulation, Graph Theory, Sorting
Acceptance Rate: 44.4%

This is a genuinely tough problem — it combines a clever graph observation with an
efficient way to answer many shortest-path queries. Don't be discouraged if the
optimal solution doesn't appear immediately; the payoff is a very reusable pattern
(sorting + greedy reach + binary lifting). Take it hint by hint.

## Hints

### Hint 1

The edge rule only depends on the *values* `nums[i]`, not on the node labels. An edge
connects two nodes whenever their values differ by at most `maxDiff`. What happens if
you stop thinking about the nodes in their original order and instead **sort them by
value**? Two nodes that are far apart in the label order but close in value are
neighbors — sorting makes that structure visible.

### Hint 2

Once the values are sorted, notice that a node connects to a *contiguous window* of
the sorted array: everything whose value lies within `maxDiff` of it. That immediately
tells you two things. First, connectivity: two sorted-adjacent nodes are directly
connected exactly when their gap is `<= maxDiff`, so the components are maximal runs of
sorted values with no gap larger than `maxDiff`. Second, movement: from any node the
*best* single move toward a target on its right is to jump to the **farthest** node
still within `maxDiff`. Think about how to precompute that farthest reach for every
position (hint: two pointers, because the reach only moves rightward as you advance).

### Hint 3

The shortest path between two nodes in the same component reduces to: "starting at the
left one, how few greedy 'jump as far right as you can' moves reach the right one?"
That is a classic *minimum-jumps* problem. Doing it naively per query is too slow for
`10^5` queries, but the greedy jump is a fixed function `far[i]` — so you can build a
**binary-lifting** table: `up[k][i]` = where you land after `2^k` greedy jumps. Then any
query counts the jumps in `O(log n)` by descending powers of two: take a `2^k` jump
whenever it still leaves you strictly left of the target, then add one final jump.

## Approach

**Step 1 — Sort by value.** Build `order`, the node indices sorted by `nums`. Record
`sortedVals[r]` = the value at sorted rank `r`, and `pos[node]` = the sorted rank of an
original node so we can translate queries into sorted-space positions.

**Step 2 — Components.** Walk the sorted array left to right. Start component `0`; each
time `sortedVals[i] - sortedVals[i-1] > maxDiff`, increment the component id. Two nodes
can reach each other **iff** they share a component. This works because within a run
where every adjacent gap is `<= maxDiff`, each node connects to its sorted neighbor, so
the whole run is connected; and a gap larger than `maxDiff` is an unbridgeable wall (no
node on one side is within `maxDiff` of any node on the other side, since values are
sorted).

**Step 3 — Greedy farthest reach.** For each sorted position `i`, `far[i]` is the
largest index `j` with `sortedVals[j] - sortedVals[i] <= maxDiff`. Because `far` is
non-decreasing in `i`, a two-pointer sweep computes all of them in `O(n)`. From node
`i`, one hop can land anywhere in `[left..far[i]]`; to travel rightward as fast as
possible, always aim for `far[i]`.

**Step 4 — Binary lifting.** Set `up[0] = far`. Then `up[k][i] = up[k-1][up[k-1][i]]`
means "apply `2^(k-1)` greedy jumps twice." Build `~log2(n)` levels.

**Step 5 — Answer queries.** For query `(u, v)`:
- Let `pu = pos[u]`, `pv = pos[v]`. If `pu == pv`, the answer is `0`.
- If `comp[pu] != comp[pv]`, they're disconnected → `-1`.
- Otherwise let `l = min(pu, pv)`, `r = max(pu, pv)`. Starting at `cur = l` with
  `steps = 0`, iterate `k` from high to low: if `up[k][cur] < r`, take that jump
  (`cur = up[k][cur]`, `steps += 2^k`). After the loop, one more greedy jump from `cur`
  reaches `r` (because it landed as far as possible while still strictly left of `r`),
  so the answer is `steps + 1`.

**Worked micro-example** (Example 2): `nums = [5,3,1,9,10]`, `maxDiff = 2`. Sorted
values are `[1,3,5,9,10]`. Gaps `2,2,4,1`, so the `4` splits things into components
`{1,3,5}` and `{9,10}`. `far` in sorted space is `[1,2,2,4,4]`. Query `[0,2]` maps to
sorted positions `2` and `0` → `l=0, r=2`: from `0` jump to `far[0]=1` (`steps=1`),
then one final jump reaches `2` → answer `2`. Query `[2,3]` crosses the component wall
→ `-1`.

## Complexity Analysis

Let `n = len(nums)`, `q = len(queries)`.

Time Complexity: `O(n log n + q log n)` — sorting is `O(n log n)`; the two-pointer reach
and component scan are `O(n)`; the binary-lifting table is `O(n log n)` to build; each
query costs `O(log n)`.

Space Complexity: `O(n log n)` for the binary-lifting table (plus `O(n)` for the sorted
arrays, components, and reach).

## Edge Cases

- **Self query (`u == v`)**: distance is `0`. Handled before any component/jump logic;
  also covers Example 3's `[0,0]`.
- **Disconnected pair**: different components → `-1` (Example 2's `[2,3]`,
  Example 3's `[0,1]` and `[1,2]`).
- **`maxDiff == 0`**: only nodes with *identical* values connect. The component scan
  naturally handles this since any positive gap exceeds `0`.
- **All nodes isolated**: every adjacent gap exceeds `maxDiff`, so each node is its own
  component; every cross-node query returns `-1`.
- **Single node (`n == 1`)**: the only valid query is `[0,0]` → `0`; the binary-lifting
  table must still be sized safely (at least one level) — guard the `LOG` computation.
- **Duplicate values**: they form zero-gap edges and belong to the same component;
  the sort keeps them adjacent so nothing special is needed.
- **Long chains**: e.g. values `1,2,3,4,5` with `maxDiff = 1` require `n-1` hops end to
  end — this is exactly what binary lifting keeps fast.
