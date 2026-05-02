# 0788. Rotated Digits

[LeetCode Link](https://leetcode.com/problems/rotated-digits/)

Difficulty: Medium
Topics: Math, Dynamic Programming
Acceptance Rate: 58.1%

## Hints

### Hint 1

This is a counting problem over a small range (`n <= 10^4`). Before reaching for anything fancy, ask yourself: can you decide whether a *single* number is "good" just by looking at its digits one at a time? If so, you can simply iterate from `1` to `n` and count.

### Hint 2

Bucket each digit into one of three categories based on its 180° rotation behavior:

- **Invalid digits** — rotating these produces nothing valid: `3, 4, 7`.
- **Self-rotating digits** — they look the same after rotation: `0, 1, 8`.
- **Changing digits** — they rotate to a different valid digit: `2 ↔ 5`, `6 ↔ 9`.

Now restate the "good" condition using only these buckets.

### Hint 3

A number `x` is good **iff** both:

1. None of its digits is in the invalid bucket, AND
2. At least one of its digits is in the changing bucket (otherwise rotating produces the same number).

So you only need two flags while walking the digits: "saw an invalid digit?" (disqualify) and "saw a changing digit?" (required for "different from `x`"). No actual rotation arithmetic is needed.

## Approach

The clean idea is to convert the rotation rule into a digit-classification rule and check each number in `[1, n]` independently.

**Algorithm**

1. For each `x` in `[1, n]`:
   - Walk each decimal digit of `x`.
   - If you see `3`, `4`, or `7`, mark the number as invalid and stop early.
   - If you see `2`, `5`, `6`, or `9`, set a "has changing digit" flag.
   - Digits `0`, `1`, `8` are fine but contribute nothing toward the "must differ" requirement.
2. After scanning, count `x` if it was not invalid **and** had at least one changing digit.

**Why it works**

Rotation is digit-local — every digit rotates independently — so the validity of a number reduces to checking each digit's class. The "must differ from `x`" condition is exactly "at least one digit actually changed under rotation," which means at least one digit from `{2, 5, 6, 9}`.

**Example (`n = 10`)**

- `1` → only self-rotating digits → not good.
- `2` → contains a changing digit `2`, no invalid → good.
- `3, 4, 7` → invalid.
- `5, 6, 9` → good (each has a changing digit).
- `8` → only self-rotating → not good.
- `10` → digits `1, 0`, both self-rotating → not good.

Count: `{2, 5, 6, 9}` → `4`. ✓

**Note on Dynamic Programming**

For much larger `n` (e.g., `10^18`), this problem is a textbook *digit DP*: count numbers `≤ n` whose digit pattern satisfies the two flags above, using a state of `(position, tight, has_changing)`. For `n ≤ 10^4` the linear scan is more than fast enough and far simpler.

## Complexity Analysis

Time Complexity: O(n · log₁₀ n) — we visit each number in `[1, n]` and inspect its `O(log n)` digits.
Space Complexity: O(1) — only a few flags and counters.

## Edge Cases

- `n = 1`: the only candidate is `1`, which rotates to itself → answer is `0`.
- `n = 2`: just `2` qualifies → answer is `1`. Important sanity check that single-digit changing numbers count.
- Numbers like `8`, `11`, `101` that only contain `0/1/8` — valid rotation but unchanged → must **not** be counted.
- Numbers containing any of `3, 4, 7` — must be rejected even if they also contain changing digits (e.g., `23` has a `2`, but the `3` disqualifies it).
- Multi-digit numbers ending in `0` like `10`, `100` — make sure your digit walk handles trailing zeros correctly (`n % 10`, `n /= 10` does).
