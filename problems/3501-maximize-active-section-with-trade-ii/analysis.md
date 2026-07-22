# 3501. Maximize Active Section with Trade II

[LeetCode Link](https://leetcode.com/problems/maximize-active-section-with-trade-ii/)

Difficulty: Hard
Topics: Array, String, Binary Search, Segment Tree
Acceptance Rate: 47.5%

This is a genuinely hard problem — it combines a non-obvious modeling step
(figuring out what a "trade" really buys you) with a data-structure step
(answering many range queries fast). Take the modeling insight first; the
efficiency machinery only makes sense once you know exactly what quantity you
need to query. Don't be discouraged if the trade definition feels slippery at
first — everyone has to reduce it to something simpler before it clicks.

## Hints

### Hint 1

Before thinking about queries at all, solve the single-string version (this is
the easier companion problem, "Maximize Active Section with Trade I"). Describe
in one sentence what a full trade accomplishes: you blank out a block of `1`s
that sits between two `0` blocks, then flood a block of `0`s that sits between
two `1` blocks. What is the *net* change to the number of `1`s? Try it on
`"0 1 0"`-style fragments and look for a pattern in terms of the surrounding
runs of zeros.

### Hint 2

A trade only ever merges **two consecutive runs of zeros** that are separated by
exactly one run of ones. If those zero runs have sizes `z1` and `z2`, the net
gain in active sections is exactly `z1 + z2` (you remove the `1`s between them,
then fill everything back with `1`s). So the whole problem reduces to: *find the
best adjacent pair of zero runs.* Also notice the parts of `s` outside the query
range never change, which means

```
answer = totalOnes(s) + bestGain(range)
```

Now the only question is how to compute `bestGain` for each of up to 10^5
ranges quickly.

### Hint 3

Only the **two extreme** zero runs of a query get clipped by the `[l, r]`
boundary — the augmenting `'1'` glued to each end keeps them "surrounded," so a
clipped zero run is still a legal trade partner. Every *interior* consecutive
pair of zero runs uses its full, precomputable size. So precompute
`pairSum[i] = size(zeroRun[i]) + size(zeroRun[i+1])` and answer range-maximum
queries over it with a sparse table (or segment tree). For each query, use
binary search to find the first and last zero runs touching `[l, r]`, take the
interior range-max, and separately compute the two boundary pairs with clipped
sizes.

## Approach

**Step 1 — Reduce the trade to a run-merge.**
A valid trade blanks a `1`-block surrounded by `0`s and then fills a `0`-block
surrounded by `1`s. The optimal move is always: pick a `1`-run that sits between
two `0`-runs, delete it (turning `z1 0…0 z2` into one big `0`-block), then fill
that entire merged `0`-block with `1`s. If the two zero runs have sizes `z1` and
`z2`, you converted `z1 + z2` former zeros into ones — a net gain of `z1 + z2`.
No trade at all gives gain `0`, and a trade is only possible when at least two
zero runs are separated by a single one run.

**Step 2 — Reframe the answer.**
The trade happens strictly inside `s[l..r]`; the rest of `s` is untouched. So the
active-section count of the whole final string is

```
answer = totalOnes(s) + bestGain(l, r)
```

where `bestGain(l, r)` is the largest `z1 + z2` over adjacent zero runs that lie
inside the augmented substring `t = '1' + s[l..r] + '1'`.

**Step 3 — Precompute.**
Decompose `s` into maximal zero runs `zStart[i]..zEnd[i]`. In a run
decomposition, consecutive zero runs are *always* separated by exactly one
`1`-run, so every adjacent zero-run pair is a valid trade candidate. Build
`pairSum[i] = size(i) + size(i+1)` and a **sparse table** over it for O(1)
range-maximum queries.

**Step 4 — Answer a query.**
Use binary search on the (sorted) run boundaries:

- `a` = first zero run with `zEnd[a] >= l`,
- `b` = last zero run with `zStart[b] <= r`.

If fewer than two zero runs overlap `[l, r]` (`b <= a`), gain is `0`. Otherwise
the zero runs `a..b` are exactly those inside the range; only `a` and `b` may be
clipped by the boundary. The best gain is the maximum of:

- the **left boundary pair**: `clip(a) + size(a+1)`,
- the **right boundary pair**: `size(b-1) + clip(b)`,
- the **interior pairs**: `rangeMaxPair(a+1, b-2)` from the sparse table,

where `clip(i)` is the length of zero run `i` intersected with `[l, r]`. (When
there are exactly two zero runs, both are boundary/clipped and the gain is simply
`clip(a) + clip(b)`.)

**Worked example.** `s = "1000100"`, query `[1, 5]`. Zero runs are `[1..3]`
(size 3) and `[5..6]` (size 2), and `totalOnes = 2`. The query touches both runs;
run `[1..3]` clips to length 3 and run `[5..6]` clips to length 1 (only index 5
is in range). Gain `= 3 + 1 = 4`, so `answer = 2 + 4 = 6`, matching the expected
output.

## Complexity Analysis

Time Complexity: O((n + q) log n) — O(n log n) to build the sparse table plus
O(log n) per query for the two binary searches and an O(1) range-max lookup,
across `q` queries.
Space Complexity: O(n log n) for the sparse table (O(n) for the run arrays).

## Edge Cases

- **No zero runs (all `1`s):** no trade is possible, so every answer is
  `totalOnes` with gain `0`.
- **A single zero run in range (including all `0`s):** a trade needs a `1`-block
  between two `0`-blocks, which doesn't exist, so gain is `0`.
- **Boundary zero run extends outside `[l, r]`:** it must be clipped to its
  intersection with the range; the augmenting `'1'` keeps it a valid trade
  partner even though it's truncated (see example 3, query `[1, 5]`).
- **Interior pair beats both boundary pairs:** a clipped boundary run can be
  small, so the best pair may be entirely in the middle — the range-max query
  over full-size interior pairs catches this.
- **Exactly two overlapping zero runs:** both are clipped; handle this directly
  instead of assuming one side is a full interior run.
- **Single-character string / smallest inputs:** the run decomposition and
  binary searches must still behave (gain stays `0`).
