# 1621. Number of Sets of K Non-Overlapping Line Segments

[LeetCode Link](https://leetcode.com/problems/number-of-sets-of-k-non-overlapping-line-segments/)

Difficulty: Medium
Topics: Math, Dynamic Programming, Combinatorics, Prefix Sum
Acceptance Rate: 58.2%

## Hints

### Hint 1

Nothing about the points matters except their order — point `i` sits at `x = i`, so a
segment is fully described by the pair of indices `(a, b)` with `a < b`. A set of `k`
non-overlapping segments is therefore just a sequence of index pairs

```
a1 < b1 <= a2 < b2 <= a3 < b3 <= ... <= ak < bk
```

(sorted left to right; `b_i <= a_{i+1}` because segments may touch at an endpoint but
may not overlap). So the question is: *how many such sequences of `2k` indices are
there?* That reframing alone is most of the work. Think "count monotone sequences,"
either with dynamic programming or with a counting argument.

### Hint 2

The straightforward route is DP. Scan the points left to right and track two things:
how many segments you have completed, and whether the segment you are currently
building is still "open." That gives states like `dp[i][j][0/1]` = ways to process the
first `i` points having finished `j` segments, with the last one closed (`0`) or still
open (`1`). Every transition is O(1), so the whole thing is O(n·k) — fast enough for
`n <= 1000`. Write that down and check it against `n = 4, k = 2 -> 5`.

But look at the chain of inequalities again. It is *almost* strictly increasing: the
only non-strict steps are the `b_i <= a_{i+1}` ones, and there are exactly `k - 1` of
them. Can you make those steps strict?

### Hint 3

Yes — shift each segment right by how many segments precede it. Map the `i`-th
segment `(a_i, b_i)` (1-indexed) to `(a_i + i - 1, b_i + i - 1)`. Each `<=` in the
chain becomes `<`, so the `2k` shifted values are **strictly increasing**. They live in
`[0, (n - 1) + (k - 1)] = [0, n + k - 2]`, a range with `n + k - 1` distinct values.

The map is reversible: given any strictly increasing `2k` values in that range,
subtract `i - 1` from the `i`-th pair to recover a valid configuration. It is a
bijection, so the answer is a single binomial coefficient:

```
answer = C(n + k - 1, 2k)  (mod 1e9 + 7)
```

## Approach

The solution used here is the closed form. The reasoning, in full:

**Step 1 — normalize the configuration.** A set of `k` non-overlapping segments has a
canonical left-to-right ordering. Writing the `i`-th segment as `(a_i, b_i)`, the
constraints "each segment covers two or more points" and "segments may share endpoints
but not overlap" become exactly

```
0 <= a1 < b1 <= a2 < b2 <= ... <= ak < bk <= n - 1
```

Sets of segments correspond one-to-one with such chains, so counting chains counts
sets.

**Step 2 — make the chain strictly increasing.** The chain alternates between strict
`<` (inside a segment) and non-strict `<=` (between segments). There are `k - 1`
non-strict steps. Define

```
a_i' = a_i + (i - 1)      b_i' = b_i + (i - 1)
```

Adding a larger offset to each successive segment breaks every tie: if `b_i = a_{i+1}`
then `b_i' = b_i + i - 1 < a_{i+1} + i = a_{i+1}'`. Now

```
0 <= a1' < b1' < a2' < b2' < ... < ak' < bk' <= (n - 1) + (k - 1) = n + k - 2
```

**Step 3 — count.** The primed values are `2k` *distinct* numbers chosen from the
`n + k - 1` values `{0, 1, ..., n + k - 2}`, and since they are sorted, the choice of
the set determines the sequence. Conversely, any `2k`-element subset, sorted and
un-shifted, yields a valid configuration (subtracting the offsets preserves the
required inequalities). The correspondence is a bijection, hence

```
answer = C(n + k - 1, 2k)
```

**Step 4 — compute it mod 1e9 + 7.** Build the binomial multiplicatively:

```
C(N, r) = product over i in [0, r) of (N - i) / (i + 1)
```

Accumulate the numerator and denominator separately modulo the prime `p = 1e9 + 7`,
then divide once at the end using Fermat's little theorem: `den^(p-2) ≡ den^(-1)
(mod p)`. Dividing only once keeps the code short and avoids any need for a full
factorial table. Note that `n + k - 1 <= 1998 < p`, so no factor in the denominator is
ever a multiple of `p` and the inverse always exists.

**Worked example (`n = 4, k = 2`).** The five configurations are `{(0,2),(2,3)}`,
`{(0,1),(1,3)}`, `{(0,1),(2,3)}`, `{(1,2),(2,3)}`, `{(0,1),(1,2)}`. Shifting the second
segment right by one turns them into the strictly increasing quadruples
`0<2<3<4`, `0<1<2<4`, `0<1<3<4`, `1<2<3<4`, `0<1<2<3` — exactly the
`C(5, 4) = 5` four-element subsets of `{0,1,2,3,4}`. Likewise `n = 3, k = 1` gives
`C(3, 2) = 3`, and `n = 30, k = 7` gives `C(36, 14) = 3796297200 ≡ 796297179`.

If the bijection does not feel convincing yet, the O(n·k) DP from Hint 2 is a perfectly
good accepted solution and is worth writing first — then compare its table against the
binomials and watch Pascal's rule fall out. Be honest with yourself: the shift trick is
the kind of insight that is obvious *afterwards*, and most people reach it by first
staring at a DP table.

## Complexity Analysis

Time Complexity: O(k + log MOD) — one pass of `2k <= 2(n-1)` multiplications to build
the binomial, plus a single modular exponentiation for the inverse. This is
O(n) overall, versus O(n·k) for the DP formulation.

Space Complexity: O(1) — only a handful of accumulators; no DP table is materialized.

## Edge Cases

- **`k = 1`**: reduces to `C(n, 2)`, the number of index pairs. A good sanity check
  that the formula is not off by one.
- **`k = n - 1` (the maximum allowed)**: `C(n + k - 1, 2k) = C(2n - 2, 2n - 2) = 1` —
  the only option is to chain every unit segment `(0,1), (1,2), ..., (n-2,n-1)`. Any
  off-by-one in the shift would break this case loudly.
- **Smallest input `n = 2, k = 1`**: `C(2, 2) = 1`. Guards against loops that assume
  at least two segments or at least three points.
- **`2k > n + k - 1`**: mathematically impossible under the stated constraints
  (`k <= n - 1`), but the binomial helper still returns `0` for out-of-range `r` rather
  than producing garbage, so the code is safe if the constraints are ever relaxed.
- **Overflow**: every intermediate product is of two values below `1e9 + 7`, which fits
  comfortably in a 64-bit Go `int`. Reducing after each multiplication is what keeps it
  that way — dropping a single `% MOD` silently corrupts large inputs like
  `n = 1000, k = 500`.
- **Modulus handling**: the answer must be taken mod `1e9 + 7` even when the true count
  is small; returning the raw product is wrong only for large inputs, which is exactly
  the kind of bug that passes the sample tests and fails on submission.
