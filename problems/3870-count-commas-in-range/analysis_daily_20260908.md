# 3870. Count Commas in Range

[LeetCode Link](https://leetcode.com/problems/count-commas-in-range/)

Difficulty: Easy
Topics: Math
Acceptance Rate: 73.5%

## Hints

### Hint 1

The constraints are small enough that a straightforward loop over every number in `[1, n]` will pass. But before you write that loop, ask a sharper question: for a *single* number, what determines how many commas it gets? It has nothing to do with the number's value — only with how many digits it has. Once you see that, you're doing counting, not simulation.

### Hint 2

A number with `d` digits gets `(d-1)/3` commas (integer division): 3 digits → 0, 4 through 6 digits → 1, 7 through 9 digits → 2. So the total answer is a sum of `(d-1)/3` over all numbers from 1 to `n`. Instead of computing that per number, try grouping numbers by digit count — how many numbers in `[1, n]` have exactly 4 digits? Exactly 5? You can count each block in O(1).

### Hint 3

There's an even cleaner reframing that avoids grouping entirely. Flip the sum around: instead of asking "how many commas does each number have," ask "how many numbers does each comma *position* apply to."

`(d-1)/3` is exactly the count of integers `k >= 1` with `3k <= d-1`, i.e. `d >= 3k+1`, i.e. the number is at least `10^(3k)`. So the k-th comma slot (thousands, millions, billions, ...) contributes one comma for every number `>= 10^(3k)`.

That turns the whole problem into a tiny sum:

```
answer = sum over k >= 1 of max(0, n - 10^(3k) + 1)
```

## Approach

The key move is exchanging the order of summation. Naively we sum over numbers and, for each, count its commas. Instead we sum over *comma slots* and, for each slot, count how many numbers use it.

**Step 1 — commas per number.** Standard formatting inserts a comma after every three digits from the right. A `d`-digit number therefore has separators at `d-3`, `d-6`, ... as long as those positions leave at least one digit to the left. That count is `(d-1)/3` with integer division.

**Step 2 — reindex the sum.** Notice that `(d-1)/3 = #{k >= 1 : 3k <= d-1}`. The condition `d >= 3k+1` says the number has at least `3k+1` digits, which is the same as saying the number is `>= 10^(3k)`. So:

```
total = sum_{x=1..n} #{k : x >= 10^(3k)}
      = sum_{k >= 1} #{x in [1, n] : x >= 10^(3k)}
      = sum_{k >= 1} max(0, n - 10^(3k) + 1)
```

**Step 3 — evaluate.** Walk `p = 1000, 1000000, 1000000000, ...` and while `p <= n`, add `n - p + 1`. Once `p > n` every later term is zero, so we stop. That's at most a handful of iterations.

**Worked example, `n = 1002`:**

- `p = 1000`: `1000 <= 1002`, add `1002 - 1000 + 1 = 3`. (These are 1000, 1001, 1002 — each written `"1,00x"`.)
- `p = 1000000`: exceeds `n`, stop.
- Answer: **3**. ✓

**Worked example, `n = 1000000`:**

- `p = 1000`: add `1000000 - 1000 + 1 = 999001`.
- `p = 1000000`: add `1000000 - 1000000 + 1 = 1`. (That single number is `"1,000,000"`, whose *second* comma this term accounts for.)
- Answer: **999002**.

Each term isolates one comma column, which is why a number with two commas correctly gets counted twice — once by the thousands term and once by the millions term.

An O(n) loop over `[1, n]` accumulating `(digits-1)/3` is perfectly acceptable for `n <= 10^5` and is a fine thing to write first. The closed form is worth understanding anyway: it's the same "count contributions instead of simulating" trick that shows up in much harder digit-DP and combinatorics problems, and here it's visible without much machinery.

## Complexity Analysis

Time Complexity: O(log n) — one iteration per group of three digits in `n`, at most 6-7 iterations for any 64-bit input.
Space Complexity: O(1) — a couple of integer accumulators.

## Edge Cases

- **`n < 1000` (e.g. `n = 1`, `n = 998`, `n = 999`)** — the loop body never executes and the answer is `0`. This is example 2, and it's the case a solution that assumes at least one comma exists would get wrong.
- **`n = 1000` exactly** — the boundary where the first comma appears. The condition must be `p <= n`, not `p < n`, and the term is `n - p + 1 = 1`, not `n - p = 0`. Off-by-one here is the most likely bug.
- **`n = 999`** — the value just below the boundary, the natural partner test to `n = 1000`.
- **Numbers with more than one comma (`n >= 10^6`)** — outside the stated constraint of `n <= 10^5`, but the reindexed sum handles them for free, and testing one confirms the multi-comma logic. A solution that only ever adds the thousands term would silently pass the official constraints and still be wrong in general.
- **Overflow when advancing `p`** — multiplying `p` by 1000 unconditionally can wrap around for large `n`. Guard the multiplication (check `p > n/1000` before scaling) so the loop terminates cleanly rather than wrapping to a negative value.
