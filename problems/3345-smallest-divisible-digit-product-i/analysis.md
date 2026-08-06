# 3345. Smallest Divisible Digit Product I

[LeetCode Link](https://leetcode.com/problems/smallest-divisible-digit-product-i/)

Difficulty: Easy
Topics: Math, Enumeration
Acceptance Rate: 68.0%

## Hints

### Hint 1

Look hard at the constraints before designing anything clever: `n <= 100` and `t <= 10`. When the input space is this tiny, the question is usually not "what is the smart formula?" but "how small is the search space, and can I just walk it?" Think about the *enumeration* pattern — start at the answer's lower bound and test candidates one at a time.

### Hint 2

The test itself is mechanical: peel digits off a number with repeated `% 10` and `/ 10`, multiply them together, and check divisibility by `t`. The only thing standing between you and a solution is convincing yourself the walk upward from `n` terminates quickly. Ask: is there any candidate whose digit product is divisible by *every* possible `t`?

### Hint 3

Yes — any number containing the digit `0` has a digit product of `0`, and `0 % t == 0` for every `t >= 1`. Starting from `n`, the next multiple of ten is at most `n + 9` away, and every multiple of ten ends in `0`. So the loop is guaranteed to stop within 10 iterations no matter what. That turns the "brute force" into a constant-time algorithm, and it also explains Example 1: `10` works immediately because its digit product is `0`, not because `0` is somehow special-cased.

## Approach

The problem asks for the smallest `x >= n` whose digit product is divisible by `t`. Two observations make this straightforward.

**1. A cheap predicate.** For a candidate `x`, compute the digit product by repeatedly taking `x % 10` and dividing `x` by `10`, multiplying the remainders together. Then check `product % t == 0`. Note the digit product of a number containing `0` is `0`, and `0 % t == 0` holds for any positive `t` — so such numbers always pass.

**2. A tight bound on the search.** Because "smallest `x >= n`" means we must check candidates in increasing order anyway, the natural algorithm is: try `n`, then `n+1`, then `n+2`, and return the first one that passes. The worry with any unbounded upward scan is termination — here it is a non-issue. From any starting point `n`, the next multiple of `10` is at most `9` steps away, and every multiple of `10` ends in the digit `0`, hence has digit product `0`, hence passes for every `t`. So the loop runs **at most 10 iterations**, guaranteed.

Walking through Example 2 (`n = 15, t = 3`):

- `15` → digits `1, 5` → product `5` → `5 % 3 = 2` ✗
- `16` → digits `1, 6` → product `6` → `6 % 3 = 0` ✓ → return `16`

And Example 1 (`n = 10, t = 2`):

- `10` → digits `1, 0` → product `0` → `0 % 2 = 0` ✓ → return `10`

The answer is the first hit, which by construction is the smallest, since we test in strictly increasing order.

This is an honest "easy" problem, but it is a good one for building the reflex of *reading the constraints first*. A reasonable interviewer's follow-up is the harder variant (LeetCode 3348, "Smallest Divisible Digit Product II") where `n` has up to 10^5 digits and enumeration is hopeless — there you must factor `t` into 2s, 3s, 5s, and 7s and construct the answer digit by digit. Recognizing *why* enumeration is legal here is what tells you when it stops being legal.

## Complexity Analysis

Time Complexity: O(log n) — the loop body runs at most 10 times (a constant, independent of the input), and each iteration extracts the digits of a number in O(log n). With the given constraints this is effectively O(1).

Space Complexity: O(1) — only a handful of integer accumulators; digits are consumed one at a time rather than materialized into a slice.

## Edge Cases

- **`t = 1`** — every product is divisible by `1`, including the product `1` of the single digit `1`, so the answer is always `n` itself. Make sure the loop returns on the very first candidate rather than assuming some work must happen.
- **`n` contains a `0` digit (e.g. `n = 10`, `n = 100`)** — the digit product is `0`, and `0 % t == 0` for all valid `t`, so the answer is `n`. This is Example 1, and it is the case people misread as "0 is not divisible."
- **`t = 10`** — no single digit is divisible by `10`, and a product of nonzero digits can only reach a multiple of `10` by containing both a `2`/`4`/`6`/`8` and a `5`. For small `n` this forces the answer to a number containing a `0`, e.g. `n = 1, t = 10` → `10`. Good test of the termination argument.
- **Crossing a decade or the 99 → 100 boundary** — e.g. `n = 99, t = 7`: `99` gives `81` ✗, and the scan rolls over to `100`, which gives `0` ✓. Verifies the digit extraction handles a changing digit count.
- **Answer exceeding `n`'s upper constraint** — `n` is capped at `100`, but the *answer* is not; it can be `100` or beyond. Do not bound the loop by `100`.
- **Integer overflow (non-issue here, real elsewhere)** — with at most 3 digits the product fits trivially, but the digit-product idiom is worth writing overflow-consciously as a habit for the larger variants.
