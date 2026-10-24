# 2409. Count Days Spent Together

[LeetCode Link](https://leetcode.com/problems/count-days-spent-together/)

Difficulty: Easy
Topics: Math, String
Acceptance Rate: 54.6%

## Hints

### Hint 1

Comparing dates in "MM-DD" form is awkward. Is there a single number you could turn each date into so that comparing and subtracting dates becomes easy?

### Hint 2

Convert each date into its "day of the year" (1 to 365) using a prefix sum of the days per month. Then each person's stay is just an integer interval.

### Hint 3

The overlap of two inclusive intervals `[a1, b1]` and `[a2, b2]` is `[max(a1, a2), min(b1, b2)]`. Its length is `min(b1, b2) - max(a1, a2) + 1`, clamped to 0 if negative.

## Approach

1. **Parse each date.** From `"MM-DD"`, read the month `MM` and day `DD` as integers (characters 0-1 and 3-4).
2. **Convert to day-of-year.** Precompute the cumulative days before each month: `prefix[m] = days[0] + ... + days[m-1]`. Then `dayOfYear = prefix[MM-1] + DD`.
3. **Intersect the intervals.** Alice is present on `[A1, A2]` and Bob on `[B1, B2]`. They overlap from `start = max(A1, B1)` to `end = min(A2, B2)`.
4. **Count.** If `end >= start`, the answer is `end - start + 1`; otherwise 0.

Example 1: Alice `08-15..08-18` → days `227..230`; Bob `08-16..08-19` → `228..231`. Overlap `228..230` → 3 days.

Why it works: mapping dates to day-of-year is order-preserving and turns the calendar into a number line, so the problem reduces to the classic interval-intersection formula.

## Complexity Analysis

Time Complexity: O(1) — fixed-length strings and a 12-entry month table.
Space Complexity: O(1)

## Edge Cases

- **No overlap** (one leaves before the other arrives): `end - start + 1` becomes ≤ 0, so clamp to 0.
- **Touching on a single day** (one leaves the day the other arrives): inclusive bounds mean the answer is 1, not 0 — watch the `+1`.
- **Same arrival and departure for both**: overlap is the whole interval.
- **One interval fully inside the other**: the answer is the inner interval's length.
- **Spanning month boundaries / whole year** (`01-01` to `12-31`): verifies the prefix sums are right (365 days).
