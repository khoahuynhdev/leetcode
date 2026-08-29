# 2948. Make Lexicographically Smallest Array by Swapping Elements

[LeetCode Link](https://leetcode.com/problems/make-lexicographically-smallest-array-by-swapping-elements/)

Difficulty: Medium
Topics: Array, Union-Find, Sorting
Acceptance Rate: 62.2%

## Hints

### Hint 1

The operation is a swap, and swaps can be repeated any number of times. That means "can be swapped with" is really a *reachability* question: if `a` can swap with `b`, and `b` can swap with `c`, then `a` and `c` can end up in each other's positions too, even if `|a - c| > limit`. Whenever a problem asks "what can reach what through repeated pairwise operations", think about connected components — union-find or a grouping pass.

### Hint 2

Building the graph explicitly is hopeless: with `n` up to `10^5`, there can be `O(n^2)` swappable pairs. But notice the edge condition only involves *values*, not positions. Sort the values. What do the connected components look like once the values are in sorted order?

### Hint 3

After sorting, two values are directly connected only if they are close in value, so each connected component is a **contiguous block of the sorted array**: scan the sorted values left to right and start a new group whenever the gap to the previous value exceeds `limit`. Within one group, the elements are freely permutable among the positions they originally occupied. So collect that group's original indices, sort them ascending, and pour the group's sorted values into them in order — smallest value into the earliest position. That greedy placement is exactly what lexicographic minimality demands.

## Approach

The whole problem collapses once you see two facts.

**Fact 1: swappability is transitive through chains.** If `|a - b| <= limit` and `|b - c| <= limit`, you can move `a` into `c`'s slot in two steps (swap `a`/`b`, then `b`/`c`), regardless of how far apart `a` and `c` are. So the real structure is the connected components of the graph where values are nodes and edges join values within `limit` of each other. Inside a component, *any* permutation of the members is achievable (a connected swap graph generates the full symmetric group on its vertices).

**Fact 2: components are contiguous in sorted order.** Sort the values ascending. If consecutive sorted values `v[k]` and `v[k+1]` differ by more than `limit`, then no value at or below `v[k]` can be within `limit` of any value at or above `v[k+1]` — the gap only widens. Conversely, if the consecutive gap is `<= limit`, they are directly connected. So the components are precisely the maximal runs of the sorted array with all consecutive gaps `<= limit`. This gives a linear scan instead of any explicit union-find, though a DSU over sorted-adjacent pairs would work identically.

The algorithm:

1. Build an index array `order = [0, 1, ..., n-1]` and sort it by `nums[i]` ascending. Sorting indices rather than values keeps each value tied to the position it came from.
2. Sweep `order` left to right. Start a new group at `k` when `nums[order[k]] - nums[order[k-1]] > limit`; otherwise extend the current group.
3. For each group, take its slice of `order` — these are the original indices of the group's members. The values are already ascending (that's the sort order). Sort the indices ascending too.
4. Write the `t`-th smallest value of the group into the `t`-th smallest index of the group.

Step 4 is the greedy core. Positions outside the group are untouchable by this group's elements, so the only freedom is how to distribute these values over these positions. To be lexicographically smallest, the earliest position we control must get the smallest value we have, then the next earliest gets the next smallest, and so on. Any swap of two assignments would put a larger value at an earlier controlled index, making the array lexicographically larger.

**Walkthrough with Example 2:** `nums = [1,7,6,18,2,1]`, `limit = 3`.

Sorted by value, as (value, index) pairs: `(1,0) (1,5) (2,4) (6,2) (7,1) (18,3)`.

Consecutive gaps: `0, 1, 4, 1, 11`. Break where the gap exceeds `3` — after `(2,4)` and after `(7,1)`.

- Group A: values `[1,1,2]`, indices `{0,5,4}` → sorted indices `0,4,5` → `nums[0]=1, nums[4]=1, nums[5]=2`.
- Group B: values `[6,7]`, indices `{2,1}` → sorted indices `1,2` → `nums[1]=6, nums[2]=7`.
- Group C: value `[18]`, index `3` → `nums[3]=18`.

Result: `[1,6,7,18,1,2]`, matching the expected output. Note the `18` sits still because it is alone in its component, and the two `1`s could not both move to the front — index 0 gets one, and index 4 (the earliest remaining controlled position) gets the other.

A small implementation nicety: sorting the group's indices ascending is a second sort, but the groups partition `order`, so the total extra work is bounded by `O(n log n)` overall — and in practice each group's index slice is short.

## Complexity Analysis

Time Complexity: O(n log n) — one sort of the `n` indices by value, plus per-group sorts of the index slices whose sizes sum to `n` (so `sum of g log g <= n log n`), plus linear scans.

Space Complexity: O(n) — the index permutation, the group slices, and the output array. The recursion-free sweep adds only O(1) beyond that.

## Edge Cases

- **Single element (`n == 1`).** No pair exists, so the answer is the input unchanged. The group sweep naturally emits one group of size one; just make sure the "look at the previous sorted element" comparison doesn't run on an empty prefix.
- **No swaps possible at all** (Example 3, `[1,7,28,19,10]` with `limit = 3`). Every consecutive sorted gap exceeds `limit`, so every group is a singleton and the output is a copy of the input. Good check that you never accidentally sort the whole array.
- **Everything in one group** (e.g. `limit >= max - min`, or a chain like `[3,1,2,8]` with `limit = 1` where `1-2-3` links up). The answer is the fully sorted array for that group's positions. This is where the transitivity insight earns its keep — `|3 - 1| = 2 > 1`, yet `3` and `1` still swap via `2`.
- **Duplicate values.** Duplicates are always within `limit` of each other (gap `0`), so they land in the same group. The sorted-index assignment handles them fine; be careful only if you sort values with an unstable comparator and expect a particular tie order — it doesn't matter here, since equal values are interchangeable.
- **Large values with a large limit.** `nums[i]` and `limit` both reach `10^9`. Differences fit comfortably in `int` on 64-bit Go, but if you ever compute a sum or use 32-bit ints, watch for overflow. Comparing `nums[a] - nums[b] > limit` on positive `int` values is safe here.
- **Already-sorted input.** Output equals input; the algorithm still does the full sort, which is fine, but it's a nice sanity test that the position mapping is an identity when values are ascending.
