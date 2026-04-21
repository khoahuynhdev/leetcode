# 1722. Minimize Hamming Distance After Swap Operations

[LeetCode Link](https://leetcode.com/problems/minimize-hamming-distance-after-swap-operations/)

Difficulty: Medium
Topics: Array, Depth-First Search, Union-Find
Acceptance Rate: 54.9%

## Hints

### Hint 1

The swap operation is transitive. If you can swap indices `(a, b)` and `(b, c)`, then by chaining swaps you can rearrange values at `a`, `b`, and `c` in any order you want. Think of `allowedSwaps` as edges in a graph over indices. What structure emerges from that graph, and how does that structure constrain what arrangements of `source` are reachable?

### Hint 2

Once you recognize the graph structure, notice that within each connected component of indices, the multiset of `source` values is invariant — swaps only shuffle those values among the component's positions. So for each component you have a bag of source values and a bag of target values (at the same positions), and you get to choose any bijection between positions and your source values. The question becomes: how many positions can you make match?

### Hint 3

For a single component, maximum matches is just the size of the multiset intersection of its source values and target values: for each value `v`, you can match `min(srcCount[v], tgtCount[v])` positions. The unmatched positions in that component contribute to the Hamming distance. Sum the mismatches over all components. Union-Find is a clean way to build components over up to 10^5 indices and 10^5 edges.

## Approach

Read the `allowedSwaps` as an undirected graph where nodes are indices `0..n-1`. Swaps are transitive, so within a connected component you can permute the `source` values however you like. That means for each component, the answer depends only on the multiset of source values and the multiset of target values at its positions — not on the exact swap wiring.

**Step-by-step:**

1. **Build connected components.** Use Union-Find (Disjoint Set Union) over `n` indices. For each `[a, b]` in `allowedSwaps`, union `a` and `b`. Use path compression for efficiency.

2. **Group values by component root.** Walk every index `i` once. For its root `r`, maintain two frequency maps: one over `source[i]` values, one over `target[i]` values.

3. **Count matches per component.** For each root, for each value `v` appearing in the source-frequency map, add `min(srcCount[v], tgtCount[v])` to a running total of matches. Target values that never appear in source contribute zero, so iterating over source keys is sufficient.

4. **Return `n - matches`.** Every position is either matched (contributes 0 to Hamming distance) or unmatched (contributes 1).

**Walk-through of Example 1:** `source = [1,2,3,4]`, `target = [2,1,4,5]`, `allowedSwaps = [[0,1],[2,3]]`. Components are `{0,1}` and `{2,3}`. Component `{0,1}`: source values `{1,2}`, target values `{2,1}` — intersection size 2, so 2 matches. Component `{2,3}`: source values `{3,4}`, target values `{4,5}` — intersection is `{4}`, size 1, so 1 match. Total matches = 3, answer = `4 - 3 = 1`.

## Complexity Analysis

Let `n = len(source)` and `m = len(allowedSwaps)`.

Time Complexity: O((n + m) · α(n)), where α is the inverse Ackermann function. Union-Find operations are effectively constant. Grouping and counting are O(n) total across all components.
Space Complexity: O(n) for the parent array and the two frequency maps (total entries summed across components is exactly `n`).

## Edge Cases

- **No swaps allowed (`allowedSwaps` is empty):** Each index is its own component; the Hamming distance is just the raw position-by-position mismatch count.
- **`source` and `target` already identical:** Answer is 0 regardless of swaps.
- **All indices in one component (fully connected swaps):** The answer reduces to `n` minus the multiset-intersection size of all source values vs. all target values.
- **Duplicate values in either array:** The multiset (not set) intersection is what matters; using `min(srcCount[v], tgtCount[v])` handles duplicates correctly. A naive set intersection would overcount.
- **Self-loops or repeated edges in `allowedSwaps`:** Harmless under Union-Find — a redundant union is a no-op.
- **`n = 1`:** One element, no swaps possible to help; answer is 0 if `source[0] == target[0]` else 1.
