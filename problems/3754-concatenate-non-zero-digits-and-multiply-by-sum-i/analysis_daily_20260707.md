# 3754. Concatenate Non-Zero Digits and Multiply by Sum I

[LeetCode Link](https://leetcode.com/problems/concatenate-non-zero-digits-and-multiply-by-sum-i/)

Difficulty: Easy
Topics: Math

Acceptance Rate: 60.8%

## Hints

### Hint 1

This is a digit-manipulation problem. Think about how you can inspect the individual
digits of an integer one at a time. Two classic tools come to mind: converting the
number to a string, or repeatedly using `% 10` and `/ 10`. Either can walk the digits.

### Hint 2

You need to build a new number `x` from only the **non-zero** digits, keeping their
original left-to-right order. As you keep each non-zero digit, you also need its
running sum. Notice that both `x` and `sum` can be accumulated in a single pass — you
do not need to store the digits separately first.

### Hint 3

If you scan the digits from most-significant to least-significant, you can grow `x`
with the rule `x = x*10 + digit` for every non-zero digit, and simultaneously add that
digit to `sum`. Skip any zero digit entirely. If no non-zero digit is ever seen, `x`
stays `0` (which also makes `sum` `0`), so the answer is naturally `0`. Finally return
`x * sum`.

## Approach

The task decomposes into two coupled accumulations over the digits of `n`:

1. **Build `x`** — the integer formed by the non-zero digits in original order.
2. **Compute `sum`** — the sum of the digits of `x` (equivalently, the sum of the
   non-zero digits of `n`, since zeros contribute nothing).

The cleanest way to preserve original order is to process digits from most significant
to least significant. The simplest way to get that order is to turn `n` into its decimal
string and iterate left to right. For each character:

- Convert it to its digit value `d`.
- If `d != 0`, extend `x` with `x = x*10 + d` and add `d` to `sum`.
- If `d == 0`, ignore it.

Walk through Example 1 with `n = 10203004`:

| digit | non-zero? | x after | sum after |
|-------|-----------|---------|-----------|
| 1     | yes       | 1       | 1         |
| 0     | no        | 1       | 1         |
| 2     | yes       | 12      | 3         |
| 0     | no        | 12      | 3         |
| 3     | yes       | 123     | 6         |
| 0     | no        | 123     | 6         |
| 0     | no        | 123     | 6         |
| 4     | yes       | 1234    | 10        |

Result: `x * sum = 1234 * 10 = 12340`. ✅

For Example 2, `n = 1000`, only the leading `1` survives, giving `x = 1`, `sum = 1`,
and `1 * 1 = 1`. ✅

An alternative is a pure-math approach using `% 10` / `/ 10`, but that yields digits
from least significant first, so you would have to reverse or use place multipliers to
rebuild `x` in the correct order. The string scan avoids that bookkeeping and stays
easy to read, which is why the reference solution uses it.

Because `0 <= n <= 10^9`, `n` has at most 10 digits, so `x` fits comfortably in a
64-bit (and even a 32-bit) integer, and `x * sum` cannot overflow Go's `int`.

## Complexity Analysis

Time Complexity: O(d) where d is the number of digits in `n` (at most 10) — effectively O(1).
Space Complexity: O(d) for the string representation (O(1) if using the pure-math variant).

## Edge Cases

- **`n = 0`**: There are no non-zero digits, so `x = 0`, `sum = 0`, and the answer is
  `0`. The single `0` digit is skipped, leaving the accumulators at their initial `0`.
- **No non-zero digits beyond a leading value that is itself zero**: Same as above —
  any all-zero input collapses to `0`.
- **Trailing/interior zeros (`n = 1000`, `n = 10203004`)**: Zeros must be dropped, not
  treated as placeholders. Building `x` only on non-zero digits handles this correctly.
- **Single non-zero digit (`n = 7`)**: `x = 7`, `sum = 7`, answer `49`. Confirms the
  loop works with one iteration.
- **Maximum input (`n = 10^9 = 1000000000`)**: Only the leading `1` is non-zero, so
  `x = 1`, `sum = 1`, answer `1`. Confirms no overflow at the constraint boundary.
