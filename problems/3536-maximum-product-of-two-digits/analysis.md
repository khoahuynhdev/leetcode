# 3536. Maximum Product of Two Digits

[LeetCode Link](https://leetcode.com/problems/maximum-product-of-two-digits/)

Difficulty: Easy
Topics: Math, Sorting
Acceptance Rate: 71.4%

## Hints

### Hint 1

There are at most 10 digits in `n`, so brute force over every pair is completely
affordable. But before you write the double loop, ask yourself what actually
makes a product large here. The two ingredients are the *digit extraction* (a
pure Math step) and the *selection* of which two digits to combine (a Sorting /
order-statistics step). Think about whether you need to compare all pairs at
all, or whether the answer is determined by only a couple of the digits.

### Hint 2

Peel the digits off with repeated `% 10` and `/ 10` — no string conversion
needed. Once you have the multiset of digits, the question becomes: which two
elements of a list of non-negative numbers have the largest product? For general
integers that question is tricky (two big negatives beat two positives), but
digits have a property that makes it much simpler.

### Hint 3

Every digit is in `[0, 9]`, so all candidates are **non-negative**. That kills
the "two negatives multiply to a big positive" trap entirely: the product is
monotonic in each factor, so picking the largest digit and the second-largest
digit is always optimal. You never need to sort, and you never need the pair
loop — one pass tracking the top two values is enough. And because the problem
lets you reuse a digit that appears more than once, "second largest" means
second largest *by position*, not second largest *distinct* value: for `n = 22`
the two factors are both `2`.

## Approach

**Step 1 — extract the digits.** Repeatedly take `n % 10` to get the least
significant digit, then `n /= 10` to drop it, until `n` becomes `0`. For
`n = 124` this yields `4`, `2`, `1`. The order doesn't matter, since we only
care about the two largest.

**Step 2 — select the two largest digits.** Because digits are non-negative,
`a * b` is non-decreasing in both `a` and `b`. So if `best` is the maximum digit
and `second` is the next digit in sorted-by-position order, no other pair can
beat `best * second`. Formally: for any valid pair `(x, y)` we have
`x <= best` and `y <= second` (after ordering `x >= y`), hence
`x * y <= best * second`.

We can find those two values in the same pass that extracts them, using the
classic two-variable "top-two" scan:

- if the current digit `d` is greater than `best`, then `best` gets demoted to
  `second` and `d` becomes the new `best`;
- otherwise, if `d` is greater than `second`, `d` becomes the new `second`;
- otherwise `d` is irrelevant.

The crucial detail is the `else if`: it uses `>` on `best` but still lets an
*equal* value fall through to the `second` slot. That is what makes duplicate
digits work. Tracing `n = 22`: first digit `2` gives `best = 2, second = 0`;
second digit `2` is not `> best`, but it is `> second`, so `second = 2`. Answer
`4`, as required.

**Step 3 — multiply.** Return `best * second`.

Initializing both trackers to `0` is safe here: digits are never negative, so a
`0` sentinel can only ever be replaced by a real digit or be a genuinely correct
answer (`n = 10` really does have `0` as its second-largest digit, and `1 * 0 =
0` is the right output).

Sorting the digit slice and multiplying the last two elements is an equally
correct and very readable alternative — `O(d log d)` instead of `O(d)`, which is
irrelevant for `d <= 10`. The one-pass version is shown because it avoids
allocating a slice at all.

## Complexity Analysis

Time Complexity: O(log n) — one iteration per decimal digit, so at most 10
iterations for the given constraints. Effectively O(1).
Space Complexity: O(1) — two integer trackers, no digit slice.

## Edge Cases

- **Repeated digits (`n = 22`, `n = 505`).** The same digit value must be usable
  twice when it occurs at two positions. This is exactly the case a "second
  largest *distinct* value" implementation gets wrong — it would report `0` for
  `22`. Use `>` when demoting into `second`, not `>=` on the value comparison.
- **Two-digit input (`n = 31`).** The minimum allowed input size. There is
  exactly one pair, so the answer is forced; a good sanity check that the
  sentinel initialization isn't leaking into the result.
- **Zeros in the number (`n = 10`, `n = 90`, `n = 1000000000`).** A zero digit
  forces the product to `0` whenever it is one of the top two. The answer really
  is `0` — don't special-case it away or skip zero digits while scanning.
- **All digits identical and maximal (`n = 999999999`).** Answer `81`, the
  largest possible output. Confirms nothing overflows and that the `else if`
  branch keeps firing correctly on a long run of ties.
- **Leading digit is the largest vs. smallest (`n = 91` vs `n = 19`).** Since we
  peel digits from the least significant end, both orders must produce `9`. This
  catches an implementation that only ever updates `best` and forgets the
  demotion step.
