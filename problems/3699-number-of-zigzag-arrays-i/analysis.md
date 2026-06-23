# 3699. Number of ZigZag Arrays I

[LeetCode Link](https://leetcode.com/problems/number-of-zigzag-arrays-i/)

Difficulty: Hard
Topics: Dynamic Programming, Prefix Sum
Acceptance Rate: 40.6%

This is a Hard counting problem, but don't let that scare you off. The "trick"
is recognizing that a counting-with-constraints problem over arrays almost
always collapses into a clean dynamic program once you find the right state.
Take it one hint at a time.

## Hints

### Hint 1

You are asked to *count* arrays, not construct them, and `n` can be up to 2000.
That rules out enumerating every array. Whenever you need to count sequences
built element-by-element under a local constraint (a rule that only looks at the
last element or two), think **dynamic programming over the position in the
array**. Ask yourself: "What is the minimal information about the prefix I need
to know in order to decide what the next element can be?"

### Hint 2

Look closely at the constraint "no three consecutive elements are strictly
increasing or strictly decreasing." Combined with "no two adjacent elements are
equal," this forces the array to **strictly alternate direction**: up, down, up,
down, ... (or down, up, down, up, ...). So to extend a valid prefix, the only
things you need to know are the **last value** and **whether the last step went
up or down** — because the next step's direction is forced to be the opposite.
That suggests a DP state of `(direction, last value)`.

### Hint 3

Define two arrays at each length `i`:

- `up[v]` = number of valid arrays of length `i` whose last element is `v` and
  whose final step went **up** (previous element `< v`).
- `down[v]` = number of valid arrays of length `i` whose last element is `v` and
  whose final step went **down** (previous element `> v`).

The transition is:

- `up[i][v]   = sum of down[i-1][u] for all u < v`
- `down[i][v] = sum of up[i-1][u]   for all u > v`

(An up-step must follow a down-step, and vice versa.) Computing those sums
naively is `O(m)` per value, giving `O(n * m^2)` which is too slow. The final
insight: the inner sums are just a **prefix sum** (for `up`) and a **suffix
sum** (for `down`) over the previous row. Maintain running sums and each
transition becomes `O(m)` per length, for `O(n * m)` overall.

## Approach

Let `m = r - l + 1` be the number of allowed distinct values. Re-index the
values `l..r` as `0..m-1`; their actual magnitudes don't matter, only their
relative order, so we work with indices.

**Base case (length 1).** A single element has no "previous" step, so it is
trivially valid in either orientation. We seed `up[v] = down[v] = 1` for every
value `v`. Counting each single-element array under both labels is a convenient
fiction that makes the length-2 transition come out right (and at length 2 each
array still gets counted exactly once, because a real two-element array has a
definite direction).

**Transition (length i from length i-1).** Using the recurrences from Hint 3:

- `newUp[v] = Σ_{u < v} down[u]` — extend every array that ended going *down* by
  appending a larger value `v`, making the new step go *up*.
- `newDown[v] = Σ_{u > v} up[u]` — extend every array that ended going *up* by
  appending a smaller value `v`, making the new step go *down*.

We evaluate `newUp` with a left-to-right **prefix sum** over `down`, and
`newDown` with a right-to-left **suffix sum** over `up`. Both are computed from
the *old* row, so we build both new rows before overwriting.

**Answer.** After processing length `n`, every valid array ends with a definite
direction, so the total is `Σ_v (up[v] + down[v])`, taken modulo `1e9 + 7`.

**Worked example** (`n = 3, l = 1, r = 3`, so `m = 3`, indices `0,1,2`):

- Length 1: `up = [1,1,1]`, `down = [1,1,1]`.
- Length 2: `up = [0,1,2]` (prefix sums of `down`), `down = [2,1,0]` (suffix
  sums of `up`).
- Length 3: `up = [0,2,3]` (prefix sums of the new `down`), `down = [3,2,0]`
  (suffix sums of the new `up`).
- Total = `(0+2+3) + (3+2+0) = 5 + 5 = 10`. ✓

## Complexity Analysis

Let `m = r - l + 1`.

Time Complexity: O(n * m) — we run `n - 1` transition steps, each doing two
linear passes over the `m` values.
Space Complexity: O(m) — we keep only the current and next `up`/`down` rows;
the row index is not retained.

## Edge Cases

- **`n = 3` (smallest allowed length).** The DP still runs two transitions and
  must produce the right alternating count; this is the canonical example case.
- **Single-step value range (`r = l + 1`, so `m = 2`).** Only two distinct
  values are available, so arrays must strictly alternate between them (e.g.
  `[4,5,4]` and `[5,4,5]`). Make sure prefix/suffix sums don't double-count.
- **Modulo discipline.** With `n` and the range both up to 2000, counts blow up
  fast. Apply `% (1e9+7)` inside the running sums and at the final answer so no
  intermediate value overflows or is forgotten.
- **No off-by-one in the sums.** `up[v]` must sum *strictly* smaller indices and
  `down[v]` *strictly* larger ones — the prefix/suffix must exclude the current
  index `v` itself, mirroring the "no two adjacent equal" rule.
