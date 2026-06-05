# 3753. Total Waviness of Numbers in Range II

[LeetCode Link](https://leetcode.com/problems/total-waviness-of-numbers-in-range-ii/)

Difficulty: Hard
Topics: Math, Dynamic Programming
Acceptance Rate: 42.9%

This one is genuinely hard — it asks you to aggregate a per-digit property over a
range that can hold up to 10^15 numbers, so brute force is hopeless. But the
structure is friendly once you spot it, and the technique (digit DP) is a
high-value pattern that shows up again and again. Take it in stages.

## Hints

### Hint 1

You cannot iterate over the range — it may contain ~10^15 numbers. Whenever a
problem says "sum/count some property over every integer in `[L, R]`," reach for
the **count-up-to-N** reduction: define `g(N)` as the answer over `[0, N]`, and
return `g(R) - g(L-1)`. That turns one range query into two prefix queries. Now
ask: what kind of DP lets me compute a digit-based statistic over `[0, N]`?

### Hint 2

The tool is **digit DP**: build numbers one digit at a time from the most
significant position, while carrying a "tight" flag that says whether the prefix
you've placed so far still hugs `N`'s prefix. The twist here is that you are not
counting numbers — you are summing a quantity. So each DP state should return
**two** things: how many numbers complete from here, and the **total waviness**
those completions accumulate.

### Hint 3

Waviness is purely *local*: whether a digit is a peak or valley depends only on
itself and its two immediate neighbors. So you only need to remember the **last
two real digits** as you move left to right. The key realization: a peak/valley
at some position can be *finalized the instant you place the digit to its right*.
If `prev2` (left), `prev1` (candidate middle), and the new digit `d` are all real
digits, then `prev1` contributes 1 to waviness exactly when
`prev1 > prev2 && prev1 > d` (peak) or `prev1 < prev2 && prev1 < d` (valley).
Add that contribution to every number flowing through the transition. Leading
zeros, and the "first/last digit can't be a peak" rule, fall out for free if you
track "no real digit yet" as a sentinel.

## Approach

**Reduction.** Define `g(N)` = sum of waviness over all integers in `[0, N]`. The
inclusive-range answer is `g(num2) - g(num1 - 1)`.

**State.** Process `N`'s digits most-significant first. At each position we carry:

- `pos` — the index we are about to fill.
- `tight` — whether the prefix so far equals `N`'s prefix (limits the next digit).
- `prev1` — the most recent **real** digit (the current peak/valley candidate).
- `prev2` — the real digit before that (the candidate's left neighbor).

We encode "no real digit placed yet" with the sentinel `10`. This single trick
handles leading zeros (a leading `0` keeps both `prev` values at the sentinel and
is not treated as a digit) **and** the boundary rules: the first real digit never
has a left neighbor, and the last digit never gets a right neighbor, so neither
can ever be scored as a peak/valley.

**Transition.** For each candidate digit `d` from `0` to the limit
(`N`'s digit at `pos` when tight, else `9`):

- If we are still in the leading-zero prefix (`prev1 == sentinel && d == 0`),
  stay in the prefix; nothing is scored.
- Otherwise `d` is a real digit. If a real `prev2` exists, then `prev1` is now an
  interior digit, and it contributes `1` iff it is a strict peak or strict valley
  relative to `prev2` and `d`. Shift the window: `prev2 ← prev1`, `prev1 ← d`.

**Combining.** Each state returns `(cnt, sum)`. A child returns how many numbers
complete (`cnt`) and their accumulated waviness (`sum`). The parent folds in the
local contribution `c` of this transition as `sum += childSum + c * childCnt` —
because the contribution applies once to *every* number that passes through.

**Memoization.** Non-tight states `(pos, prev1, prev2)` repeat across many
prefixes, so memoize them. Tight states lie on a single path and aren't cached.

**Worked check.** For `4848`: digits `4,8,4,8`. When we place the third digit
`4`, the window is `prev2=4, prev1=8` → `8 > 4 && 8 > 4`, a peak (+1). When we
place the fourth digit `8`, the window is `prev2=8, prev1=4` → `4 < 8 && 4 < 8`,
a valley (+1). Total waviness `2`, matching the example.

## Complexity Analysis

Let `D ≈ 16` be the number of digits in `N` (since `N ≤ 10^15`).

Time Complexity: O(D · 11 · 11 · 10) per `g(N)` call — positions × `prev1`
states × `prev2` states × the 10 digit choices per transition. With the sentinel
that is a few thousand operations; effectively O(1) for this constraint. We call
`g` twice, so the whole solution is constant-bounded work.

Space Complexity: O(D · 11 · 11) for the memo table plus O(D) recursion depth.

## Edge Cases

- **Numbers with fewer than 3 digits.** They can have no interior digit, so their
  waviness is 0. The sentinel-window logic never scores them because a real
  `prev2` never materializes.
- **Leading zeros while building shorter numbers.** Building fixed-width prefixes
  means a value like `7` is represented as `0…07`; the leading zeros must *not*
  count as digits. The sentinel keeps them out of the waviness window.
- **First and last digit can never be peaks/valleys.** Handled structurally: the
  first real digit is never a `prev1` with a real `prev2`, and the last digit is
  never followed by another digit to finalize it.
- **`num1 - 1` underflow / `num1 = 1`.** `g(0)` must return 0, and `countWaviness`
  guards negative inputs by returning 0 so `g(num1 - 1)` is always well defined.
- **Strictness.** Peaks/valleys require *strict* inequalities, so equal neighbors
  (e.g. the middle `8` in `188`) are neither — easy to get wrong with `>=`/`<=`.
- **Magnitude of the result.** Up to ~10^15 numbers, each with up to ~13 interior
  digits, so the sum can reach ~10^16; use 64-bit integers throughout.
