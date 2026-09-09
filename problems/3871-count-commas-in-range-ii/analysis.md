# 3871. Count Commas in Range II

[LeetCode Link](https://leetcode.com/problems/count-commas-in-range-ii/)

Difficulty: Medium
Topics: Math
Acceptance Rate: 44.7%

## Hints

### Hint 1

`n` can be as large as `10^15`, so any loop that visits every integer from `1` to `n`
is hopeless. That constraint is the whole hint: the answer has to come from a closed-form
count, not from simulation. Ask yourself what property of a single number decides how many
commas it gets, and whether you can count how many numbers in `[1, n]` share that property
without enumerating them.

### Hint 2

For one number `x` with `d` digits, the comma count is `(d - 1) / 3` (integer division):
4-6 digits get 1 comma, 7-9 digits get 2, and so on. So the total is
`sum over x in [1, n] of (digits(x) - 1) / 3`.

You could group numbers by digit length — there are `9 * 10^(d-1)` numbers with exactly `d`
digits — and multiply each group size by its comma count. That works, but there is a
reformulation that is even shorter and avoids the "partial last group" bookkeeping.

### Hint 3

Instead of asking "how many commas does each number have," flip it and ask
"how many numbers does each comma position appear in."

A number gets its **k-th** comma exactly when it has at least `3k + 1` digits, i.e. when
`x >= 10^(3k)`. So the k-th comma contributes one unit for every `x` in `[10^(3k), n]`.

That turns the whole problem into a sum of at most five terms:

```
answer = sum over k >= 1 of max(0, n - 10^(3k) + 1)
```

Since `n <= 10^15`, only `k = 1..5` can ever contribute (`10^3, 10^6, 10^9, 10^12, 10^15`).

## Approach

The key move is exchanging the order of summation. The naive view is

```
answer = Σ_{x=1}^{n} commas(x)
```

where `commas(x) = (digits(x) - 1) / 3`. Rewrite `commas(x)` as an indicator sum:

```
commas(x) = Σ_{k>=1} [ x has a k-th comma ]
```

A number written in standard formatting places a comma after every group of three digits
counted from the right. The first comma appears once the number reaches 4 digits, the
second once it reaches 7 digits, the third at 10 digits, and generally the k-th comma
appears exactly when the number has at least `3k + 1` digits — which is precisely
`x >= 10^(3k)`.

Substituting and swapping the two sums:

```
answer = Σ_{x=1}^{n} Σ_{k>=1} [ x >= 10^(3k) ]
       = Σ_{k>=1} Σ_{x=1}^{n} [ x >= 10^(3k) ]
       = Σ_{k>=1} |{ x : 10^(3k) <= x <= n }|
       = Σ_{k>=1} max(0, n - 10^(3k) + 1)
```

Each inner count is just the size of a contiguous integer interval, which is why no
enumeration is needed. The algorithm is then a loop over powers of one thousand:

1. Start with `total = 0` and `p = 1000`.
2. While `p <= n`, add `n - p + 1` to `total` and multiply `p` by `1000`.
3. Return `total`.

The loop runs at most 5 times for the given constraints (`1000, 10^6, 10^9, 10^12, 10^15`),
so it is effectively constant time.

**Worked example, `n = 1002`:**

- `p = 1000`: `1000 <= 1002`, so add `1002 - 1000 + 1 = 3`. These are `1,000`, `1,001`,
  `1,002` — each earning its first comma.
- `p = 1000000`: `1000000 > 1002`, loop ends.
- Total = `3`, matching the expected output.

**Worked example, `n = 1000002`:**

- `p = 1000`: add `1000002 - 999 = 999003` (every number from 1000 up gets a first comma).
- `p = 1000000`: add `1000002 - 1000000 + 1 = 3` (the three 7-digit numbers each get a
  second comma).
- `p = 10^9`: too big, stop.
- Total = `999006`.

Multiplying `p` by 1000 each iteration is the natural way to walk the thresholds, but note
the overflow trap discussed below — the loop guard must be written so `p` never has to grow
past a representable value.

## Complexity Analysis

Time Complexity: O(log n) — more precisely `O(log_1000 n)`, which is at most 5 iterations
for `n <= 10^15`, so constant in practice.
Space Complexity: O(1) — only a running total and the current power of one thousand.

## Edge Cases

- **`n < 1000` (including `n = 1`, the minimum).** No number has four digits, so the loop
  body never executes and the answer is `0`. This is Example 2 (`n = 998`) and is the main
  reason the `max(0, ...)` / loop guard matters: without it you would add a negative term.
- **`n = 999` vs `n = 1000`.** The boundary where the first comma appears. `999` yields `0`,
  `1000` yields exactly `1`. Off-by-one errors here usually come from writing `n - p`
  instead of `n - p + 1` (the interval `[p, n]` is inclusive on both ends).
- **`n` exactly equal to a power of one thousand** (`10^6`, `10^9`, ...). The new comma
  threshold contributes exactly `1`, not `0`. Using `p <= n` rather than `p < n` is what
  gets this right.
- **`n = 10^15`, the maximum.** Five thresholds all contribute, and the answer is
  `3998998998999005` — far beyond 32-bit range. The accumulator must be a 64-bit type.
- **Overflow while advancing `p`.** If you multiply `p` by 1000 unconditionally, after the
  `10^15` iteration `p` becomes `10^18`, which still fits in `int64` — but one more step
  would be `10^21` and would overflow. Guarding the multiplication (or bounding the loop by
  a fixed count of five) keeps this safe regardless of how the loop is structured.
- **Numbers with a leading digit group shorter than three** (e.g. `1,000` has a one-digit
  lead group). This is not a special case at all under the threshold formulation, which is
  a nice sanity check that the reformulation is the right one — the digit-length grouping
  approach would need explicit care here.
