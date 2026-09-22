# 3525. Find X Value of Array II

[LeetCode Link](https://leetcode.com/problems/find-x-value-of-array-ii/)

Difficulty: Hard
Topics: Array, Math, Segment Tree
Acceptance Rate: 55.2%

## Hints

### Hint 1

Start by translating the wordy operation into something concrete. Removing the
prefix `nums[0..start-1]` and then removing a suffix leaves exactly
`nums[start..j]` for some `j` in `[start, n-1]`. So each query is really asking:
**how many `j >= start` have `nums[start] * ... * nums[j] % k == x`?**

Point updates plus range questions is a strong signal for a segment tree. The
real work is figuring out what a node must store so that two nodes can be
combined.

### Hint 2

Look hard at the constraint `1 <= k <= 5`. That is tiny, and tiny constraints
are never accidental. Every element only matters through `nums[i] % k`, so any
range only has `k` possible "product" values and each of its prefixes lands in
one of `k` buckets.

That suggests a node summary of size `O(k)`: a small histogram of residues, not
a list of prefixes.

### Hint 3

Store two things per node covering a range:

- `prod` — the product of the entire range modulo `k`
- `base[p]` — how many prefixes of the range have product `p` modulo `k`

The merge is the insight. Prefixes of the left child are already complete
prefixes of the parent, so they carry over unchanged. Every prefix of the right
child, however, is preceded by the *whole* left child, so its residue `p` shifts
to `(L.prod * p) % k`:

```
parent.prod    = L.prod * R.prod % k
parent.base[v] = L.base[v] + sum of R.base[p] for all p with (L.prod * p) % k == v
```

Notice this merge is **not commutative** — order matters, which is fine for a
segment tree but rules out anything like a Fenwick tree over an unordered
aggregate.

## Approach

**Step 1 — restate the query.** After the persistent update `nums[index] = value`,
the answer for `(start, x)` is the number of indices `j` in `[start, n-1]` with

```
nums[start] * nums[start+1] * ... * nums[j] ≡ x (mod k)
```

This is a count over the prefixes of the suffix `nums[start..]`.

**Step 2 — design the node.** Because `k <= 5`, a node covering a range is fully
described by `prod` (the range product mod `k`) and the histogram `base[0..k-1]`
counting its prefixes by residue. That is `k + 1` small integers per node.

**Step 3 — merge.** As in Hint 3, the left child's prefixes are untouched and the
right child's prefixes all get multiplied by `L.prod`. Each merge costs `O(k)`,
so building the tree is `O(n·k)` and a point update is `O(k·log n)`.

**Step 4 — answer a query.** The suffix `[start, n-1]` decomposes into `O(log n)`
nodes. Visit them **left to right**, carrying a running residue `r` that starts
at `1 % k` (the empty product; `1 % k` rather than `1` so `k == 1` behaves). At
each node:

- add `sum of node.base[p]` over the `p` with `(r * p) % k == x` — these are the
  prefixes that end inside this node
- then advance `r = r * node.prod % k` before moving to the next node

The recursion in `solution.go` threads `r` through as a return value, which keeps
the left-to-right ordering automatic.

**Worked example.** `nums = [1,2,3,4,5]`, `k = 3`, query `[2,2,0,2]`. The update
makes `nums = [1,2,2,4,5]`, and `start = 0` so we look at every prefix. The
prefix products mod 3 are `1, 2, 1, 1, 2`. Two of them equal `x = 2`, matching
the expected answer of 2. The segment tree never materializes that list; it
reaches the same count from a handful of `O(k)` histograms.

## Complexity Analysis

Time Complexity: O(n·k + q·k·log n) — building the tree touches each of the `O(n)`
nodes with an `O(k)` merge; each query does one point update and one suffix walk,
both `O(k·log n)`. With `k <= 5` this is effectively `O(n + q·log n)`.

Space Complexity: O(n·k) for the `prod` and `base` arrays (`int32` counts keep
the constant small; a count never exceeds `n <= 10^5`).

## Edge Cases

- **`k == 1`.** Every product is `0 mod 1` and `x` must be `0`, so the answer is
  always `n - start`. The code handles this only because the identity residue is
  written as `1 % k` (which is `0` here) instead of a hard-coded `1`.
- **Elements divisible by `k`.** Once a prefix product hits `0 mod k`, every
  longer prefix stays `0`. Nothing special is needed, but it is a good sanity
  check: `[3,3,4]` with `k = 2` gives residues `1, 1, 0`.
- **`start == n - 1`.** The query range is a single element; make sure the suffix
  decomposition still includes it and that the empty prefix is never counted (the
  removed suffix must leave `nums` non-empty, so `j` starts at `start`, not
  `start - 1`).
- **`n == 1`.** The tree is a single leaf; `index` and `start` are both forced to
  `0`. Recursive builds with `lo == hi` at the root must not recurse.
- **Updates persist.** `nums[index] = value` carries into all later queries, so
  the tree must be mutated rather than rebuilt per query — and a query may update
  the same index repeatedly, or update an index that lies *before* `start` and
  therefore does not affect that query's answer.
- **Large values.** `nums[i]` and `value` go up to `10^9`; reduce mod `k` at the
  leaf so products never overflow.
- **Order sensitivity.** The merge is non-commutative. Swapping the children, or
  visiting the query's segments out of order, silently produces wrong answers on
  asymmetric inputs — the randomized brute-force test in
  `solution_test.go` exists to catch exactly that.
