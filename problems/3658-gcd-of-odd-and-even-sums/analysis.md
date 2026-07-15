# 3658. GCD of Odd and Even Sums

[LeetCode Link](https://leetcode.com/problems/gcd-of-odd-and-even-sums/)

Difficulty: Easy
Topics: Math, Number Theory
Acceptance Rate: 86.4%

## Hints

### Hint 1

Before writing any loops, try to find closed-form expressions for the two sums. The sum of the first `n` odd numbers and the sum of the first `n` even numbers are both classic arithmetic series with well-known formulas. Recognizing them turns this into a pure number-theory question.

### Hint 2

Work out each formula. The first `n` odd numbers are `1, 3, 5, ..., (2n-1)` and the first `n` even numbers are `2, 4, 6, ..., 2n`. Once you have `sumOdd` and `sumEven` as simple expressions in `n`, look at how they factor. A shared factor between them will jump out.

### Hint 3

The sum of the first `n` odd numbers is exactly `n²`, and the sum of the first `n` even numbers is `n(n+1)`. Since `n` and `n+1` are consecutive integers, they are coprime — `gcd(n, n+1) = 1`. Therefore `gcd(n², n(n+1)) = n · gcd(n, n+1) = n`. The answer is simply `n`.

## Approach

The problem asks for `gcd(sumOdd, sumEven)`, so start by simplifying both sums.

- **Sum of the first `n` odd numbers:** `1 + 3 + 5 + ... + (2n-1)`. This is a well-known identity that equals `n²`. (Quick check for `n = 4`: `1 + 3 + 5 + 7 = 16 = 4²`.)
- **Sum of the first `n` even numbers:** `2 + 4 + 6 + ... + 2n = 2(1 + 2 + ... + n) = 2 · n(n+1)/2 = n(n+1)`. (Quick check for `n = 4`: `2 + 4 + 6 + 8 = 20 = 4 · 5`.)

Now we need `gcd(n², n(n+1))`. Factor out the common `n`:

```
gcd(n², n(n+1)) = n · gcd(n, n+1)
```

Because `n` and `n+1` are consecutive integers, their greatest common divisor is always `1` (any common divisor would have to divide their difference, which is `1`). So:

```
gcd(n², n(n+1)) = n · 1 = n
```

The result is just `n`. This gives an `O(1)` solution that returns the input directly.

For learning value, the accompanying solution still computes `sumOdd = n²` and `sumEven = n(n+1)` explicitly and runs the Euclidean algorithm on them. This mirrors the literal problem statement, is easy to trust, and — reassuringly — always produces `n`, confirming the derivation. Either the one-line `return n` or the explicit-GCD version is fully correct; the explicit version is shown to make the reasoning tangible.

Example walk-through for `n = 5`:
- `sumOdd = 5² = 25`
- `sumEven = 5 · 6 = 30`
- `gcd(25, 30) = 5`, which equals `n`. ✓

## Complexity Analysis

Time Complexity: O(1) — with the `return n` shortcut. The explicit Euclidean-algorithm variant is O(log n), which is still effectively constant here.
Space Complexity: O(1) — only a few integer variables are used.

## Edge Cases

- **`n = 1`:** `sumOdd = 1`, `sumEven = 2`, `gcd(1, 2) = 1`, which equals `n`. The smallest input still fits the formula.
- **Upper bound `n = 1000`:** `sumOdd = 1_000_000` and `sumEven = 1_001_000` comfortably fit in a 32-bit `int`, but Go's `int` is 64-bit on common platforms, so there is no overflow risk. Worth confirming that intermediate products don't overflow for larger constraints.
- **Overflow awareness in general:** if the constraints were much larger, `n²` and `n(n+1)` could exceed a fixed-width integer. Here the constraints are tiny, so ordinary `int` arithmetic is safe.
